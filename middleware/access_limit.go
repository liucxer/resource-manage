package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccessLimit(allowHosts []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(allowHosts) != 0 {

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
