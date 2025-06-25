package adslot

type AdSlotCoreParams interface {
	GetParamsDictMd5() (string, error) // 计算广告位md5
	SetParamsDictMd5(md5 string)       // 设置广告位md5
	ValidateCoreParams() bool          // 验证广告位是否有效
	GetAdUnitPushLog() *AdSlotPushLog  //
}

type CommonAdSlotCoreParams struct {
	ParamsDict    string `json:"params_dict" gorm:"params_dict"`                 // 参数字典
	ParamsDictMd5 string `json:"params_dict_md5" gorm:"params_dict_md5"`         // 参数字典的md5值，主要用于去重
	Os            uint8  `json:"os" gorm:"os"`                                   // 设备类型1:ios,2:android 3:pc
	IsAvailable   int8   `json:"is_available" gorm:"is_available"`               // 广告位状态，0-无效。1-有效
	SourceApp     string `json:"source_app" gorm:"source_app" form:"source_app"` // 来源的app名称
}
