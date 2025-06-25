package initialize

import (
	"app-insight/core/downloader"
	"app-insight/global"
)

func InitDownloader() {
	downloaderList := make(map[string]downloader.Downloader)
	switch global.Config.Downloader.Source {
	case "apkpure":
		global.Downloader = &downloader.ApkpureDownloader{ProxyUrl: global.Config.Downloader.Proxy, Timeout: global.Config.Downloader.Timeout}
	case "googleplay":
		global.Downloader = &downloader.GooglePlayDownloader{ProxyUrl: global.Config.Downloader.Proxy}
	default:
		global.Downloader = &downloader.ApkpureDownloader{ProxyUrl: global.Config.Downloader.Proxy}
	}
	downloaderList["googleplay"] = &downloader.GooglePlayDownloader{ProxyUrl: global.Config.Downloader.Proxy}
	global.DownloaderList = downloaderList
}
