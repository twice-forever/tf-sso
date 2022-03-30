package common

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}

type PageResponse struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}
