package payload

const (
	MessageTypeError   = "error"
	MessageTypeSuccess = "success"
	Message            = "message"
)

type Response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func NewResponse(messageType, message string, data interface{}) Response {
	return Response{
		MessageType: messageType,
		Message:     message,
		Data:        data,
	}
}
