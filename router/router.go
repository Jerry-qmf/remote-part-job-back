package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"jerry_something/config"
	"jerry_something/modules/info_manager"
	"net/http"
)

var _router *gin.Engine

func InitRouter() {
	_router = gin.New()
	_router.Use(gin.LoggerWithFormatter(customLogFormat))
	_router.Use(gin.Recovery())
	_router.GET("/remote_part_job", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome to 远程零工.")
	})

	info_manager.InitRouter(_router)
	//image_manger.InitRouter(_router)
	//config_manager.InitRouter(_router)
	//storekeeper_manager.InitRouter(_router)
}

func GetRouter() *gin.Engine {
	return _router
}

func customLogFormat(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s|%s|%d|%s|%s|%s|%s\n",
		param.ClientIP,
		param.Method,
		param.StatusCode,
		param.Latency,
		param.Path,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func Run(ctx context.Context) {
	srv := http.Server{Addr: config.ConfigHolder.ListenAddr, Handler: _router}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//logs.Fatal("listen server error: %+v", err)
		}
	}()
	//logs.Info("jerry-something Server Started: listening at http://%s", config_manager.ConfigHolder.ListenAddr)
	<-ctx.Done()
	//logs.Info("Stopping uss-data-summary...")
}
