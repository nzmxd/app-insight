package request

import (
	"time"
)

type ScrapyApplovinCoreParamsRequest struct {
	ID            int64                  `json:"id" gorm:"id" form:"id"`
	ParamsId      int64                  `json:"params_id" gorm:"params_id" form:"params_id"`                   // 对应泛化参数表中的主键id
	ParamsDict    map[string]interface{} `json:"params_dict" gorm:"params_dict" form:"params_dict"`             // 参数字典
	ParamsDictMd5 string                 `json:"params_dict_md5" gorm:"params_dict_md5" form:"params_dict_md5"` // 参数字典的md5值，主要用于去重
	IsAvailable   int8                   `json:"is_available" gorm:"is_available" form:"is_available"`          // 0：不可用，1: 可用
	Os            uint8                  `json:"os" gorm:"os" form:"os"`                                        // 设备类型1:ios,2:android 3:pc
	SourceApp     string                 `json:"source_app" gorm:"source_app" form:"source_app"`                // 来源的app名称
	CreatedAt     time.Time              `json:"created_at" gorm:"created_at" form:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at" gorm:"updated_at" form:"updated_at"`
	BkInt         int64                  `json:"bk_int" gorm:"bk_int" form:"bk_int"`          // int 类型备用字段
	BkString      string                 `json:"bk_string" gorm:"bk_string" form:"bk_string"` // string类型备用字段
	Geo           string                 `json:"geo" gorm:"geo" form:"geo"`                   // 相关国家编码
	Lang          string                 `json:"lang" gorm:"lang" form:"lang"`                // 国家语言代码
	Version       int64                  `json:"version" gorm:"version" form:"version"`
	IsPublished   uint8                  `json:"is_published" gorm:"is_published" form:"is_published"`
}
