package downloader

import (
	"errors"
	"fmt"
	"github.com/nzmxd/bserver/utils"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var (
	AppNotFoundErr = errors.New("app not found")
)

type GenericAppDetail map[string]interface{}
type GenericAppVersionList []GenericAppDetail

// Downloader 定义了应用下载器的接口，提供应用信息获取、版本管理、下载、校验等功能。
// 适用于对接应用商店、企业分发平台、第三方市场等多种应用源。
type Downloader interface {
	// GetAppDetail 返回指定 appid 的应用基础信息（如名称、版本、包名等）。
	GetAppDetail(appid string) (GenericAppDetail, error)

	// GetAppDownloadUrl 获取指定 appid 和版本的下载链接（URL），不执行实际下载。
	GetAppDownloadUrl(appid string, version string) string

	// ListVersions 返回指定 appid 的所有可用历史版本。
	ListVersions(appid string) (GenericAppVersionList, error)

	// Download 执行下载操作，返回本地保存路径。
	Download(appid string, version string, saveDir string) (string, error)

	// CheckUpdate 判断本地版本是否需要更新，返回是否更新、最新版本号。
	CheckUpdate(appid string, localVersion string) (bool, string, error)

	// Validate 校验指定 appid 和版本是否有效或可用。
	Validate(appid string, version string) error

	// SetProxyUrl 设置下载所使用的代理地址（如 HTTP/SOCKS5），便于跨境或局域网访问。
	SetProxyUrl(proxyUrl string) error
}

func DownloadFile(fileURL, saveDir, proxyAddr string, timeout int) (string, error) {
	if timeout == 0 {
		timeout = 600
	}
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	if proxyAddr != "" {
		proxyURL, err := url.Parse(proxyAddr)
		if err != nil {
			return "", fmt.Errorf("代理地址无效: %w", err)
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
	resp, err := client.Get(fileURL)
	if err != nil {
		return "", fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	var fileName string
	if cd := resp.Header.Get("Content-Disposition"); cd != "" {
		_, params, err := mime.ParseMediaType(cd)
		if err == nil && params["filename"] != "" {
			fileName = params["filename"]
		}
	}
	if fileName == "" {
		parsedURL, err := url.Parse(fileURL)
		if err != nil {
			return "", fmt.Errorf("URL 解析失败: %w", err)
		}
		fileName = path.Base(parsedURL.Path)
		if fileName == "" || strings.HasSuffix(fileName, "/") {
			fileName = "downloaded_file"
		}
	}
	// 创建保存目录
	_ = utils.CreateDir(saveDir)
	//if err = utils.CreateDir(saveDir); err != nil {
	//	return "", fmt.Errorf("创建目录失败: %w", err)
	//}

	// 拼接绝对路径
	fullPath := filepath.Join(saveDir, fileName)
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return "", fmt.Errorf("获取绝对路径失败: %w", err)
	}

	// 创建文件并保存内容
	out, err := os.Create(absPath)
	if err != nil {
		return "", fmt.Errorf("创建文件失败: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("写入文件失败: %w", err)
	}

	return absPath, nil
}
