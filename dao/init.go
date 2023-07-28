package dao

import (
	"remote-part-job-back/common/mysql"
	"remote-part-job-back/config"
	"time"
)

func Init() {
	initMysql()
	initJobInfo()
}

func initMysql() {
	cfg := config.ConfigHolder.DB
	mysql.Init(&mysql.Config{
		Driver:          cfg.Driver,
		DataBase:        cfg.DataBase,
		DSN:             cfg.DSN,
		MaxIdleConns:    cfg.MaxIdleConns,
		MaxOpenConns:    cfg.MaxOpenConns,
		ConnMaxLifetime: time.Duration(cfg.ConnMaxLifeTime) * time.Second,
		LogMode:         cfg.LogLevel,
	})
}
