package analysis

import (
	"app-insight/model/analysis"
	"app-insight/model/download"
	"errors"
	"github.com/nzmxd/bserver/global"
	"github.com/nzmxd/bserver/utils"
	"go.uber.org/zap"
	"time"
)

type AppAnalysisTaskService struct{}

func (a *AppAnalysisTaskService) CreateAppAnalysisTask(task analysis.AppAnalysisTask) (err error) {
	if task.AppID == nil || *task.AppID == "" {
		return errors.New("AppID 不能为空")
	}

	// 判断 AppID 是否已存在
	var count int64
	err = global.DB.Model(&analysis.AppAnalysisTask{}).
		Where("app_id = ?", *task.AppID).
		Where("version_code = ?", *task.VersionCode).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		// 已存在，不插入
		global.LOG.Info("该版本 AppID 已存在，跳过插入", zap.String("appID", *task.AppID), zap.String("versionCode", *task.VersionCode))
		return nil
	}
	if task.FileAnalysisStatus == nil {
		task.FileAnalysisStatus = utils.Ptr(analysis.StatusPending)
	}
	// 插入记录
	err = global.DB.Create(task).Error
	return err
}

func (a *AppAnalysisTaskService) UpdateAppAnalysisTask(task *analysis.AppAnalysisTask) (err error) {
	err = global.DB.Model(&analysis.AppAnalysisTask{}).Where("id = ?", task.ID).Updates(&task).Error
	return err
}

func (a *AppAnalysisTaskService) GetAppAnalysisTask(id string) (task analysis.AppAnalysisTask, err error) {
	err = global.DB.Where("id = ?", id).First(&task).Error
	return
}

func (a *AppAnalysisTaskService) GetAppAnalysisTaskByAppId(appId string) (task analysis.AppAnalysisTask, err error) {
	err = global.DB.Where("app_id = ?", appId).First(&task).Error
	return
}

func (a *AppAnalysisTaskService) GetAppAnalysisTaskByAppIdAndVersion(appId string, versionCode string) (task analysis.AppAnalysisTask, err error) {
	err = global.DB.
		Where("app_id = ?", appId).
		Where("version_code = ?", versionCode).
		First(&task).Error
	return
}

func (a *AppAnalysisTaskService) GetAnalysisTaskFormDownloadTask(limit int) (task []analysis.AppAnalysisTask, err error) {
	var downloadTasks []download.AppDownloadTask

	// 获取当天零点时间
	//today := time.Now().Truncate(24 * time.Hour)

	// 查询满足条件的下载任务
	err = global.DB.Table("app_download_task AS d").
		Select("d.*").
		Joins(`LEFT JOIN app_analysis_task AS a 
		       ON d.app_id = a.app_id AND d.version_code = a.version_code`).
		Where("d.status = ?", analysis.StatusSuccess).
		Where("a.app_id IS NULL").
		//Where("d.created_at >= ?", today).
		Order("d.created_at DESC").
		Limit(limit).
		Find(&downloadTasks).Error
	if err != nil {
		return nil, err
	}

	// 转换为 AppAnalysisTask 列表
	var analysisTasks []analysis.AppAnalysisTask
	now := time.Now()

	for _, dt := range downloadTasks {
		analysisTask := analysis.AppAnalysisTask{
			AppID:                 dt.AppID,
			VersionCode:           dt.VersionCode,
			VersionName:           dt.VersionName,
			IsGpListing:           dt.IsGpListing,
			Developer:             dt.Developer,
			FileAnalysisStatus:    utils.Ptr(analysis.StatusPending),
			FileAnalysisStartedAt: &now,
		}
		analysisTasks = append(analysisTasks, analysisTask)
	}

	// 插入新任务
	if len(analysisTasks) > 0 {
		err = global.DB.Create(&analysisTasks).Error
		if err != nil {
			return nil, err
		}
	}

	return analysisTasks, nil
}
