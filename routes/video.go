package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/liucxer/resource-manage/protocol"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// VideoCreate godoc
// @Summary 创建视频
// @Tags 视频
// @Security ApiKeyAuth
// @Accept multipart/form-data
// @Param video	formData file true	"待上传视频"
// @Produce application/json
// @Success 200 {object} protocol.VideoCreateReply
// @Router /resource-manage/v1/videos [post]
func VideoCreate(c *gin.Context) {
	var (
		reply protocol.VideoCreateReply
		err   error
	)

	contentType := c.ContentType()
	logrus.Warnf("contentType:%s", contentType)
	// 获取上传文件
	video, err := c.FormFile("video")
	if err != nil {
		logrus.Errorf("c.FormFile err:%v", err)
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}
	fileItems := strings.Split(video.Filename, ".")
	if len(fileItems) != 2 {
		logrus.Errorf("c.FormFile err:%v", err)
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	fileExt := fileItems[1]
	uuidStr := uuid.NewString()
	videoPath := "/var/video/" + uuidStr + "." + fileExt

	// 上传文件到指定的路径
	err = c.SaveUploadedFile(video, videoPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	reply.Url = "/video/" + uuidStr + "/" + uuidStr + ".m3u8"
	c.JSON(http.StatusOK, reply)
}
