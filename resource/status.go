package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Status struct{}

func (*Status) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
