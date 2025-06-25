package analysis

// AppSdkMatchRule SDK匹配规则
type AppSdkMatchRule struct {
	ID              int64  `json:"id" gorm:"id"`                               // 主键ID
	SdkMetadataId   int64  `json:"sdk_metadata_id" gorm:"sdk_metadata_id"`     // SDK信息主键ID
	XmlScanRules    string `json:"xml_scan_rules" gorm:"xml_scan_rules"`       // xml匹配规则
	GlobalScanRules string `json:"global_scan_rules" gorm:"global_scan_rules"` // 全局匹配规则
	BkString1       string `json:"bk_string1" gorm:"bk_string1"`               // 备用字符串字段1
	BkString2       string `json:"bk_string2" gorm:"bk_string2"`               // 备用字符串字段2
	BkInt1          int64  `json:"bk_int1" gorm:"bk_int1"`                     // 备用Int字段1
}

// TableName 表名称
func (*AppSdkMatchRule) TableName() string {
	return "app_sdk_match_rule"
}
