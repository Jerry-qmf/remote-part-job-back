package config_manager

import (
	"github.com/gin-gonic/gin"
	"remote-part-job-back/common/response"
	"remote-part-job-back/config"
)

func ConfigUpdate(c *gin.Context) {
	err := config.Update(c.Query("path"))
	if err != nil {
		response.ErrorHandle(c, err, nil)
	}
	response.RespondJSON(c, nil, "更新成功")
}

func ConfigGet(c *gin.Context) {
	response.RespondJSON(c, nil, config.Get())
}
