package image_manger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"remote-part-job-back/common"
	"remote-part-job-back/common/errors"
	"remote-part-job-back/common/response"
)

func ImageDownload(c *gin.Context) {
	fileName := c.Query("image")
	path := common.GetPathByFileName(fileName)
	c.File(path)
}

func ImageUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		//errLog.Error(logrus.Fields{"err":err.Error(),"source":pkg.GetPath()},"controller - admin - upload")
		response.ErrorHandle(c, errors.InvalidParamsError(err), nil)
		return
	}

	imageName := c.Query("image_name")
	filePath, url := common.GenFilePathAndUrl(imageName)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		fmt.Println(c.Request.URL, err)
		response.ErrorHandle(c, errors.InternalDBError(err), nil)
	}

	fmt.Println(filePath, url)
	response.RespondJSON(c, nil, url)

	go common.ConvertImageSize(filePath)
}

