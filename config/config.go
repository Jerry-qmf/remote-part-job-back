package config

import (
	"encoding/json"
	"os"
	"sync"
)

type Config struct {
	Debug          bool     `json:"debug"`
	ListenAddr     string   `json:"listen_addr"`
	ServiceDomain  string   `json:"service_domain"`
	ImageRoot      string   `json:"image_root"`
	ImageSizeLimit int      `json:"image_size_limit"`
	CarouselInfo   []string `json:"carousel_info"`
	Log            Log      `json:"log"`
	DB             DB       `json:"db"`
}

type DB struct {
	Driver          string `json:"driver"`
	DataBase        string `json:"database"`
	DSN             string `json:"dsn"`
	MaxIdleConns    int    `json:"maxidleconns"`
	MaxOpenConns    int    `json:"maxopenconns"`
	ConnMaxLifeTime int    `json:"connmaxlifetime"`
	LogLevel        string `json:"log_level"`
}

type Log struct {
	LogLevel               string `json:"log_level"`
	LogFolder              string `json:"log_folder"`
	LogFile                string `json:"log_file"`
	MaxAge                 int    `json:"max_age"`
	RotationTime           int    `json:"rotation_time"`
	IsEnableRecordFileInfo bool   `json:"is_enable_record_file_info"`
}

var (
	configPath   string
	ConfigHolder Config
)

func Init(filePath string) {
	new(sync.Once).Do(func() {
		configPath = filePath
		if err := loadConfig(); err != nil {
			panic(err)
		}
	})
}

func Update(filePath string) error {
	if filePath != "" {
		configPath = filePath
	}
	if err := loadConfig(); err != nil {
		return err
	}
	return nil
}

func Get() Config {
	return ConfigHolder
}

func loadConfig() error {
	d, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(d, &ConfigHolder); err != nil {
		return err
	}
	return nil
}
