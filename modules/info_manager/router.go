package info_manager

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {
	infoRoute := router.Group("/remote_part_job/api/v1")
	infoRoute.GET("/job/list", InfoJobList)
	//infoRoute.GET("/store/info", InfoStore)
	//infoRoute.GET("/product/list", InfoProductList)
	infoRoute.GET("/job/detail", InfoJobDetail) //?product_id=2
	infoRoute.GET("/image/carousel", InfoCarouselList)
}
