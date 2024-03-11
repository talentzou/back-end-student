package request

type PageInfo struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page" form:"page"`         // 页码
	PageSize int         `json:"pageSize" form:"pageSize"` // 每页大小
}

// type CustomContext struct {
// 	*gin.Context
// }

// func (c *CustomContext) HandlerModel(s interface{}) gin.HandlerFunc {

// 	var tempArr []apidorm.Floors_api
// 	err := c.ShouldBindJSON(&floors)
// 	if err != nil {
// 		return func(c *gin.Context) {
// 			// This will handle the response when the handler is invoked
// 			response.FailWithMessage("系统错误", c)
// 		}
// 	}
// 	switch c.Request.Method {
// 	case "GET":

// 		return Dorm_api.QueryApi
// 	case "POST":
// 		return Dorm_api.CreateApi
// 	case "DELETE":
// 		return Dorm_api.DeleteApi
// 	case "PUT":
// 		return Dorm_api.UpdateApi
// 	}

// 	return Dorm_api.UpdateApi
// }
