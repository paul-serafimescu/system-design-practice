package models

type MessageType int

const (
	ClientHello MessageType = iota
	// TODO: add others as we go
)

// base structure for any chat message
type ChatMessage struct {
	Type    MessageType            `json:"type"`
	Payload map[string]interface{} `json:"payload"`
}
