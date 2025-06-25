package adslot

import "time"

// ScrapyUnityadsCoreParams undefined
type ScrapyUnityadsCoreParams struct {
	ID            int64     `json:"id" gorm:"id"`
	ParamsId      int64     `json:"params_id" gorm:"params_id"`             // 对应泛化参数表中的主键id
	ParamsDict    string    `json:"params_dict" gorm:"params_dict"`         // 参数字典
	ParamsDictMd5 string    `json:"params_dict_md5" gorm:"params_dict_md5"` // 参数字典的md5值，主要用于去重
	IsAvailable   int8      `json:"is_available" gorm:"is_available"`       // 0：不可用，1: 可用
	Os            uint8     `json:"os" gorm:"os"`                           // 设备类型1:ios,2:android 3:pc
	Geo           string    `json:"geo" gorm:"geo"`                         // 可以泛化的代理地区
	SourceApp     string    `json:"source_app" gorm:"source_app"`           // 来源的app名称
	CreatedAt     time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"updated_at"`
	BkInt         int64     `json:"bk_int" gorm:"bk_int"`       // int 类型备用字段
	BkString      string    `json:"bk_string" gorm:"bk_string"` // string类型备用字段
	Lang          string    `json:"lang" gorm:"lang"`           // 国家语言代码
}

// TableName 表名称
func (s *ScrapyUnityadsCoreParams) TableName() string {
	return "scrapy_unityads_core_params"
}

func (s *ScrapyUnityadsCoreParams) Platform() string {
	return "unity"
}

func (s *ScrapyUnityadsCoreParams) SetParamsDictMd5(md5 string) {
	s.ParamsDictMd5 = md5
}

func (s *ScrapyUnityadsCoreParams) GetParamsDictMd5() (string, error) {
	if s.ParamsDictMd5 != "" {
		return s.ParamsDictMd5, nil
	}
	return "", nil
}

func (s *ScrapyUnityadsCoreParams) ValidateCoreParams() bool {
	s.IsAvailable = 1
	return true
}

func (s *ScrapyUnityadsCoreParams) GetAdUnitPushLog() *AdSlotPushLog {
	return &AdSlotPushLog{
		Platform: s.Platform(),
		CommonAdSlotCoreParams: CommonAdSlotCoreParams{
			ParamsDict:    s.ParamsDict,
			ParamsDictMd5: s.ParamsDictMd5,
			Os:            s.Os,
			IsAvailable:   s.IsAvailable,
			SourceApp:     s.SourceApp,
		},
	}
}
