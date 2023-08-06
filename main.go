package main

import (
	gins "github.com/gin-gonic/gin/ginS"
	"github.com/liucxer/resource-manage/config"
	_ "github.com/liucxer/resource-manage/docs"
	"github.com/liucxer/resource-manage/logger"
	"github.com/liucxer/resource-manage/routes"
	_ "github.com/liucxer/resource-manage/routes"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func main() {
	var err error

	config.G_GlobalConfig, err = config.ReadConfig(os.Args[1])
	if err != nil {
		logrus.Warnf("config.ReadConfig err:%v, path:%s, use default config", err, "./config.json")
	}

	logger.InitLogger(config.G_GlobalConfig.LogPath, config.G_GlobalConfig.AppName)
	routes.InitRouter(config.G_GlobalConfig.ResourcePath)
	routes.InitResourcePath()
	err = gins.Run("0.0.0.0:" + strconv.Itoa(int(config.G_GlobalConfig.ListeningPort)))
	if err != nil {
		logrus.Errorf("gins.Run err:%v", err)
		panic(err)
	}
}
