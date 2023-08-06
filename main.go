package main

import (
	"fmt"
	gins "github.com/gin-gonic/gin/ginS"
	"github.com/liucxer/resource-manage/config"
	_ "github.com/liucxer/resource-manage/docs"
	"github.com/liucxer/resource-manage/logger"
	"github.com/liucxer/resource-manage/mgrs"
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
	routes.InitResourcePath(config.G_GlobalConfig.ResourcePath)

	if len(os.Args) >= 3 {
		if os.Args[2] == "upload" {
			dirPath := os.Args[3]
			res, err := mgrs.GlobalMgr.Upload(dirPath)
			if err != nil {
				logrus.Errorf("mgrs.GlobalMgr.Upload err:%v", err)
				panic(err)
			}
			for _, item := range res {
				str := fmt.Sprintf("<lazy-component><video controls playsinline poster=\"%s\" preload=\"metadata\"  src=\"%s\" ></video></lazy-component>",
					item.Picture, item.Video)
				fmt.Println(str)
			}
		}
	} else {
		routes.InitRouter(config.G_GlobalConfig.ResourcePath)
		err = gins.Run("0.0.0.0:" + strconv.Itoa(int(config.G_GlobalConfig.ListeningPort)))
		if err != nil {
			logrus.Errorf("gins.Run err:%v", err)
			panic(err)
		}
	}
}
