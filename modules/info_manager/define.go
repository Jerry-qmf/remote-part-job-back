package info_manager

type ListParam struct {
	Page     int `form:"page,default=1"`
	PageSize int `form:"page_size,default=10"`
}

type ListJobShowInfoResp struct {
	Total int           `json:"total"`
	Data  []JobShowInfo `json:"data"`
}

type JobShowInfo struct {
	JobId    string `json:"job_id"`
	JobTitle string `json:"job_title"`
	JobPay   string `json:"job_pay"`
	JobLabel string `json:"job_label"`
}

type JobDetail struct {
	JobId           string    `json:"job_id"`
	JobTitle        string    `json:"job_title"`
	JobPay          string    `json:"job_pay"`
	JobLabel        string    `json:"job_label"`
	JobDescribe     string    `json:"job_describe"`
	JobCarouselList []string  `json:"job_carousel_list"`
	HRContact       HRContact `json:"HR_contact"`
}

type HRContact struct {
	WechatUrl string `json:"wechat_url"`
	WechatNum string `json:"wechat_num"`
}
