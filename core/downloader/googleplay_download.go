package downloader

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GooglePlayDownloader struct {
	ProxyUrl string
}

func (g *GooglePlayDownloader) GetAppDetail(appid string) (GenericAppDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GooglePlayDownloader) GetAppDownloadUrl(appid string, version string) string {
	//TODO implement me
	panic("implement me")
}

func (g *GooglePlayDownloader) ListVersions(appid string) (GenericAppVersionList, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GooglePlayDownloader) Download(appid string, version string, saveDir string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GooglePlayDownloader) CheckUpdate(appid string, localVersion string) (bool, string, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GooglePlayDownloader) Validate(appid string, version string) error {
	googlePlayerUrl := fmt.Sprintf("https://play.google.com/store/apps/details?id=%s", appid)
	method := "GET"
	client := &http.Client{}
	if g.ProxyUrl != "" {
		proxyURL, _ := url.Parse(g.ProxyUrl) // 你的代理地址
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}
	req, err := http.NewRequest(method, googlePlayerUrl, nil)

	if err != nil {
		return err
	}
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36")
	req.Header.Add("Cookie", "NID=524=pbfci_R7f3ZAnQll_VT5e-YKdY7wlBoFjNCK78-jd5EQPb_kZjL6X4HcypoBegD3UKdcE1outArTs34gSNgGtTS4LeM6jdRNhiRBP-kRqYkm_omhkfIOmlh5ICfn9w9q4_biM29AcpBMJsWSVEE_Ilf_5ASDOA9dMF83FOzf_3RghFouBWQiIQF1RZH9k2qg")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("package not fount")
	}
	return nil
}

func (g *GooglePlayDownloader) SetProxyUrl(proxyUrl string) error {
	if !isValidURL(proxyUrl) {
		return errors.New("invalid proxyUrl")
	}
	g.ProxyUrl = proxyUrl
	return nil
}
