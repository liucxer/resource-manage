package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/liucxer/resource-manage/protocol"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// PictureCreate godoc
// @Summary 创建图片
// @Tags 图片
// @Security ApiKeyAuth
// @Accept multipart/form-data
// @Param picture	formData file true	"待上传图片"
// @Produce application/json
// @Success 200 {object} protocol.PictureCreateReply
// @Router /resource-manage/v1/pictures [post]
func PictureCreate(c *gin.Context) {
	var (
		reply protocol.PictureCreateReply
		err   error
	)

	contentType := c.ContentType()
	logrus.Warnf("contentType:%s", contentType)
	// 获取上传文件
	object, err := c.FormFile("picture")
	if err != nil {
		logrus.Errorf("c.FormFile err:%v", err)
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	fileItems := strings.Split(object.Filename, ".")
	if len(fileItems) != 2 {
		logrus.Errorf("c.FormFile err:%v", err)
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	fileExt := fileItems[1]
	uuidStr := uuid.NewString()
	picturePath := "/var/picture/" + uuidStr + "." + fileExt

	// 上传文件到指定的路径
	err = c.SaveUploadedFile(object, picturePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	reply.Url = "/picture/" + uuidStr + "/" + uuidStr + "." + fileExt
	c.JSON(http.StatusOK, reply)
}
