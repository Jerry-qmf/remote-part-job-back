package info_manager

import (
	"github.com/gin-gonic/gin"
	"remote-part-job-back/common/response"
	"remote-part-job-back/dao"
)

func InfoCarouselList(c *gin.Context) {
	var urlList []string
	carouselList, err := dao.GetCarouselInfoList()
	if err != nil {
		response.ErrorHandle(c, err, nil)
		return
	}
	for _, c := range carouselList {
		if !c.Deleted {
			urlList = append(urlList, c.CarouselUrl)
		}
	}
	response.RespondJSON(c, nil, urlList)
}
