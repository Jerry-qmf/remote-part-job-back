package dao

type JobInfo struct {
	Id              uint   `gorm:"primaryKey"`
	JobTitle        string `gorm:"column:job_title"`
	JobPay          string `gorm:"column:job_pay"`
	JobLabel        string `gorm:"column:job_label"`
	JobDescribe     string `gorm:"column:job_describe"`
	JobCarouselList string `gorm:"column:job_carousel_list"`
	WechatUrl       string `gorm:"column:wechat_url"`
	WechatNum       string `gorm:"column:wechat_num"`
	Expires         int    `gorm:"column:expires"`
	IsTop           bool   `gorm:"column:is_top"`
	Deleted         bool   `gorm:"column:deleted"`
	CreatedAt       int
	UpdatedAt       int
}

type CarouselInfo struct {
	Id          uint   `gorm:"primaryKey"`
	CarouselUrl string `gorm:"column:carousel_url"`
	Expires     int    `gorm:"column:expires"`
	Deleted     bool   `gorm:"column:deleted"`
	IsTop       bool   `gorm:"column:is_top"`
	CreatedAt   int
	UpdatedAt   int
}

func (JobInfo) TableName() string {
	return "job_info_tab"
}

func (CarouselInfo) TableName() string {
	return "carousel_info_tab"
}
