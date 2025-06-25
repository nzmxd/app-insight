package analysis

import "time"

const (
	StatusPending  = iota // 0: 初始状态，任务刚创建，未入队
	StatusQueued          // 1: 已加入队列
	StatusAnalysis        // 2: 正在分析
	StatusSuccess         // 3: 分析成功
	StatusFailed          // 4: 分析失败
	StatusRetrying        // 5: 正在重试
)

type AppAnalysisTaskPayload struct {
	AppID       *string `json:"app_id"`
	VersionCode *string `json:"version_code" gorm:"version_code"`
}

// AppAnalysisTask 应用分析任务表
type AppAnalysisTask struct {
	ID                        int64      `json:"id" gorm:"id"`                                     // 主键ID
	AppID                     *string    `json:"app_id" gorm:"app_id"`                             // 应用ID
	VersionCode               *string    `json:"version_code" gorm:"version_code"`                 // 版本号
	VersionName               *string    `json:"version_name" gorm:"version_name"`                 // 版本名
	IsGpListing               *bool      `json:"is_gp_listing" gorm:"is_gp_listing"`               // GooglePlay上架:0=否,1=是
	Developer                 *string    `json:"developer" gorm:"developer"`                       // 开发商
	ErrorMessage              *string    `json:"error_message" gorm:"error_message"`               // 错误信息（仅在失败时记录）
	FileAnalysisStatus        *int       `json:"file_analysis_status" gorm:"file_analysis_status"` // 文件分析状态
	FileAnalysisStartedAt     *time.Time `json:"file_analysis_started_at" gorm:"file_analysis_started_at"`
	FileAnalysisFinishedAt    *time.Time `json:"file_analysis_finished_at" gorm:"file_analysis_finished_at"`
	CodeAnalysisStatus        *int       `json:"code_analysis_status" gorm:"code_analysis_status"` // 代码分析状态
	CodeAnalysisStartedAt     *time.Time `json:"code_analysis_started_at" gorm:"code_analysis_started_at"`
	CodeAnalysisFinishedAt    *time.Time `json:"code_analysis_finished_at" gorm:"code_analysis_finished_at"`
	DynamicAnalysisStatus     *int       `json:"dynamic_analysis_status" gorm:"dynamic_analysis_status"` // 动态分析状态
	DynamicAnalysisStartedAt  *time.Time `json:"dynamic_analysis_started_at" gorm:"dynamic_analysis_started_at"`
	DynamicAnalysisFinishedAt *time.Time `json:"dynamic_analysis_finished_at" gorm:"dynamic_analysis_finished_at"`
	CreatedAt                 time.Time  `json:"created_at" gorm:"created_at"` // 任务创建时间
	UpdatedAt                 time.Time  `json:"updated_at" gorm:"updated_at"` // 最后更新时间
}

// TableName 表名称
func (a *AppAnalysisTask) TableName() string {
	return "app_analysis_task"
}
