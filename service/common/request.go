package common

// Paginate 分页操作
type Paginate struct {
	Page     int `qurty:"page"`
	PageSize int `qurty:"page_size"`
}
