package analysis

import (
	"app-insight/core/analyzer"
	aGloabl "app-insight/global"
	"app-insight/model/analysis"
	"encoding/json"
	"fmt"
	"strings"
)

type AppSdkMatchRuleService struct{}

// CreateAppSdkMatchRule 创建sdkMatchRule表记录
// Author [yourname](https://github.com/yourname)
func (appSdkMatchRuleService *AppSdkMatchRuleService) CreateAppSdkMatchRule(appSdkMatchRule *analysis.AppSdkMatchRule) (err error) {
	err = aGloabl.AppRankDB.Create(appSdkMatchRule).Error
	return err
}

// DeleteAppSdkMatchRule 删除sdkMatchRule表记录
// Author [yourname](https://github.com/yourname)
func (appSdkMatchRuleService *AppSdkMatchRuleService) DeleteAppSdkMatchRule(id string) (err error) {
	err = aGloabl.AppRankDB.Delete(&analysis.AppSdkMatchRule{}, "id = ?", id).Error
	return err
}

// DeleteAppSdkMatchRuleByIds 批量删除sdkMatchRule表记录
// Author [yourname](https://github.com/yourname)
func (appSdkMatchRuleService *AppSdkMatchRuleService) DeleteAppSdkMatchRuleByIds(ids []string) (err error) {
	err = aGloabl.AppRankDB.Delete(&[]analysis.AppSdkMatchRule{}, "id in ?", ids).Error
	return err
}

// UpdateAppSdkMatchRule 更新sdkMatchRule表记录
func (appSdkMatchRuleService *AppSdkMatchRuleService) UpdateAppSdkMatchRule(appSdkMatchRule analysis.AppSdkMatchRule) (err error) {
	err = aGloabl.AppRankDB.Model(&analysis.AppSdkMatchRule{}).Where("id = ?", appSdkMatchRule.ID).Updates(&appSdkMatchRule).Error
	return err
}

// GetAppSdkMatchRule 根据id获取sdkMatchRule表记录
func (appSdkMatchRuleService *AppSdkMatchRuleService) GetAppSdkMatchRule(id string) (appSdkMatchRule analysis.AppSdkMatchRule, err error) {
	err = aGloabl.AppRankDB.Where("id = ?", id).Preload("SdkInfo").First(&appSdkMatchRule).Error
	return
}

func (appSdkMatchRuleService *AppSdkMatchRuleService) GetAllAppSdkMatchRule() ([]analyzer.ScanRule, error) {
	var matchRules []analysis.AppSdkMatchRule
	var sdkInfos []analysis.SdkMetadata
	var scanRules []analyzer.ScanRule

	// 1. 查询所有匹配规则
	if err := aGloabl.AppRankDB.Find(&matchRules).Error; err != nil {
		return nil, fmt.Errorf("查询 AppSdkMatchRule 失败: %w", err)
	}

	// 2. 查询所有 SDK 信息
	if err := aGloabl.AppRankDB.Find(&sdkInfos).Error; err != nil {
		return nil, fmt.Errorf("查询 SdkMetadata 失败: %w", err)
	}

	// 3. 构建 sdkInfoId → sdkName 映射（可选）
	sdkMap := make(map[int64]string)
	for _, sdk := range sdkInfos {
		sdkMap[sdk.ID] = sdk.SdkName
	}

	// 4. 转换成 []ScanRule
	for _, rule := range matchRules {
		sdkName := sdkMap[rule.SdkMetadataId]

		if strings.TrimSpace(rule.XmlScanRules) != "" {
			var patterns []string
			if err := json.Unmarshal([]byte(rule.XmlScanRules), &patterns); err != nil {
				return nil, fmt.Errorf("解析 xml_scan_rules 失败，RuleID=%d: %w", rule.ID, err)
			}
			scanRules = append(scanRules, analyzer.ScanRule{
				ID:            int(rule.ID),
				SdkMetadataID: int(rule.SdkMetadataId),
				SdkName:       sdkName,
				RuleType:      "xml",
				Patterns:      patterns,
			})
		}

		if strings.TrimSpace(rule.GlobalScanRules) != "" {
			var patterns []string
			if err := json.Unmarshal([]byte(rule.GlobalScanRules), &patterns); err != nil {
				return nil, fmt.Errorf("解析 global_scan_rules 失败，RuleID=%d: %w", rule.ID, err)
			}
			scanRules = append(scanRules, analyzer.ScanRule{
				ID:            int(rule.ID),
				SdkMetadataID: int(rule.SdkMetadataId),
				SdkName:       sdkName,
				RuleType:      "global",
				Patterns:      patterns,
			})
		}
	}

	return scanRules, nil
}
