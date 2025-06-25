package global

import (
	"app-insight/config"
	"app-insight/core/analyzer"
	"app-insight/core/downloader"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

const (
	DownloadTask       = "app_insight:download"
	StaticAnalysisTask = "app_insight:static_analysis"
)

var (
	Config          config.Server
	AppRankDB       *gorm.DB
	AppRankOnlineDB *gorm.DB
	SpRawDB         *gorm.DB
	Analyzer        analyzer.StaticAnalyzer
	Downloader      downloader.Downloader
	DownloaderList  map[string]downloader.Downloader
	AsynqClient     *asynq.Client
)
