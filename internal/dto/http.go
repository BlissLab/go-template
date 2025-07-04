package dto

const (
	SUCCESS     = 2000
	SUCCESS_MSG = "ok"
	FAILED      = 10000
	FAILED_MSG  = "failed"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Page struct {
	Page  int `json:"page"`
	Total int `json:"total"`
	Size  int `json:"size"`
}

type PageRequest struct {
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
}

type PageResponse[T any] struct {
	List []T `json:"list"`
	Page `json:"pagination"`
}
