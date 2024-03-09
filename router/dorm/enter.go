package dorm

import (
	"fmt"

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
	dorm := d.Group("/dorm")
	fmt.Println("wo执行到111111")
	{
		i.Api_Bed(dorm)
		i.Api_Dorm(dorm)
		i.Api_Stay(dorm)
		i.Api_Rate(dorm)
		i.Api_Floor(dorm)
	}

}
