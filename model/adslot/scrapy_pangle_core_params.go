package adslot

import "time"

// ScrapyPangleCoreParams TODO 通用参数不适用于Pangle渠道
type ScrapyPangleCoreParams struct {
	ID            int64     `json:"id" gorm:"id"`
	AppId         int64     `json:"app_id" gorm:"app_id"`
	CodeId        int64     `json:"code_id" gorm:"code_id"`
	Cypher        int8      `json:"cypher" gorm:"cypher"`
	IsAvailable   int64     `json:"is_available" gorm:"is_available"`
	Os            int64     `json:"os" gorm:"os"`
	PackageName   string    `json:"package_name" gorm:"package_name"`
	SlotType      int8      `json:"slot_type" gorm:"slot_type"`
	SlotParamDict string    `json:"slot_param_dict" gorm:"slot_param_dict"`
	SdkVersion    string    `json:"sdk_version" gorm:"sdk_version"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"updated_at"`
	CreatedAt     time.Time `json:"created_at" gorm:"created_at"`
	BkInt1        int64     `json:"bk_int1" gorm:"bk_int1"`
	BkString1     string    `json:"bk_string1" gorm:"bk_string1"`
	BkString2     string    `json:"bk_string2" gorm:"bk_string2"`
	BkInt2        int64     `json:"bk_int2" gorm:"bk_int2"`
}

// TableName 表名称
func (s *ScrapyPangleCoreParams) TableName() string {
	return "scrapy_pangle_core_params"
}

func (s *ScrapyPangleCoreParams) Platform() string {
	return "pangle"
}

//
//func (s *ScrapyPangleCoreParams) SetParamsDictMd5(md5 string) {
//	s.ParamsDictMd5 = md5
//}
//
//func (s *ScrapyPangleCoreParams) GetParamsDictMd5() (string, error) {
//	if s.ParamsDictMd5 != "" {
//		return s.ParamsDictMd5, nil
//	}
//	return "", nil
//}
//
//func (s *ScrapyPangleCoreParams) ValidateCoreParams() bool {
//	s.IsAvaliable = 1
//	return true
//}
//
//func (s *ScrapyPangleCoreParams) GetAdUnitPushLog() *AdSlotPushLog {
//	return &AdSlotPushLog{
//		Platform: s.Platform(),
//		CommonAdSlotCoreParams: CommonAdSlotCoreParams{
//			ParamsDict:    s.ParamsDict,
//			ParamsDictMd5: s.ParamsDictMd5,
//			Os:            s.Os,
//			IsAvailable:   s.IsAvaliable,
//			SourceApp:     s.SourceApp,
//		},
//	}
//}
