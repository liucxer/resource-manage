package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccessLimit(allowIPs []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(allowIPs) != 0 {
			clientIP := ctx.ClientIP()
			findIp := false

			for _, ip := range allowIPs {
				if ip == clientIP {
					findIp = true
				}
			}

			if !findIp {
				_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("access limit"))
				return
			}
		}

		ctx.Next()
	}
}
