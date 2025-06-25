package request

import "time"

// AdSlotPushLogSearch  广告位推送记录
type AdSlotPushLogSearch struct {
	Platform   *string   `json:"platform" gorm:"platform" form:"platform"`          // 渠道名称
	PushStatus *uint8    `json:"push_status" gorm:"push_status" form:"pushStatus" ` // 0-推送失败 1-推送成功
	StartTime  time.Time `gorm:"column:startTime" json:"startTime" form:"startTime" time_format:"unix"`
	EndTime    time.Time `gorm:"column:endTime" json:"endTime" form:"endTime"   time_format:"unix"`
}
