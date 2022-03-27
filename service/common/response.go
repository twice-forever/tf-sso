package common

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type PageResponse struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}
