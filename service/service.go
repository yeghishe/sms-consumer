package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yeghishe/sms-consumer/resource"
)

func Run() {
	var (
		event  resource.Event
		status resource.Status
	)

	r := gin.Default()
	r.GET("/status", status.Get)
	r.POST("/event", event.Create)

	r.Run(":5000")
}
