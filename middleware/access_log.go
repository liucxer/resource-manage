package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

// LoggerAccessToFile 访问日志记录到文件
func LoggerAccessToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		referer := c.Request.Referer()
		// 请求host
		reqHost := c.Request.Host
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 请求IP
		clientIP := c.ClientIP()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 状态码
		statusCode := c.Writer.Status()

		// 日志格式
		logrus.WithFields(logrus.Fields{
			"req_time":    FormatTime(startTime, ""),
			"req_cost":    fmt.Sprintf("%0.2fms", float64(latencyTime)/float64(time.Millisecond)),
			"req_host":    reqHost,
			"req_uri":     reqUri,
			"req_method":  reqMethod,
			"status_code": statusCode,
			"client_ip":   clientIP,
			"referer":     referer,
		}).Info("runtime")
	}
}

const TimeFormatLayout = "2006-01-02 15:04:05"

// FormatTime formats time to string
func FormatTime(t time.Time, layout string) string {
	if layout == "" {
		layout = TimeFormatLayout
	}
	return t.Format(layout)
}
