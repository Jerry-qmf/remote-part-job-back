package storekeeper_manager

type UserAppidPasswordLoginReq struct {
	Appid    string `form:"appid" json:"appid"`
	Password string `form:"password" json:"password"`
}

type UserAppidPasswordLogoutReq struct {
	Appid string `form:"appid" json:"appid"`
}

type UserPasswordUpdateReq struct {
	Appid       string `form:"appid" json:"appid"`
	Password    string `form:"password" json:"password"`
	NewPassword string `form:"new_password" json:"new_password"`
}

type UpdateJobReq struct {
	Id              string   `json:"job_id"`
	JobTitle        string   `json:"job_title"`
	JobPay          string   `json:"job_pay"`
	JobLabel        string   `json:"job_label"`
	JobDescribe     string   `json:"job_describe"`
	JobCarouselList []string `json:"job_carousel_list"`
	WechatUrl       string   `json:"wechat_url"`
	WechatNum       string   `json:"wechat_num"`
	Expires         int      `json:"expires"`
	IsTop           bool     `json:"is_top"`
	Deleted         bool     `json:"deleted"`
}

type DeleteProductReq struct {
	ProductId string `json:"job_id"`
}

type UpdateCarouselReq struct {
	CarouselList []string `json:"carousel_list"`
}
