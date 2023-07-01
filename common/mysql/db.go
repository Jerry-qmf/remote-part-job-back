package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var Orm *DBOrm

// 使用继承，如果后面要实现方法重载的功能，需要改为组合机制
type DBOrm struct {
	*gorm.DB
}

// mysql配置
type Config struct {
	Driver          string
	DataBase        string
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	LogMode         string
}

func getLog(logMode string) logger.LogLevel {
	switch logMode {
	case "info":
		return logger.Info
	case "warn":
		return logger.Warn
	case "error":
		return logger.Error
	case "silent":
		return logger.Silent
	}
	return logger.Info
}

func New(c *Config) *DBOrm {
	// createDataBase(c) //如果数据库不存在，则尝试创建数据库

	//mysqlOrm, err := gorm.Open(mysql.Open(c.DSN), &gorm.Config{Logger: NewMysqlLogger(
	//	logger.Config{
	//		SlowThreshold: time.Second * 2,   // 慢 SQL 阈值
	//		LogLevel:      getLog(c.LogMode), // Log level
	//		Colorful:      false,             // 禁用彩色打印
	//	},
	//)})
	mysqlOrm, err := gorm.Open(mysql.Open(c.DSN))
	if err != nil {
		//logs.Error("mysql dsn(%v) error: %v", c.DSN, err)
		fmt.Printf("mysql dsn(%v) error: %v\n", c.DSN, err)
		panic(err)
	}
	orm := &DBOrm{mysqlOrm}
	//logs.Info("set max idle connection：%d", c.MaxIdleConns)
	fmt.Printf("set max idle connection：%d\n", c.MaxIdleConns)
	mysqlDB, err := orm.DB.DB()
	if err != nil {
		//logs.Error("mysql dsn(%v) error: %v", c.DSN, err)
		fmt.Printf("mysql dsn(%v) error: %v\n", c.DSN, err)
		panic(err)
	}
	mysqlDB.SetMaxIdleConns(c.MaxIdleConns)
	//logs.Info("set max open connection：%d", c.MaxOpenConns)
	fmt.Printf("set max open connection：%d\n", c.MaxOpenConns)
	mysqlDB.SetMaxOpenConns(c.MaxOpenConns)
	//logs.Info("set connection max life time：%d", c.ConnMaxLifetime)
	fmt.Printf("set connection max life time：%d\n", c.ConnMaxLifetime)
	mysqlDB.SetConnMaxLifetime(c.ConnMaxLifetime)
	if err = mysqlDB.Ping(); err != nil {
		//logs.Error(err.Error())
		fmt.Printf(err.Error() + "\n")
		panic(err)
	}
	return orm
}

func Init(c *Config) {
	Orm = New(c)
}

// todo 分页查询
// 分页器
type Pagination struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"limit"`
	TotalCount int64 `json:"total"`
}

// 分页数据
type PageData struct {
	Pagination
	Data interface{} `json:"items"`
}
