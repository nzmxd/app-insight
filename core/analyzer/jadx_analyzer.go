package analyzer

import (
	"fmt"
	"github.com/nzmxd/bserver/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type JadxAnalyzer struct {
	WorkDir   string
	Exec      string
	scanRules []ScanRule
}

func (j *JadxAnalyzer) AddRule(rule ScanRule) {
	j.scanRules = append(j.scanRules, rule)
}

func (j *JadxAnalyzer) LoadRules(rules []ScanRule) {
	j.scanRules = rules
}

func (j *JadxAnalyzer) Analysis(localPath string) ([]MatchResult, error) {
	// 生成反编译输出路径
	jadxOutPath := filepath.Join(j.WorkDir, "jadx")
	absJadxOutPath, err := filepath.Abs(jadxOutPath)
	if err != nil {
		return nil, err
	}
	err = utils.CreateDir()
	if err != nil {
		return nil, err
	}
	baseName := sanitizeBaseName(localPath)
	datePath := time.Now().Format("2006-01-02")
	outDir := filepath.Join(absJadxOutPath, datePath, baseName)
	// 构造并执行 jadx 命令
	cmd := exec.Command(j.Exec, "-s", "-d", outDir, localPath)
	defer utils.DeLFile(outDir)
	_, err = cmd.CombinedOutput()
	if err != nil && !strings.Contains("exit status 1", err.Error()) {
		return nil, fmt.Errorf("error running jadx: %s", err)
	}
	xmlFiles, err := utils.FindFiles(outDir, "AndroidManifest.xml")
	if err != nil {
		return nil, err
	}
	results := j.matchXmlRules(xmlFiles)
	return results, nil
}

func (j *JadxAnalyzer) matchXmlRules(xmlFiles []string) []MatchResult {
	var results []MatchResult

	// 1. 提前过滤出 xml 类型的规则
	var xmlRules []ScanRule
	for _, rule := range j.scanRules {
		if rule.RuleType == "xml" {
			xmlRules = append(xmlRules, rule)
		}
	}

	// 2. 遍历每个 XML 文件
	for _, file := range xmlFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			continue // 可选：log.Printf("failed to read file %s: %v", file, err)
		}
		text := string(content)

		// 3. 遍历所有 XML 规则并匹配
		for _, rule := range xmlRules {
			if matchPatterns(text, rule.Patterns) {
				results = append(results, MatchResult{
					RuleID:    rule.ID,
					RuleType:  rule.RuleType,
					SdkInfoID: rule.SdkMetadataID,
					Matched:   true,
					SdkName:   rule.SdkName,
				})
			}
		}
	}

	return results
}
