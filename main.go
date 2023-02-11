package main

import (
	gins "github.com/gin-gonic/gin/ginS"
	"github.com/liucxer/resource-manage/config"
	_ "github.com/liucxer/resource-manage/docs"
	"github.com/liucxer/resource-manage/logger"
	_ "github.com/liucxer/resource-manage/routes"
	"github.com/sirupsen/logrus"
	"strconv"
)

func main() {
	var err error
	globalConfig, err := config.ReadConfig("./config.json")
	if err != nil {
		logrus.Warnf("config.ReadConfig err:%v, path:%s, use default config", err, "./config.json")
	}

	logger.InitLogger(globalConfig.LogPath, globalConfig.AppName)
	err = gins.Run("0.0.0.0:" + strconv.Itoa(int(globalConfig.ListeningPort)))
	if err != nil {
		logrus.Errorf("gins.Run err:%v", err)
		panic(err)
	}
}
