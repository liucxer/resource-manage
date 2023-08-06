package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/liucxer/resource-manage/config"
	"net/http"
)

func AccessLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		allowHosts := config.G_GlobalConfig.LimitHosts
		if config.G_GlobalConfig.EnableLimitHost && len(allowHosts) != 0 {
			requestHost := ctx.Request.Referer()
			findHost := false
			for _, host := range allowHosts {
				if host == requestHost {
					findHost = true
				}
			}

			if !findHost {
				_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("access limit"))
				return
			}
		}

		ctx.Next()
	}
}
