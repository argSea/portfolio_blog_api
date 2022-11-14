package service

type BaseResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Items   interface{} `json:"items"`
}
