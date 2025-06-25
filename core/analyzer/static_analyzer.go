package analyzer

import (
	"path/filepath"
	"strings"
)

type ScanRule struct {
	ID            int
	SdkName       string
	SdkMetadataID int
	RuleType      string   // class, permission, meta, xml, ...
	Patterns      []string // 匹配的字符串列表
	Key           string   // meta 匹配字段
	ExpectedVal   string   // meta 匹配值
}

type MatchResult struct {
	RuleID       int
	RuleType     string
	SdkInfoID    int
	SdkName      string
	Matched      bool
	MatchedItems []string
	Message      string
}

// StaticAnalyzer 分析应用sdk信息
type StaticAnalyzer interface {
	AddRule(rule ScanRule)
	LoadRules(rules []ScanRule)
	Analysis(localPath string) ([]MatchResult, error)
}

func sanitizeBaseName(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)

	// 去掉扩展名（支持 .apk, .xapk, .apks）
	base = strings.TrimSuffix(base, ext)

	// 替换空格 → 下划线
	base = strings.ReplaceAll(base, " ", "_")
	return base
}

// 匹配任意一个 pattern 即可返回 true
func matchPatterns(text string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.Contains(text, pattern) {
			return true
		}
	}
	return false
}
