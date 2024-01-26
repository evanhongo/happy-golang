package api

type ErrorCode string

const (
	NOT_FOUND             ErrorCode = "NOT_FOUND"
	INVALID_REQUEST       ErrorCode = "INVALID_REQUEST"
	INTERNAL_SERVER_ERROR ErrorCode = "INTERNAL_SERVER_ERROR"
)

type HttpErrorBody struct {
	Code  string `json:"code" example:"NOT_FOUND"`
	Error string `json:"error" example:"status bad request"`
}
