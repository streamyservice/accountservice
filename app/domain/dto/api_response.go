package dto

type ApiResponse[T any] struct {
	ResponseKey     string `json:"status"`
	ResponseMessage string `json:"message"`
	Data            T      `json:"data"`
}
