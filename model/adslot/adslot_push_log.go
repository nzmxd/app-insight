package adslot

import "time"

// AdSlotPushLog 广告位推送记录
type AdSlotPushLog struct {
	ID         int64      `json:"id" gorm:"id"`                   // 主键ID
	Platform   string     `json:"platform" gorm:"platform"`       // 渠道名称
	PushStatus uint8      `json:"push_status" gorm:"push_status"` // 0-推送失败 1-推送成功
	CreatedAt  *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"deleted_at"`
	CommonAdSlotCoreParams
}

// TableName 表名称
func (*AdSlotPushLog) TableName() string {
	return "ad_unit_push_log"
}

type AdUnitPushLogStatsResult struct {
	Platform     string `json:"platform"`
	Os           uint8  `json:"os"`
	SuccessValid int64  `json:"success_valid"` // 推送成功 且 有效
	FailValid    int64  `json:"fail_valid"`    // 推送失败 且 有效
	FailInvalid  int64  `json:"fail_invalid"`  // 推送失败 且 无效
}
