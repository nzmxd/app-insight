package scheduler

import (
	"context"
	aGlobal "github.com/nzmxd/app-insight/global"
	"github.com/nzmxd/app-insight/model/apprank"
	"github.com/nzmxd/app-insight/model/download"
	"github.com/nzmxd/bserver/global"
	"github.com/nzmxd/bserver/utils/record"
	"go.uber.org/zap"
)

// SyncDownloadTask 同步ASO的Android包名到AppDownloadTask下载任务列表中
func SyncDownloadTask() error {
	androidBasicInfoRecord := record.GormRecord[apprank.AppBasicInfoAndroid]{
		DB:            aGlobal.AppRankOnlineDB,
		Redis:         global.REDIS,
		RedisKey:      "app_sys_v1:app_downloader:id_record:app_basic_info_android",
		StartRowID:    20000000,
		SelectFields:  []string{"id", "app_id"},
		QueryModifier: nil,
	}
	var downloadPendingCount int64
	err := global.DB.Model(&download.AppDownloadTask{}).Where("status = ?", download.StatusPending).Count(&downloadPendingCount).Error
	if err != nil {
		global.LOG.Error("查询待下载任务数失败", zap.Error(err))
		return err
	}
	if downloadPendingCount > 500 {
		global.LOG.Info("当前下载任务数大于指定值无需添加下载任务", zap.Int64("download_pending_count", downloadPendingCount))
		return nil
	}

	basicInfoAndroids, err := androidBasicInfoRecord.Fetch(context.Background(), 1000)
	if err != nil {
		global.LOG.Error("从apprank数据库获取待下载包名失败", zap.Error(err))
		return err
	}
	var appIdList []string
	for _, basicInfoAndroid := range basicInfoAndroids {
		appIdList = append(appIdList, basicInfoAndroid.AppId)
	}
	inserted, skipped, err := downloadTaskService.BatchInsertByAppId(appIdList)
	if err != nil {
		global.LOG.Error("批量插入Android包名失败", zap.Error(err))
		return err
	}
	global.LOG.Info("批量插入Android包名成功", zap.Int("total", len(appIdList)), zap.Int64("inserted", inserted), zap.Int64("skipped", skipped))
	return nil
}
