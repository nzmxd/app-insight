package adslot

import "time"

// ScrapyIronsourceCoreParams undefined
type ScrapyIronsourceCoreParams struct {
	ID            int64     `json:"id" gorm:"id"`
	ParamsId      int64     `json:"params_id" gorm:"params_id"`             // 对sp_etl.app_sdk_info中的主键id
	ParamsDict    string    `json:"params_dict" gorm:"params_dict"`         // 参数字典
	ParamsDictMd5 string    `json:"params_dict_md5" gorm:"params_dict_md5"` // 参数字典的md5值，主要用于去重
	IsAvaliable   int8      `json:"is_avaliable" gorm:"is_avaliable"`       // 0：不可用，1: 可用
	Os            uint8     `json:"os" gorm:"os"`                           // 设备类型1:ios,2:android 3:pc
	Geo           string    `json:"geo" gorm:"geo"`                         // 国家地区编码
	SourceApp     string    `json:"source_app" gorm:"source_app"`           // 来源的app名称
	Lang          string    `json:"lang" gorm:"lang"`                       // 国家语言代码
	CreatedAt     time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"updated_at"`
	BkInt         int64     `json:"bk_int" gorm:"bk_int"`       // int 类型备用字段
	BkString      string    `json:"bk_string" gorm:"bk_string"` // string类型备用字段
}

// TableName 表名称
func (s *ScrapyIronsourceCoreParams) TableName() string {
	return "scrapy_ironsource_core_params"
}

func (s *ScrapyIronsourceCoreParams) Platform() string {
	return "ironsource"
}

func (s *ScrapyIronsourceCoreParams) SetParamsDictMd5(md5 string) {
	s.ParamsDictMd5 = md5
}

func (s *ScrapyIronsourceCoreParams) GetParamsDictMd5() (string, error) {
	if s.ParamsDictMd5 != "" {
		return s.ParamsDictMd5, nil
	}
	return "", nil
}

func (s *ScrapyIronsourceCoreParams) ValidateCoreParams() bool {
	s.IsAvaliable = 1
	return true
}

func (s *ScrapyIronsourceCoreParams) GetAdUnitPushLog() *AdSlotPushLog {
	return &AdSlotPushLog{
		Platform: s.Platform(),
		CommonAdSlotCoreParams: CommonAdSlotCoreParams{
			ParamsDict:    s.ParamsDict,
			ParamsDictMd5: s.ParamsDictMd5,
			Os:            s.Os,
			IsAvailable:   s.IsAvaliable,
			SourceApp:     s.SourceApp,
		},
	}
}
