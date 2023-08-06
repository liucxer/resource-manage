package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/liucxer/resource-manage/config"
	"github.com/liucxer/resource-manage/protocol"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
)

// SetLimitHost godoc
// @Summary 设置host访问限制配置
// @Tags 配置
// @Security ApiKeyAuth
// @Accept multipart/form-data
// @Produce application/json
// @Router /resource-manage/v1/limit-host [post]
func SetLimitHost(c *gin.Context) {
	var (
		args protocol.SetLimitHostArgs
		err  error
	)

	err = c.BindJSON(&args)
	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	config.G_GlobalConfig.EnableLimitHost = args.Enable
	config.G_GlobalConfig.LimitHosts = strings.Split(args.LimitHosts, ",")
	err = config.WriteConfig(os.Args[1], config.G_GlobalConfig)
	if err != nil {
		logrus.Warnf("config.WriteConfig err:%v, path:%s, use default config", err, "./config.json")
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	c.Status(http.StatusOK)
}

// GetLimitHost godoc
// @Summary 获取host访问限制配置
// @Tags 配置
// @Security ApiKeyAuth
// @Accept multipart/form-data
// @Produce application/json
// @Router /resource-manage/v1/limit-host [post]
func GetLimitHost(c *gin.Context) {
	var (
		reply protocol.GetLimitHostReply
	)

	reply.LimitHosts = config.G_GlobalConfig.LimitHosts.Str()
	reply.Enable = config.G_GlobalConfig.EnableLimitHost

	var amisReply AmisReply
	amisReply.Data = reply
	c.JSON(http.StatusOK, amisReply)
}

type AmisReply struct {
	Status int64       `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
