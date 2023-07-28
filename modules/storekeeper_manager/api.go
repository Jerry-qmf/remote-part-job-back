package storekeeper_manager

import (
	"github.com/gin-gonic/gin"
)

//func Login(c *gin.Context) {
//	context := NewAppContext(c)
//	context.Run(
//		context.CheckCookie,
//		context.Login,
//		context.ParseLoginResp,
//	)
//}

func UpdateJob(c *gin.Context) {
	context := NewAppContext(c)
	context.Run(
		//context.AuthRequired,
		context.ParseUpdateProductReq, //获取参数
		context.GetMysqlProductInfo,
		context.UpdateMysqlProductInfo,
		context.ParseUpdateProductResp,
	)
}

func CreateJob(c *gin.Context) {
	context := NewAppContext(c)
	context.Run(
		//context.AuthRequired,
		context.ParseUpdateProductReq,  //获取参数
		context.GenMysqlJobInfo,        //生成一个空信息
		context.UpdateMysqlProductInfo, //更新到job info并刷盘
		context.ParseUpdateProductResp,
	)
}

func DeleteJob(c *gin.Context) {
	context := NewAppContext(c)
	context.Run(
		//context.AuthRequired,
		context.ParseDeleteProductReq, //获取参数，并从数据库获取product info
		context.DeleteMysqlProductInfo,
		context.ParseDeleteProductResp,
	)
}

func UpdateCarousel(c *gin.Context) {
	context := NewAppContext(c)
	context.Run(
		context.UpdateCarouselReq,
	)
}

//func ProductImageUpload(c *gin.Context) {
//	context := NewAppContext(c)
//	context.Run(
//		context.AuthRequired,
//		context.SaveImageAndSetUrl,
//	)
//}
