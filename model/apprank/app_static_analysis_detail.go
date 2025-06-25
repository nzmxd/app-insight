package apprank

import "time"

type AppStaticAnalysisDetail struct {
	AppID           string    `json:"app_id"`
	SdkNames        []string  `gorm:"column:sdk_names;type:Array(String)"`
	VersionCode     string    `json:"version_code"`
	VersionName     string    `json:"version_name"`
	Developer       string    `json:"developer"`
	IsGoogleListing bool      `json:"is_google_listing"`
	FilePath        string    `json:"file_path"`
	CreatedAt       time.Time `json:"created_at"`
}

func (a *AppStaticAnalysisDetail) TableName() string {
	return "app_static_analysis_detail"
}

type StaticAnalysisStatsResp struct {
	SdkName               string `json:"sdk_name"`
	GoogleListingCount    int    `json:"google_listing_count"`
	NonGoogleListingCount int    `json:"non_google_listing_count"`
}
