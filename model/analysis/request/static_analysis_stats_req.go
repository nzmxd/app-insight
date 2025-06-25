package request

import "time"

type StaticAnalysisStatsSearch struct {
	StartTime time.Time `json:"startTime" form:"startTime" time_format:"unix"`
	EndTime   time.Time `json:"endTime" form:"endTime"   time_format:"unix"`
}
