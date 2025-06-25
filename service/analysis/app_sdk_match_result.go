package analysis

import (
	aGloabl "app-insight/global"
	"app-insight/model/analysis"
)

type AppSdkMatchResultService struct{}

// CreateAppSdkMatchResult 创建SDK匹配结果记录
// Author [yourname](https://github.com/yourname)
func (appSdkMatchResultService *AppSdkMatchResultService) CreateAppSdkMatchResult(appSdkMatchResult *analysis.AppSdkMatchResult) (err error) {
	err = aGloabl.AppRankDB.Create(appSdkMatchResult).Error
	return err
}

// DeleteAppSdkMatchResult 删除SDK匹配结果记录
// Author [yourname](https://github.com/yourname)
func (appSdkMatchResultService *AppSdkMatchResultService) DeleteAppSdkMatchResult(id string) (err error) {
	err = aGloabl.AppRankDB.Delete(&analysis.AppSdkMatchResult{}, "id = ?", id).Error
	return err
}

// DeleteAppSdkMatchResultByIds 批量删除SDK匹配结果记录
// Author [yourname](https://github.com/yourname)
func (appSdkMatchResultService *AppSdkMatchResultService) DeleteAppSdkMatchResultByIds(ids []string) (err error) {
	err = aGloabl.AppRankDB.Delete(&[]analysis.AppSdkMatchResult{}, "id in ?", ids).Error
	return err
}

// UpdateAppSdkMatchResult 更新SDK匹配结果记录
// Author [yourname](https://github.com/yourname)
func (appSdkMatchResultService *AppSdkMatchResultService) UpdateAppSdkMatchResult(appSdkMatchResult analysis.AppSdkMatchResult) (err error) {
	err = aGloabl.AppRankDB.Model(&analysis.AppSdkMatchResult{}).Where("id = ?", appSdkMatchResult.ID).Updates(&appSdkMatchResult).Error
	return err
}

// GetAppSdkMatchResult 根据id获取SDK匹配结果记录
// Author [yourname](https://github.com/yourname)
func (appSdkMatchResultService *AppSdkMatchResultService) GetAppSdkMatchResult(id string) (appSdkMatchResult analysis.AppSdkMatchResult, err error) {
	err = aGloabl.AppRankDB.Where("id = ?", id).First(&appSdkMatchResult).Error
	return
}

func (appSdkMatchResultService *AppSdkMatchResultService) BatchInsertAppSdkMatchResult(results []analysis.AppSdkMatchResult) (err error) {
	if len(results) == 0 {
		return nil // 避免空插入
	}
	err = aGloabl.AppRankDB.Create(&results).Error
	return err
}
