package analysis

import "time"

// SdkMetadata SDK元信息表
type SdkMetadata struct {
	ID               int64     `json:"id" gorm:"id"`                                 // 主键ID
	SdkName          string    `json:"sdk_name" gorm:"sdk_name"`                     // SDK名称
	DeveloperName    string    `json:"developer_name" gorm:"developer_name"`         // 开发者名称
	IconFileUrl      string    `json:"icon_file_url" gorm:"icon_file_url"`           // 图标文件URL
	MavenIdentifiers string    `json:"maven_identifiers" gorm:"maven_identifiers"`   // Maven标识符（JSON数组格式）
	IsPlayRegistered int8      `json:"is_play_registered" gorm:"is_play_registered"` // 是否在Play商店注册（0表示否，1表示是）
	SdkPageUrl       string    `json:"sdk_page_url" gorm:"sdk_page_url"`             // SDK页面URL
	Repository       int8      `json:"repository" gorm:"repository"`                 // 是否为仓库（0表示否，1表示是）
	PrivateRepoUrl   string    `json:"private_repo_url" gorm:"private_repo_url"`     // 私有仓库URL
	PublicCategories string    `json:"public_categories" gorm:"public_categories"`   // 公开分类（JSON数组格式）
	DataSafetyUrl    string    `json:"data_safety_url" gorm:"data_safety_url"`       // 数据安全URL
	IsGoogleOwned    int8      `json:"is_google_owned" gorm:"is_google_owned"`       // 是否为Google所有（0表示否，1表示是）
	IsDeprecated     int8      `json:"is_deprecated" gorm:"is_deprecated"`           // 是否已弃用（0表示否，1表示是）
	CreatedAt        time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at" gorm:"deleted_at"`
}

// TableName 表名称
func (*SdkMetadata) TableName() string {
	return "sdk_metadata"
}
