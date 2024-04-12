package response
import "back-end/model/system"

type SysBtnsResponse struct {
	Btns []system.SysBtn `json:"btns"`
}
type SysBtnsTree struct {
	BtnTree []system.MenuTree `json:"btnTree"`
}