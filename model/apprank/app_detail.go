package apprank

import "time"

type AppDetail struct {
	Title              string    `json:"title"`
	Label              string    `json:"label"`
	IconURL            string    `json:"icon_url"`
	PackageName        string    `json:"package_name"`
	VersionCode        string    `json:"version_code"`
	VersionName        string    `json:"version_name"`
	Sign               []string  `json:"sign" gorm:"type:Array(String)"`
	ReviewStars        float64   `json:"review_stars"`
	Description        string    `json:"description"`
	DescriptionShort   string    `json:"description_short"`
	Whatsnew           string    `json:"whatsnew"`
	AssetUsability     string    `json:"asset_usability"`
	Developer          string    `json:"developer"`
	IsShowCommentScore bool      `json:"is_show_comment_score"`
	CommentScore1      string    `json:"comment_score1"`
	CommentScore2      string    `json:"comment_score2"`
	CommentScore3      string    `json:"comment_score3"`
	CommentScore4      string    `json:"comment_score4"`
	CommentScore5      string    `json:"comment_score5"`
	CommentTotal       string    `json:"comment_total"`
	CommentScoreTotal  string    `json:"comment_score_total"`
	CommentScoreStars  float64   `json:"comment_score_stars"`
	Price              string    `json:"price"`
	InAppProducts      string    `json:"in_app_products"`
	Introduction       string    `json:"introduction"`
	CategoryName       string    `json:"category_name"`
	UpdateDate         time.Time `json:"update_date"`
	CreateDate         time.Time `json:"create_date"`
	IsFree             bool      `json:"is_free"`
	Tags               []string  `json:"tags"  gorm:"type:Array(String)"`
	SHA1               string    `json:"sha1"`
	Size               string    `json:"size"`
	IsGoogleListing    bool      `json:"is_google_listing"`
	DownloadURL        string    `json:"download_url"  gorm:"-"`
	DownloadType       string    `json:"download_type" gorm:"-"`
	DownloadCount      string    `json:"download_count"`
	VersionID          string    `json:"version_id"`
	AppID              string    `json:"app_id"`
	NativeCode         []string  `json:"native_code" gorm:"type:Array(String)"`
	APKType            uint8     `json:"apk_type"`
	RealPackageName    string    `json:"real_package_name"`
	SDKVersion         string    `json:"sdk_version"`
	TargetSDKVersion   string    `json:"target_sdk_version"`
	CreatedAt          time.Time `json:"created_at"`
}

func (AppDetail) TableName() string {
	return "android_app_detail"
}
