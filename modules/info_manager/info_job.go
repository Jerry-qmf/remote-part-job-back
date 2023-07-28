package info_manager

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"remote-part-job-back/common"
	"remote-part-job-back/common/errors"
	"remote-part-job-back/common/response"
	"remote-part-job-back/dao"
	"strconv"
)

func InfoJobList(c *gin.Context) {
	//分页展示
	var lp ListParam
	if err := c.ShouldBindQuery(&lp); err != nil {
		response.ErrorHandle(c, errors.InvalidParamsError(err), nil)
	}

	allInfoList := dao.GetJobInfoList()
	//if err != nil {
	//	fmt.Printf("get job list error, err=%s \n", err.Error())
	//	response.ErrorHandle(c, errors.ForbiddenError(err), nil)
	//	return
	//}

	key := c.Query("key")
	if key != "" {
		allInfoList = allInfoList.FilterKey(key)
	}

	if len(allInfoList) == 0 {
		response.ErrorHandle(c, errors.ForbiddenError(fmt.Errorf("no jobs")), nil)
		return
	}

	fmt.Println("all jobs: total=", len(allInfoList), lp.Page, lp.PageSize)
	if (lp.Page-1)*lp.PageSize >= len(allInfoList) || lp.Page < 1 || lp.PageSize < 1 {
		err := fmt.Errorf("page is invalid")
		response.ErrorHandle(c, errors.InvalidParamsError(err), nil)
		return
	}

	resp := ListJobShowInfoResp{
		Total: len(allInfoList) / lp.PageSize,
		Data:  make([]JobShowInfo, 0),
	}
	if len(allInfoList)%lp.PageSize > 0 {
		resp.Total += 1
	}
	upLimit := lp.Page * lp.PageSize
	if len(allInfoList) < lp.Page*lp.PageSize {
		upLimit = len(allInfoList)
	}
	for i := (lp.Page - 1) * lp.PageSize; i < upLimit; i++ {
		info := allInfoList[i]
		resp.Data = append(resp.Data, parseStoreInfoToResp(info))
	}
	c.JSON(http.StatusOK, resp)
	//response.RespondJSON(c, nil, resp)
}

func parseStoreInfoToResp(info dao.JobInfo) JobShowInfo {
	return JobShowInfo{
		JobId:    strconv.Itoa(int(info.Id)),
		JobTitle: info.JobTitle,
		JobPay:   info.JobPay,
		JobLabel: info.JobLabel,
	}
}

func InfoJobDetail(c *gin.Context) {
	id := c.Query("job_id")
	if id == "" {
		err := fmt.Errorf("job_id为空")
		response.ErrorHandle(c, errors.ForbiddenError(err), nil)
		return
	}

	info, err := dao.GetJobInfoByJobId(id)
	if err != nil {
		fmt.Println(err)
		response.ErrorHandle(c, errors.ForbiddenError(err), nil)
		return
	}

	resp := JobDetail{
		JobId:           strconv.Itoa(int(info.Id)),
		JobTitle:        info.JobTitle,
		JobPay:          info.JobPay,
		JobLabel:        info.JobLabel,
		JobDescribe:     info.JobDescribe,
		JobCarouselList: common.SplitImgUrls(info.JobCarouselList),
		WechatNum:       info.WechatNum,
		WechatUrl:       info.WechatUrl,
	}
	response.RespondJSON(c, nil, resp)
}
