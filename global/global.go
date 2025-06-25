package global

import (
	"github.com/hibiken/asynq"
	"github.com/nzmxd/app-insight/config"
	"github.com/nzmxd/app-insight/core/analyzer"
	"github.com/nzmxd/app-insight/core/downloader"
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
