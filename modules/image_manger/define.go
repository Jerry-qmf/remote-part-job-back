package image_manger

type ImageDownloadReq struct {
	StoreId   string `form:"store_id"`
	ProductId string `form:"product_id"`
	ImageId   string `form:"image_id"`
	Avatar    bool   `form:"2-avatar.jpg.jpg"` //是不是头像
	WeChat    bool   `form:"wechat"`           //是不是微信号
	Carousel  bool   `form:"carousel"`
}

type ImageDownloadResp struct {
}
