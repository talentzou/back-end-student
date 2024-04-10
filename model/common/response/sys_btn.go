package response
import "back-end/model/system"

type SysBtnsResponse struct {
	Btns []system.SysBtn `json:"btns"`
}