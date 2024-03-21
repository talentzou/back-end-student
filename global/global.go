package global

import (
	"back-end/config"
	"gorm.io/gorm"
)

var (
	Global_Db             *gorm.DB
	Global_Font_End_Route config.RouterConfig
)
