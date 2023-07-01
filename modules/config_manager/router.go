package config_manager

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {
	infoRoute := router.Group("/api/config")
	infoRoute.GET("/update", ConfigUpdate)
	infoRoute.GET("/get", ConfigGet)
}
