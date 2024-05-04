package global

import (
	"back-end/config"
	"gorm.io/gorm"
)

var (
	Global_Db             *gorm.DB
	// Global_Web_Route config.RouterConfig
	Global_Config         config.Server
)
