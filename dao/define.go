package dao

type JobInfo struct {
	Id              uint   `gorm:"primaryKey"`
	JobTitle        string `json:"job_title"`
	JobPay          string `json:"job_pay"`
	JobLabel        string `json:"job_label"`
	JobDescribe     string `json:"job_describe"`
	JobCarouselList string `json:"job_carousel_list"`
	WechatUrl       string `json:"wechat_url"`
	WechatNum       string `json:"wechat_num"`
	Expires         int    `json:"expires"`
	IsTop           bool   `gorm:"column:is_top"`
	Deleted         bool   `gorm:"column:deleted"`
	CreatedAt       int
	UpdatedAt       int
}

func (JobInfo) TableName() string {
	return "job_info_tab"
}
