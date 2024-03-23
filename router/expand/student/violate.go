package student

import (
	"back-end/controller"
	"github.com/gin-gonic/gin"
)

type Violate struct{}

func (v *Violate) Api_violate(r *gin.RouterGroup) {
	violateApiRouter := r.Group("/violate")
	violateApi := controller.Vio_api
	{
		violateApiRouter.GET("getViolate/:Page/:PageSize", violateApi.QueryVioApi)
		violateApiRouter.POST("createViolate", violateApi.CreateVioApi)
		violateApiRouter.PUT("updateViolate", violateApi.UpdateVioApi)
		violateApiRouter.DELETE("deleteViolateById", violateApi.DeleteVioApi)
	}
}
