package apprank

import "time"

type AppBasicInfoAndroid struct {
	ID               int64     `json:"id" gorm:"id"`
	AppId            string    `json:"app_id" gorm:"app_id"`             // app唯一标识
	LogoCdnUrl       string    `json:"logo_cdn_url" gorm:"logo_cdn_url"` // logo cdn 地址唯一编码
	ReleaseDate      time.Time `json:"release_date" gorm:"release_date"` // 发布日期
	PublisherId      string    `json:"publisher_id" gorm:"publisher_id"` // app包名
	IsPaid           int8      `json:"is_paid" gorm:"is_paid"`           // 是否收费，0免费，1收费
	IsIap            int8      `json:"is_iap" gorm:"is_iap"`             // 是否内购，0非内购，1内购
	Categories       string    `json:"categories" gorm:"categories"`     // app类型，例[6014, 6017]
	StoreStatus      string    `json:"store_status" gorm:"store_status"` // 商店状态，保存各国家上下架信息 json：{\"1\":\"1\",\"2\":\"2\",\"3\":\"1\"}，索引为国家映射，0预约、1上架、2下架
	CreatedAt        time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"updated_at"`
	AppIdMd5         int64     `json:"app_id_md5" gorm:"app_id_md5"`               // app唯一标识md5前18位数值
	DeveloperWebsite string    `json:"developer_website" gorm:"developer_website"` // 开发者官网地址
	TagIds           string    `json:"tag_ids" gorm:"tag_ids"`                     // 标签id的list
	DeveloperEmail   string    `json:"developer_email" gorm:"developer_email"`     // app服务邮箱
	DeveloperAddress string    `json:"developer_address" gorm:"developer_address"` // 开发者地址
	PrivacyPolicy    string    `json:"privacy_policy" gorm:"privacy_policy"`       // app隐私政策链接
	CopyRight        string    `json:"copy_right" gorm:"copy_right"`               // app版权信息
}

// TableName 表名称
func (a AppBasicInfoAndroid) TableName() string {
	return "app_basic_info_android"
}

func (a AppBasicInfoAndroid) GetID() int64 {
	return a.ID
}
