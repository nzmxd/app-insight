package adslot

import (
	"github.com/nzmxd/app-insight/model/adslot"
	"github.com/nzmxd/app-insight/model/adslot/request"
	"github.com/nzmxd/bserver/global"
)

type AdSlotPushLogService struct{}

func (a *AdSlotPushLogService) CreatAdUnitPushLog(pushLog *adslot.AdSlotPushLog) (err error) {
	err = global.DB.Create(pushLog).Error
	return err
}

func (a *AdSlotPushLogService) GetAdUnitPushLogStats(search request.AdSlotPushLogSearch) ([]adslot.AdUnitPushLogStatsResult, error) {
	var results []adslot.AdUnitPushLogStatsResult

	// 构造初始查询
	db := global.DB.Model(&adslot.AdSlotPushLog{})

	// 条件：平台
	if search.Platform != nil && *search.Platform != "" {
		db = db.Where("platform = ?", *search.Platform)
	}

	// 条件：推送状态（如果你希望以某一状态筛选整体数据）
	if search.PushStatus != nil {
		db = db.Where("push_status = ?", *search.PushStatus)
	}

	// 条件：时间区间
	if !search.StartTime.IsZero() && !search.EndTime.IsZero() {
		db = db.Where("created_at BETWEEN ? AND ?", search.StartTime, search.EndTime)
	}

	// 构建分组查询
	db = db.Select(`
		platform,
		os,
		SUM(CASE WHEN push_status = 1 AND is_available = 1 THEN 1 ELSE 0 END) AS success_valid,
		SUM(CASE WHEN push_status = 0 AND is_available = 1 THEN 1 ELSE 0 END) AS fail_valid,
		SUM(CASE WHEN push_status = 0 AND is_available = 0 THEN 1 ELSE 0 END) AS fail_invalid
	`).Group("platform, os")

	err := db.Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (a *AdSlotPushLogService) ExistsByParamsDictMd5(md5 string) (bool, error) {
	var count int64
	err := global.DB.Model(&adslot.AdSlotPushLog{}).
		Where("params_dict_md5 = ?", md5).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
