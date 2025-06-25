package analysis

import "app-insight/service/download"

type ServiceGroup struct {
	AppAnalysisTaskService
	AppSdkMatchRuleService
	AppSdkMatchResultService
	SdkMetadataService
	AppStaticAnalysisTaskService
}

var appAnalysisTaskService AppAnalysisTaskService

// var appSdkMatchRuleService AppSdkMatchRuleService
var appSdkMatchResultService AppSdkMatchResultService

// var sdkMetadataService SdkMetadataService
// var appStaticAnalysisTaskService AppStaticAnalysisTaskService
var downloadTaskService download.DownloadTaskService
