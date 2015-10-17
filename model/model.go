package model

type Event struct {
	EventType string `json:"type" binding:"required"`
	Payload   string `json:"payload" binding:"required"`
}
