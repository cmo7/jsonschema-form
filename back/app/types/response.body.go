package types

type StatusType string

const (
	Success StatusType = "success"
	Error   StatusType = "error"
)

type ResponseBody[T any] struct {
	Status  StatusType `json:"status"`
	Message string     `json:"message"`
	Data    T          `json:"data"`
}

func NewResponseBody[T interface{}](status StatusType, message string, data T) *ResponseBody[T] {
	return &ResponseBody[T]{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

type Empty struct{}
