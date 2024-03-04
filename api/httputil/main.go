package httputil

type HttpErrorBody struct {
	Error string `json:"error" example:"status bad request"`
}
