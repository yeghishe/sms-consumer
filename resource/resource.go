package resource

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yeghishe/sms-consumer/model"
)

type Event struct{}

func (*Event) Create(c *gin.Context) {
	var event model.Event
	if c.BindJSON(&event) == nil {
		log.Printf("Received event: %#v", event)
	}
}
