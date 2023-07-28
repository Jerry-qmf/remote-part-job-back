package storekeeper_manager

import (
	"fmt"
	"os"
	"remote-part-job-back/common"
	"remote-part-job-back/common/errors"
	"remote-part-job-back/common/response"
	"remote-part-job-back/dao"
	"strconv"
	"strings"
)

func (a *AppContext) ParseDeleteProductResp() {
	response.RespondJSON(a.GinContext(), nil, "success")
}

func (a *AppContext) DeleteMysqlProductInfo() {
	err := dao.DeleteJobInfo(strconv.Itoa(int(a.JobInfo().Id)))
	if err != nil {
		a.ResponseWhenError(err, errors.InternalDBError(err))
		return
	}

	//删除商品时删除图片文件
	for _, url := range append(common.SplitImgUrls(a.JobInfo().JobCarouselList), a.JobInfo().WechatNum) {
		list := strings.Split(url, "=")
		var name string
		if len(list) >= 2 {
			name = list[1]
		}
		_ = os.Remove(common.GetPathByFileName(name))
	}
}

func (a *AppContext) ParseDeleteProductReq() {
	var req DeleteProductReq
	if err := a.GinContext().ShouldBindJSON(&req); err != nil {
		a.ResponseWhenError(err, errors.InvalidParamsError(err))
		return
	}

	info, err := dao.GetJobInfoByJobId(req.ProductId)
	if err != nil || info == nil {
		err := fmt.Errorf("商品已不在")
		a.ResponseWhenError(err, errors.InvalidParamsError(err))
		return
	}
	a.SetJobInfo(info)
}

func (a *AppContext) ParseUpdateProductResp() {
	response.RespondJSON(a.GinContext(), nil, a.UpdateJobReq())
}

func (a *AppContext) UpdateMysqlProductInfo() {
	req := a.UpdateJobReq()
	info := a.JobInfo()

	info.JobTitle = req.JobTitle
	info.JobDescribe = req.JobDescribe
	info.JobPay = req.JobPay
	info.JobLabel = req.JobLabel
	info.JobCarouselList = common.MergeImgUrls(req.JobCarouselList)
	info.WechatUrl = req.WechatUrl
	info.WechatNum = req.WechatNum
	info.Expires = req.Expires
	info.IsTop = req.IsTop

	var err error
	if info.Id == 0 {
		var id uint
		id, err = dao.CreateJobInfo(info)
		req.Id = strconv.Itoa(int(id))
		a.SetUpdateJobReq(req)
	} else {
		err = dao.UpdateJobInfo(info)
	}
	if err != nil {
		a.ResponseWhenError(err, errors.InternalDBError(err))
		return
	}
}

func (a *AppContext) GenMysqlJobInfo() {
	info := &dao.JobInfo{}
	a.SetJobInfo(info)
}

func (a *AppContext) GetMysqlProductInfo() {
	info, err := dao.GetJobInfoByJobId(a.updateJobInfoReq.Id)
	if err != nil {
		err = fmt.Errorf("no this job: %s", err)
		a.ResponseWhenError(err, errors.ForbiddenError(err))
		return
	}
	a.SetJobInfo(info)
}

func (a *AppContext) ParseUpdateProductReq() {
	var req UpdateJobReq
	if err := a.GinContext().ShouldBindJSON(&req); err != nil {
		a.ResponseWhenError(err, errors.InvalidParamsError(err))
		return
	}
	if len(req.JobTitle) == 0 || len(req.JobTitle) > 64 {
		err := fmt.Errorf("工作名字不能为空，且长度不能大于64")
		fmt.Println("error: job_title=", req.JobTitle)
		a.ResponseWhenError(err, errors.InvalidParamsError(err))
		return
	}
	a.SetUpdateJobReq(req)
}

//
//func (a *AppContext) SaveImageAndSetUrl() {
//	file, err := a.GinContext().FormFile("file")
//	if err != nil {
//		fmt.Println(a.GinContext().Request.URL, err)
//		a.ResponseWhenError(err, errors.InvalidParamsError(err))
//		return
//	}
//
//	imageName := a.GinContext().Query("image_name")
//	path, url := common.GenFilePathAndUrl(imageName, a.GetStoreId())
//	err = a.GinContext().SaveUploadedFile(file, path)
//	if err != nil {
//		fmt.Println(a.GinContext().Request.URL, err)
//		a.ResponseWhenError(err, errors.InternalDBError(err))
//	}
//	fmt.Println(a.GetStoreId(), path, url)
//	response.RespondJSON(a.GinContext(), nil, url)
//
//	go common.ConvertImageSize(path)
//}

func (a *AppContext) UpdateCarouselReq() {
	var req UpdateCarouselReq
	if err := a.GinContext().ShouldBindJSON(&req); err != nil {
		a.ResponseWhenError(err, errors.InvalidParamsError(err))
		return
	}

	updateCarousel(req.CarouselList)
	response.RespondJSON(a.GinContext(), nil, nil)
}

func updateCarousel(list []string) {
	oldL, _ := dao.GetCarouselInfoList()
	for _, v := range oldL {
		dao.DeleteCarouselInfo(strconv.Itoa(int(v.Id)))
	}
	for _, v := range list {
		dao.CreateCarouselInfo(&dao.CarouselInfo{
			CarouselUrl: v,
		})
	}
}