package analysis

// AppSdkMatchResult APK-SDK匹配结果表
type AppSdkMatchResult struct {
	ID                int64 `json:"id" gorm:"id"`                                     // 主键ID
	AppAnalysisTaskId int64 `json:"app_analysis_task_id" gorm:"app_analysis_task_id"` // 包名主键ID
	SdkMetadataId     int64 `json:"sdk_metadata_id" gorm:"sdk_metadata_id"`           // SDK信息主键ID
}

// TableName 表名称
func (*AppSdkMatchResult) TableName() string {
	return "app_sdk_match_result"
}
