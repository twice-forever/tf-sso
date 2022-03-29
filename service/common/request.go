package common

// Paginate 分页操作
type Paginate struct {
	Page     int `qurty:"page"`
	PageSize int `qurty:"page_size"`
}

// Login 登录结构
type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
