package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartChunk(c *gin.Context) {

	c.Status(http.StatusOK)
}
