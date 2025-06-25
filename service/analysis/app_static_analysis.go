package analysis

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/nzmxd/app-insight/core/downloader"
	aGlobal "github.com/nzmxd/app-insight/global"
	"github.com/nzmxd/app-insight/model/analysis"
	"github.com/nzmxd/app-insight/model/analysis/request"
	"github.com/nzmxd/app-insight/model/apprank"
	"github.com/nzmxd/bserver/global"
	"github.com/nzmxd/bserver/utils"
	"go.uber.org/zap"
	"strings"
	"time"
)

type AppStaticAnalysisTaskService struct{}

func (a *AppStaticAnalysisTaskService) EnqueueStaticAnalysisTask(staticAnalysisTask analysis.AppAnalysisTaskPayload) error {
	payload, err := json.Marshal(staticAnalysisTask)
	if err != nil {
		return err
	}
	downloadTask := asynq.NewTask(aGlobal.StaticAnalysisTask, payload, asynq.Retention(24*time.Hour))
	_, err = aGlobal.AsynqClient.Enqueue(downloadTask,
		asynq.Queue("analysis"),
		asynq.TaskID(*staticAnalysisTask.AppID),
		asynq.MaxRetry(1),
	)
	return err
}

func (a *AppStaticAnalysisTaskService) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var staticAnalysisTask analysis.AppAnalysisTaskPayload
	// Step 1: 解码任务
	if err := json.Unmarshal(t.Payload(), &staticAnalysisTask); err != nil {
		return err
	}

	// Step 2: 获取任务信息
	appDownloadTask, err := downloadTaskService.GetDownloadTaskByAppIdAndVersion(*staticAnalysisTask.AppID, *staticAnalysisTask.VersionCode)
	if err != nil {
		return err
	}
	appAnalysisTask, err := appAnalysisTaskService.GetAppAnalysisTaskByAppIdAndVersion(*staticAnalysisTask.AppID, *staticAnalysisTask.VersionCode)
	if err != nil {
		return err
	}
	defer func() {
		if err = appAnalysisTaskService.UpdateAppAnalysisTask(&appAnalysisTask); err != nil {
			global.LOG.Error("Failed to update download task", zap.Error(err))
		}
	}()
	now := time.Now()
	appAnalysisTask.FileAnalysisStatus = utils.Ptr(analysis.StatusAnalysis)
	appAnalysisTask.FileAnalysisFinishedAt = &now
	err = appAnalysisTaskService.UpdateAppAnalysisTask(&appAnalysisTask)
	if err != nil {
		return a.setTaskError(&appAnalysisTask, analysis.StatusFailed, err, true)
	}
	// Step 3: 检查本地文件
	downloadFilePath := *appDownloadTask.FilePath
	var localFilePath string
	if strings.HasPrefix(downloadFilePath, "http") {
		localFilePath, err = downloader.DownloadFile(downloadFilePath, aGlobal.Config.Downloader.SavePath, "", 0)
		if err != nil {
			return a.setTaskError(&appAnalysisTask, analysis.StatusFailed, err, true)
		}
	} else {
		localFilePath = downloadFilePath
	}
	defer func() {
		if downloadFilePath != localFilePath {
			utils.DeLFile(localFilePath)
		}
	}()

	// Step 4: 执行分析任务
	results, err := aGlobal.Analyzer.Analysis(localFilePath)
	if err != nil {
		return a.setTaskError(&appAnalysisTask, analysis.StatusFailed, err, true)
	}
	var sdkNames []string
	var sdkMatchResults []analysis.AppSdkMatchResult
	for _, result := range results {
		sdkMatchResults = append(sdkMatchResults, analysis.AppSdkMatchResult{
			AppAnalysisTaskId: appAnalysisTask.ID,
			SdkMetadataId:     int64(result.SdkInfoID),
		})
		sdkNames = append(sdkNames, result.SdkName)
	}

	// Step 5: 插入分析结果
	err = appSdkMatchResultService.BatchInsertAppSdkMatchResult(sdkMatchResults)
	if err != nil {
		return a.setTaskError(&appAnalysisTask, analysis.StatusFailed, err, true)
	}

	// Step 6: 更新任务状态
	now = time.Now()
	appAnalysisTask.VersionCode = appDownloadTask.VersionCode
	appAnalysisTask.Developer = appDownloadTask.Developer
	appAnalysisTask.IsGpListing = appDownloadTask.IsGpListing
	appAnalysisTask.FileAnalysisFinishedAt = &now
	appAnalysisTask.FileAnalysisStatus = utils.Ptr(analysis.StatusSuccess)

	// Step 7: 插入CH记录
	var staticAnalysisDetail apprank.AppStaticAnalysisDetail
	staticAnalysisDetail.AppID = *appAnalysisTask.AppID
	staticAnalysisDetail.VersionCode = *appAnalysisTask.VersionCode
	staticAnalysisDetail.VersionName = *appAnalysisTask.VersionName
	staticAnalysisDetail.Developer = *appAnalysisTask.Developer
	staticAnalysisDetail.IsGoogleListing = *appAnalysisTask.IsGpListing
	staticAnalysisDetail.SdkNames = sdkNames
	if downloadFilePath != localFilePath {
		staticAnalysisDetail.FilePath = downloadFilePath
	}
	staticAnalysisDetail.CreatedAt = now

	_ = global.CH.Create(&staticAnalysisDetail).Error
	return nil
}

// setTaskError 封装错误处理
func (a *AppStaticAnalysisTaskService) setTaskError(task *analysis.AppAnalysisTask, status int, err error, skipRetry bool) error {
	task.FileAnalysisStatus = utils.Ptr(status)
	msg := err.Error()
	task.ErrorMessage = &msg
	global.LOG.Error("ProcessTask error",
		zap.String("app_id", *task.AppID),
		zap.String("error", msg),
		zap.Int("status", status),
	)
	if skipRetry {
		return asynq.SkipRetry
	}
	return err
}

func (a *AppStaticAnalysisTaskService) GetStaticAnalysisStats(req request.StaticAnalysisStatsSearch) ([]apprank.StaticAnalysisStatsResp, error) {
	sql := `
		SELECT
			sdk_name,
			countIf(app_id, is_google_listing = 1) AS google_listing_count,
			countIf(app_id, is_google_listing = 0) AS non_google_listing_count
		FROM (
			SELECT
				app_id,
				is_google_listing,
				arrayJoin(sdk_names) AS sdk_name,
				created_at
			FROM app_static_analysis_detail
			WHERE created_at >= ? AND created_at < ?
		)
-- 		WHERE sdk_name IN 
		GROUP BY sdk_name
		ORDER BY google_listing_count DESC
	`
	var results []apprank.StaticAnalysisStatsResp
	err := global.CH.Raw(sql, req.StartTime, req.EndTime).Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
