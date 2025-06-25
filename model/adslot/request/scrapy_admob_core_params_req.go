package request

import (
	"time"
)

// ScrapyAdmobCoreParamsRequest undefined
type ScrapyAdmobCoreParamsRequest struct {
	ID            int64                  `json:"id" gorm:"id"`
	ParamsId      int64                  `json:"params_id" gorm:"params_id"`             // 对应泛化参数表中的主键id
	ParamsDict    map[string]interface{} `json:"params_dict" gorm:"params_dict"`         // 参数字典
	ParamsDictMd5 string                 `json:"params_dict_md5" gorm:"params_dict_md5"` // 参数字典的md5值，主要用于去重
	IsAvailable   int8                   `json:"is_available" gorm:"is_available"`       // 0：不可用，1: 可用
	Os            uint8                  `json:"os" gorm:"os"`                           // 设备类型1:ios,2:android 3:pc
	SourceApp     string                 `json:"source_app" gorm:"source_app"`           // 来源的app名称
	AppName       string                 `json:"app_name" gorm:"app_name"`               // app的名称
	Geo           string                 `json:"geo" gorm:"geo"`                         // 国家信息
	Lang          string                 `json:"lang" gorm:"lang"`                       // 语言信息
	CreatedAt     time.Time              `json:"created_at" gorm:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at" gorm:"updated_at"`
	BkInt         int64                  `json:"bk_int" gorm:"bk_int"`                   // int 类型备用字段
	BkString      string                 `json:"bk_string" gorm:"bk_string"`             // string类型备用字段
	Format        string                 `json:"format" gorm:"format"`                   // 泛化需要的参数
	GeoList       string                 `json:"geo_list" gorm:"geo_list"`               // 可以抓取到广告的国家列表
	IsApp         int8                   `json:"is_app" gorm:"is_app"`                   // 抓取到数据中，是否为app
	SourceAppType int8                   `json:"source_app_type" gorm:"source_app_type"` // 0: 应用类型, 1: 益智游戏, 2: 探险游戏, 3: 音乐游戏, 4: 休闲游戏, 5: 卡牌游戏, 6: 动作游戏, 7: 策略游戏, 8: 百科游戏, 9: 街机游戏, 10: 文字游戏, 11: 风格化游戏, 12: 教育游戏, 13: 模拟游戏, 14: 角色扮演游戏, 15: 赌场游戏, 16: 运动游戏, 17:赛车游戏
	AdUnitQuality int64                  `json:"ad_unit_quality" gorm:"ad_unit_quality"` // 0: 默认值，没有进行质量筛选 1：低质量  2：中质量 3：高质量
}
