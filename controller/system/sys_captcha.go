package system

// import (
// 	"back-end/global"
// 	systemRes "back-end/model/common/response"
// 	"github.com/gin-gonic/gin"
// 	"github.com/mojocn/base64Captcha"
// 	// "go.uber.org/zap"
// )

// var store = base64Captcha.DefaultMemStore

// func (b *BaseApi) Captcha(c *gin.Context) {
// 	// 生成默认数字
// 	driver := base64Captcha.NewDriverDigit(global.Global_Config.Captcha.ImgHeight, global.Global_Config.Captcha.ImgWidth, global.Global_Config.Captcha.KeyLong, 0.7, 80)
// 	cp := base64Captcha.NewCaptcha(driver, store)
// 	if id, b64s, _, err := cp.Generate(); err != nil {
// 		// global.MAY_LOGGER.Error("验证码获取失败", zap.Error(err))
// 		systemRes.FailWithMessage("验证码获取失败", c)
// 	} else {
// 		systemRes.OkWithDetailed(systemRes.SysCaptchaResponse{
// 			CaptchaId:     id,
// 			PicPath:       b64s,
// 			CaptchaLength: global.Global_Config.Captcha.KeyLong,
// 		}, "验证码获取成功", c)
// 	}
// }
