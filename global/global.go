package global

import (
	"back-end/config"
	"gorm.io/gorm"
)
// $2a$10$8HFIPEmF/bTcloYrjeiRtuhERB05/WF7TtQOGxveB9c/cEu.VRmie
// $2a$10$WJlt5NDjyyFiY5jMzqyRKOfY9lyuiXWwrvlDZJ.Wu1Hju64zR.a8K
var (
	Global_Db             *gorm.DB
	// Global_Web_Route config.RouterConfig
	Global_Config         config.Server
)
