package storekeeper_manager

import (
	"github.com/gin-gonic/gin"
)

const SessionName = "session_id"

var TokenMap = map[string]string{}

func InitRouter(router *gin.Engine) {
	//infoRoute := router.Group("/api/v1/store/keeper")
	//infoRoute.POST("login", Login)
	//infoRoute.POST("logout", Logout)
	//infoRoute.POST("register", Register)
	//infoRoute.POST("/password/update", PasswordUpdate)

	// Private group, require authentication to access
	private := router.Group("remote_part_job/api/v1")
	{
		private.POST("/job/create", CreateJob)
		private.POST("/job/update", UpdateJob)
		private.POST("/job/delete", DeleteJob)
		//private.POST("/image/upload", ProductImageUpload)
		private.POST("/carousel/update", UpdateCarousel)
	}
}
