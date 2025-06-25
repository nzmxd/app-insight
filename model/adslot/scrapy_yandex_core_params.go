package adslot

import "time"

// ScrapyYandexCoreParams  TODO 通用参数不适用于Yandex渠道
type ScrapyYandexCoreParams struct {
	ID          int64     `json:"id" gorm:"id"` // 主键
	Name        string    `json:"name" gorm:"name"`
	NetworkName string    `json:"network_name" gorm:"network_name"`
	Type        string    `json:"type" gorm:"type"`
	AdUnitId    string    `json:"ad_unit_id" gorm:"ad_unit_id"`     // 广告位ID
	PackageName string    `json:"package_name" gorm:"package_name"` // 包名
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`     // 调度时间
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`     // 广告位创建时间
	BkInt1      int64     `json:"bk_int1" gorm:"bk_int1"`           // 状态码
	BkInt2      int64     `json:"bk_int2" gorm:"bk_int2"`           // 调度成功次数
	BkString1   string    `json:"bk_string1" gorm:"bk_string1"`     // 调度日期
	BkString2   string    `json:"bk_string2" gorm:"bk_string2"`
}

// TableName 表名称
func (*ScrapyYandexCoreParams) TableName() string {
	return "scrapy_yandex_core_params"
}

func (s *ScrapyYandexCoreParams) Platform() string {
	return "yandex"
}
