package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/liucxer/resource-manage/logger"
	"github.com/liucxer/resource-manage/mgrs"
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
	videoPath := ResourcePath + "/video/" + uuidStr + "." + fileExt

	// 上传文件到指定的路径
	err = c.SaveUploadedFile(video, videoPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	reply.Url = "/video/" + uuidStr + "." + fileExt
	c.JSON(http.StatusOK, reply)
}

func InitResourcePath() {
	err := logger.NewFileMgr(ResourcePath, "").CheckAndCreateDir()
	if err != nil {
		panic(fmt.Sprintf("CheckAndCreateDir %s error. err:%v", ResourcePath, err))
	}
	err = logger.NewFileMgr(ResourcePath+"/video", "").CheckAndCreateDir()
	if err != nil {
		panic(fmt.Sprintf("CheckAndCreateDir %s error. err:%v", ResourcePath, err))
	}
	err = logger.NewFileMgr(ResourcePath+"/picture", "").CheckAndCreateDir()
	if err != nil {
		panic(fmt.Sprintf("CheckAndCreateDir %s error. err:%v", ResourcePath, err))
	}
}

// MultipartVideoCreate godoc
// @Summary 批量创建视频
// @Tags 视频
// @Security ApiKeyAuth
// @Accept multipart/form-data
// @Param video1	formData file false	"待上传视频"
// @Param video2	formData file false	"待上传视频"
// @Param video3	formData file false	"待上传视频"
// @Param video4	formData file false	"待上传视频"
// @Param video5	formData file false	"待上传视频"
// @Param video6	formData file false	"待上传视频"
// @Param video7	formData file false	"待上传视频"
// @Param video8	formData file false	"待上传视频"
// @Produce application/json
// @Success 200 {object} protocol.VideoCreateReply
// @Router /resource-manage/v1/videos/multipart [post]
func MultipartVideoCreate(c *gin.Context) {
	var (
		reply protocol.MultipartVideoCreateReply
		err   error
	)

	contentType := c.ContentType()
	logrus.Warnf("contentType:%s", contentType)
	// 获取上传文件
	video, err := c.FormFile("video1")
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
	videoPath := ResourcePath + "/video/" + uuidStr + "." + fileExt

	// 上传文件到指定的路径
	err = c.SaveUploadedFile(video, videoPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	// 封面获取
	coverPath := ResourcePath + "/picture/" + uuidStr + ".jpg"
	cmd := "ffmpeg -i " + videoPath + " -vframes 1 " + coverPath
	_, err = mgrs.GlobalMgr.SyncExecute(cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	reply.Items = append(reply.Items,
		protocol.Item{
			Picture: "/picture/" + uuidStr + ".jpg",
			Video:   "/video/" + uuidStr + "." + fileExt,
		})
	c.JSON(http.StatusOK, reply)
}
