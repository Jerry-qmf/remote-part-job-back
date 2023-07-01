package info_manager

import (
	"github.com/gin-gonic/gin"
	"remote-part-job-back/common/response"
	"remote-part-job-back/config"
)

func InfoCarouselList(c *gin.Context) {
	urlList := config.ConfigHolder.CarouselInfo
	response.RespondJSON(c, nil, urlList)
}
