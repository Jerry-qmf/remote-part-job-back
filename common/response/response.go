package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"remote-part-job-back/common/errors"
	"runtime"
)

func obj(code errors.ErrCode, msg string, data interface{}) gin.H {
	return gin.H{
		"code":    code,
		"message": msg,
		"data":    data,
	}
}

func ErrorHandle(c *gin.Context, responseError error, data interface{}) {
	if responseError == nil {
		//logs.Error("params error")
		return
	}
	//logs.Error(responseError.Error())
	RespondJSON(c, responseError, data)
	call(2)
}

func call(skip int) {
	pc, file, line, _ := runtime.Caller(skip)
	pcName := runtime.FuncForPC(pc).Name() //获取函数名
	//fmt.Println(fmt.Sprintf("%v   %s   %d   %t   %s", pc, file, line, ok, pcName))
	fmt.Println(fmt.Sprintf("%s %d %s", file, line, pcName))
}

func RespondJSON(c *gin.Context, err error, data interface{}) {
	if err == nil {
		c.JSON(http.StatusOK, obj(errors.StatusOK, "ok", data))
		return
	}
	if e, ok := err.(*errors.Error); ok {
		errCode := e.ErrorCode()
		c.JSON(errCode.ToHTTPCode(), obj(errCode, e.Error(), data))
		return
	}
	c.JSON(http.StatusInternalServerError, obj(errors.StatusUnknown, err.Error(), data))
}
