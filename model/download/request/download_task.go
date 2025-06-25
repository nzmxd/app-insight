package request

import "time"

type DownloadStatsSearch struct {
	StartTime time.Time `gorm:"column:startTime" json:"startTime" form:"startTime" time_format:"unix"`
	EndTime   time.Time `gorm:"column:endTime" json:"endTime" form:"endTime"   time_format:"unix"`
}

type DownloadUrlSearch struct {
	AppID       *string `gorm:"column:app_id" json:"appId"  form:"appId"`                   // 应用 ID
	VersionCode *string `gorm:"column:version_code" json:"versionCode"  form:"versionCode"` // 应用版本号（code）
}
