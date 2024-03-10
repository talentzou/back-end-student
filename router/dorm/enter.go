package dorm

import (
	"github.com/gin-gonic/gin"
)

type DormGroup struct {
	Dorm
	Bed
	Floor
	Rate
	Stay
}

func (i *DormGroup) UseDormRouter(d *gin.RouterGroup) {
	dorm := d.Group("/Dormitory")
	{
		i.Api_Bed(dorm)
		i.Api_Dorm(dorm)
		i.Api_Stay(dorm)
		i.Api_Rate(dorm)
		i.Api_Floor(dorm)
	}

}
