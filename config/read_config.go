package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type GlobalConfig struct {
	LogPath       string `json:"log_path"`
	ListeningPort int64  `json:"listening_port"`
	// 上传文件的资源路径
	ResourcePath string   `json:"resource_path"`
	AppName      string   `json:"app_name"`
	LimitSources []string `json:"limit_sources"`
}

func ReadConfig(configFile string) (GlobalConfig, error) {
	var (
		err          error
		globalConfig GlobalConfig
	)
	bts, err := ioutil.ReadFile(configFile)
	if err != nil {
		return DefaultGlobalConfig(), err
	}

	err = json.Unmarshal(bts, &globalConfig)
	if err != nil {
		return GlobalConfig{}, err
	}
	if globalConfig.ListeningPort == 0 {
		globalConfig.ListeningPort = 80
	}
	if globalConfig.ResourcePath == "" {
		globalConfig.ResourcePath = "./resource"
	}
	if globalConfig.LogPath == "" {
		globalConfig.LogPath = "./log"
	}
	if globalConfig.AppName == "" {
		globalConfig.AppName = os.Args[0]
	}
	return globalConfig, err
}

func DefaultGlobalConfig() GlobalConfig {
	var globalConfig GlobalConfig
	globalConfig.ListeningPort = 80
	globalConfig.ResourcePath = "./resource"
	globalConfig.LogPath = "./log"
	globalConfig.AppName = os.Args[0]
	return globalConfig
}
