//apibasura/api/domain/entities
package entities

type Message struct {
	Text      string `json:"text"`
	Action    string `json:"action"`
}
func NewMessage(text string, action string) *Message {
	return &Message{
		Text: text,
		Action: action,

	}
}