package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"remote-part-job-back/config"
	"remote-part-job-back/dao"
	"remote-part-job-back/router"
	"sync"
	"syscall"
)
import "flag"
import "path/filepath"

var configPath string

func parseArgs() {
	flag.StringVar(&configPath, "config_manager", filepath.Join("config.json"), "config_manager file path")
	flag.Parse()
}

func main() {
	parseArgs()
	config.Init(configPath)
	router.InitRouter()
	dao.Init()
	//logs
	//包容gin logs

	fmt.Println("hello ", config.ConfigHolder.ListenAddr)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gracefulExitHandle()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		router.Run(ctx)
	}()

	wg.Wait()
}

func gracefulExitHandle() {
	chExit := make(chan os.Signal, 1)
	signal.Notify(chExit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-chExit
		//todo 程序优雅退出处理
		//logs.Info("received signal %v, exiting...", s)
		os.Exit(1)
	}()
}
