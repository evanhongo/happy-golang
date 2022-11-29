package entity

type JobRequest struct {
	Name string      `json:"name" example:"compressImage"`
	Data interface{} `json:"data"`
}

type Job struct {
	Id     string      `json:"id" example:"e6e07f18-cae7-4ea9-a4ba-2c8ac364ea5b"`
	Name   string      `json:"name" example:"compressImage,omitempty"`
	State  string      `json:"state" example:"SUCCESS,omitempty"`
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
}
