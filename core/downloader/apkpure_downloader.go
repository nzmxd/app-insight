package downloader

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
	"time"
)

type ApkpureDownloader struct {
	ProxyUrl string
	Timeout  int
}

func (a *ApkpureDownloader) GetAppDetail(appid string) (GenericAppDetail, error) {
	getAppDetailUrl := "https://tapi.pureapk.com/v3/get_app_detail"

	payload := []byte(fmt.Sprintf(`{"action":"Download","ad":false,"ad_source_type":0,"arg":"","filter_package_name":[],"has_installed_version_code":0,"package_name":"%s","page":"Detail","sdk_ads":{"module_ads":[{"ads":[],"module_id":"","module_name":"recommend_ad","cached_size":-1}],"cached_size":-1},"cached_size":-1}`, appid))

	req, err := http.NewRequest("POST", getAppDetailUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Host", "tapi.pureapk.com")
	req.Header.Set("Cookie", "")
	req.Header.Set("user-agent-webview", "Mozilla/5.0 (Linux; Android 9; RMX3820; zh-CN) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Mobile Safari/537.36 Mobile Safari/537.36")
	req.Header.Set("user-agent", "Dalvik/2.1.0 (Linux; U; Android 9; RMX3820 Build/PQ3A.190605.06051204); APKPure/3.20.49 (Aegon)")
	req.Header.Set("ual-access-businessid", "projecta")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	if a.ProxyUrl != "" {
		proxyURL, _ := url.Parse(a.ProxyUrl)
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	retcode := gjson.GetBytes(body, "retcode").Int()
	errmsg := gjson.GetBytes(body, "errmsg").String()
	if retcode != 0 {
		if retcode == 1001 {
			return nil, AppNotFoundErr
		}
		return nil, fmt.Errorf("retcode: %d, errmsg: %s", retcode, errmsg)
	}

	appDetailJson := gjson.GetBytes(body, "app_detail")
	if !appDetailJson.Exists() {
		return nil, fmt.Errorf("no app_detail found in response")
	}
	genericAppDetail := a.ParseGenericAppDetail(appDetailJson)
	return genericAppDetail, nil
}

func (a *ApkpureDownloader) ListVersions(appid string) (GenericAppVersionList, error) {
	getAppVersionUrl := fmt.Sprintf("https://tapi.pureapk.com/v3/get_app_his_version?package_name=%s&hl=cn", appid)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, getAppVersionUrl, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Host", "tapi.pureapk.com")
	req.Header.Add("user-agent-webview", "Mozilla/5.0 (Linux; Android 9; RMX3820; zh-CN) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Mobile Safari/537.36 Mobile Safari/537.36")
	req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 9; RMX3820 Build/PQ3A.190605.06051204); APKPure/3.20.49 (Aegon)")
	req.Header.Add("ual-access-businessid", "projecta")
	req.Header.Add("ual-access-projecta", "{\"device_info\":{\"abis\":[\"x86_64\",\"x86\",\"arm64-v8a\",\"armeabi-v7a\",\"armeabi\"],\"android_id\":\"eebd4c1aa736f267\",\"brand\":\"realme\",\"country\":\"\",\"country_code\":\"US\",\"imei\":\"\",\"language\":\"zh-CN\",\"manufacturer\":\"realme\",\"mode\":\"RMX3820\",\"os_ver\":\"28\",\"os_ver_name\":\"9\",\"platform\":1,\"product\":\"RMX3820\",\"screen_height\":1920,\"screen_width\":1080},\"host_app_info\":{\"build_no\":\"641\",\"channel\":\"\",\"md5\":\"9e4824b79e3a27b9c706a8e0194ca20a\",\"pkg_name\":\"com.apkpure.aegon\",\"sdk_ver\":\"3.20.49\",\"version_code\":3204937,\"version_name\":\"3.20.49\"},\"net_info\":{\"carrier_code\":0,\"ipv4\":\"127.0.0.1\",\"ipv6\":\"\",\"mac_address\":\"00:DB:F1:C9:72:A4\",\"net_type\":1,\"use_vpn\":true,\"wifi_bssid\":\"02:00:00:00:00:00\",\"wifi_ssid\":\"<unknown ssid>\"}}") //req.Header.Add("ual-access-extinfo", "{\"ext_info\":\"{\\\"gaid\\\":\\\"\\\",\\\"oaid\\\":\\\"\\\"}\",\"lbs_info\":{\"accuracy\":0.0,\"city\":\"\",\"city_code\":0,\"country\":\"\",\"country_code\":\"\",\"district\":\"\",\"latitude\":0.0,\"longitude\":0.0,\"province\":\"\",\"street\":\"\"}}")
	//req.Header.Add("Cookie", "tgw_l7_route=8680cc90c90b4cde6bf3fc068564664d")
	//req.Header.Add("ual-access-sequence", "14d02393-a478-4431-8f5b-810825947686")
	//req.Header.Add("ual-access-signature", "")
	//req.Header.Add("ual-access-nonce", "40199865")
	//req.Header.Add("ual-access-timestamp", "1749454934063")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	retcode := gjson.GetBytes(body, "retcode").Int()
	errmsg := gjson.GetBytes(body, "errmsg").String()
	if retcode != 0 {
		return nil, errors.New(fmt.Sprintf("retcode:%d,errmsg:%s", retcode, errmsg))
	}
	versionList := gjson.GetBytes(body, "version_list").Array()
	list := GenericAppVersionList{}
	for _, result := range versionList {
		genericAppDetail := a.ParseGenericAppDetail(result)
		list = append(list, genericAppDetail)

	}
	return list, nil
}

func (a *ApkpureDownloader) SetProxyUrl(proxyUrl string) error {
	if !isValidURL(proxyUrl) {
		return errors.New("invalid proxyUrl")
	}
	a.ProxyUrl = proxyUrl
	return nil
}

func (a *ApkpureDownloader) GetAppDownloadUrl(appid string, version string) string {
	return ""
}

func (a *ApkpureDownloader) Download(appid string, version string, saveDir string) (string, error) {
	// 暂时不使用版本
	_ = version
	appDetail, err := a.GetAppDetail(appid)
	if err != nil {
		return "", err
	}
	downloadUrl := appDetail["download_url"].(string)
	if downloadUrl == "" {
		return "", errors.New("invalid download url")
	}
	savePath, err := DownloadFile(downloadUrl, saveDir, a.ProxyUrl, a.Timeout)
	if err != nil {
		return "", err
	}
	return savePath, nil
}

func (a *ApkpureDownloader) CheckUpdate(appid string, localVersion string) (bool, string, error) {

	return false, "", nil
}

func (a *ApkpureDownloader) Validate(appid string, version string) error {

	return nil
}

func (a *ApkpureDownloader) ParseGenericAppDetail(appDetailJson gjson.Result) GenericAppDetail {
	genericAppDetail := make(GenericAppDetail)

	fields := []string{
		"title", "label", "icon_url", "package_name", "version_code", "version_name",
		"review_stars", "description", "description_short", "whatsnew", "asset_usability",
		"developer", "is_show_comment_score", "comment_score1", "comment_score2", "comment_score3",
		"comment_score4", "comment_score5", "comment_total", "comment_score_total",
		"comment_score_stars", "price", "in_app_products", "introduction", "category_name",
		"is_free", "download_count", "version_id", "app_id", "real_package_name",
		"sdk_version", "target_sdk_version", "apk_type",
	}

	for _, field := range fields {
		if v := appDetailJson.Get(field); v.Exists() {
			genericAppDetail[field] = v.Value()
		}
	}

	// 特殊字段处理
	var tags []string
	for _, tag := range appDetailJson.Get("tags.#.name").Array() {
		tags = append(tags, tag.String())
	}
	genericAppDetail["tags"] = tags

	var sign []string
	for _, s := range appDetailJson.Get("sign").Array() {
		sign = append(sign, s.String())
	}
	genericAppDetail["sign"] = sign

	var nativeCode []string
	for _, n := range appDetailJson.Get("native_code").Array() {
		nativeCode = append(nativeCode, n.String())
	}
	genericAppDetail["native_code"] = nativeCode

	if t := appDetailJson.Get("update_date").String(); t != "" {
		if parsed, err := time.Parse(time.RFC3339, t); err == nil {
			genericAppDetail["update_date"] = parsed
		}
	}
	if t := appDetailJson.Get("create_date").String(); t != "" {
		if parsed, err := time.Parse(time.RFC3339, t); err == nil {
			genericAppDetail["create_date"] = parsed
		}
	}

	// 下载信息
	genericAppDetail["sha1"] = appDetailJson.Get("asset.sha1").String()
	genericAppDetail["size"] = appDetailJson.Get("asset.size").String()
	genericAppDetail["download_url"] = appDetailJson.Get("asset.url").String()
	genericAppDetail["download_type"] = appDetailJson.Get("asset.type").String()
	genericAppDetail["created_at"] = time.Now()

	return genericAppDetail
}

func isValidURL(rawURL string) bool {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}
	// 检查 Scheme 和 Host 是否存在
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}
	return true
}
