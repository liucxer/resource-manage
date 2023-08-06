package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type LimitHostList []string

func (l LimitHostList) Str() string {
	var res string
	for i := 0; i < len(l); i++ {
		res += l[i]
		if i != len(l)-1 {
			res += ","
		}
	}
	return res
}

type GlobalConfig struct {
	LogPath       string `json:"log_path"`
	ListeningPort int64  `json:"listening_port"`
	// 上传文件的资源路径
	ResourcePath    string        `json:"resource_path"`
	WebPath         string        `json:"web_path"`
	AppName         string        `json:"app_name"`
	LimitHosts      LimitHostList `json:"limit_hosts"`
	EnableLimitHost bool          `json:"enable_limit_host"`
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

func WriteConfig(configFile string, config GlobalConfig) error {
	var (
		err error
	)

	bts, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configFile, bts, os.ModePerm)
	if err != nil {
		return err
	}

	return err
}

func DefaultGlobalConfig() GlobalConfig {
	var globalConfig GlobalConfig
	globalConfig.ListeningPort = 80
	globalConfig.ResourcePath = "./resource"
	globalConfig.LogPath = "./log"
	globalConfig.AppName = os.Args[0]
	return globalConfig
}

var (
	G_GlobalConfig = GlobalConfig{}
)
