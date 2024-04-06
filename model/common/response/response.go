package response
type PageInfo struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page" form:"page"`         // 页码
	PageSize int         `json:"pageSize" form:"pageSize"` // 每页大小
}