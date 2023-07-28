package storekeeper_manager

import (
	"github.com/gin-gonic/gin"
	"remote-part-job-back/common/response"
	"remote-part-job-back/dao"
	"time"
)

type AppContext struct {
	ginContext      *gin.Context
	storeId         string
	aborted         bool
	startCreateTime time.Time

	jobInfo          *dao.JobInfo
	updateJobInfoReq UpdateJobReq

	carouselList []string
}

func (a *AppContext) GetCarouselList() []string {
	return a.carouselList
}

func (a *AppContext) SetCarouselList(list []string) {
	a.carouselList = list
}

func (a *AppContext) UpdateJobReq() UpdateJobReq {
	return a.updateJobInfoReq
}

func (a *AppContext) SetUpdateJobReq(createJobReq UpdateJobReq) {
	a.updateJobInfoReq = createJobReq
}

func (a *AppContext) JobInfo() *dao.JobInfo {
	return a.jobInfo
}

func (a *AppContext) SetJobInfo(productInfo *dao.JobInfo) {
	a.jobInfo = productInfo
}

type ApolloHandler func()

func NewAppContext(g *gin.Context) *AppContext {
	a := AppContext{}
	a.SetGinContext(g)
	//a.SetLog(logs.GetCustomLog("rid", g.GetHeader("X-Request-Id"))) 不用request_id
	return &a
}

func (a *AppContext) ResponseWhenError(inErr error, responseErr error) {
	if inErr != nil && !a.Aborted() {
		response.ErrorHandle(a.GinContext(), responseErr, inErr)
		a.Abort()
	}
}

func (a *AppContext) Aborted() bool {
	return a.aborted
}

func (a *AppContext) Abort() {
	a.aborted = true
}

func (a *AppContext) GinContext() *gin.Context {
	return a.ginContext
}

func (a *AppContext) SetGinContext(ginContext *gin.Context) {
	a.ginContext = ginContext
}

func (a *AppContext) Run(handlers ...ApolloHandler) {
	for i := range handlers {
		handlers[i]()
		if a.Aborted() {
			return
		}
	}
}
