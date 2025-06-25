package download

import "time"

const (
	StatusPending     = iota // 0: 初始状态，任务刚创建，未入队
	StatusQueued             // 1: 已加入下载队列，等待下载
	StatusDownloading        // 2: 正在下载
	StatusSuccess            // 3: 下载成功
	StatusFailed             // 4: 下载失败
	StatusRetrying           // 5: 正在重试
)

type AppDownloadTask struct {
	ID           int64      `gorm:"column:id" json:"id"`                       // 主键 ID
	AppID        *string    `gorm:"column:app_id" json:"app_id"`               // 应用 ID
	VersionCode  *string    `gorm:"column:version_code" json:"version_code"`   // 应用版本号（code）
	VersionName  *string    `gorm:"column:version_name" json:"version_name"`   // 应用版本名（name）
	IsGpListing  *bool      `gorm:"column:is_gp_listing" json:"is_gp_listing"` // 应用是否在GooglePlay上架
	Developer    *string    `gorm:"column:developer" json:"developer"`         // 应用开发商
	Source       *string    `gorm:"column:source" json:"source"`               // 下载来源
	Status       *int       `gorm:"column:status" json:"status"`               // 下载状态（0-pending，1-queued 2-downloading，3-success，4-failed，5-retrying）
	RetryCount   *int       `gorm:"column:retry_count" json:"retry_count"`     // 下载重试次数
	ErrorMessage *string    `gorm:"column:error_message" json:"error_message"` // 错误信息
	FilePath     *string    `gorm:"column:file_path" json:"file_path"`         // 下载文件路径
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`       // 创建时间
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`       // 更新时间
	StartedAt    *time.Time `gorm:"column:started_at" json:"started_at"`       // 下载任务开始时间
	FinishedAt   *time.Time `gorm:"column:finished_at" json:"finished_at"`     // 下载完成时间（任务成功时更新）
}

func (AppDownloadTask) TableName() string {
	return "app_download_task"
}

type DownloadStatsResult struct {
	PendingCount       int64   `json:"pending_count"`         // 等待中的任务数
	QueueCount         int64   `json:"queue_count"`           // 任务队列中的总数
	RetryCount         int64   `json:"retry_count"`           // 重试中的任务数
	FailedCount        int64   `json:"failed_count"`          // 失败的任务数
	FinishedGpCount    int64   `json:"finished_gp_count"`     // 下载完成并上架 GP 的任务数
	FinishedNonGpCount int64   `json:"finished_non_gp_count"` // 下载完成但未上架 GP 的任务数
	AvgDownloadSeconds float64 `json:"avg_download_seconds"`  // 平均下载耗时（秒）
	TotalCount         int64   `json:"total_count"`           // 所有任务总数（匹配条件内）
}
