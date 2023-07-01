package image_manger

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {
	infoRoute := router.Group("/remote_part_job/api/v1")
	infoRoute.POST("/image/upload", ImageUpload)
	infoRoute.GET("/image/download", ImageDownload)
}
