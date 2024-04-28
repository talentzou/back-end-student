package repair

import (
	"back-end/model/test/dorm"
	"time"
)

type Repair struct {
	Id           uint      `json:"id" gorm:"primarykey"`
	Problems     string    `json:"problems" gorm:"size:256"`
	RepairStatus string    `json:"repairStatus" gorm:"size:256;default:未完成"`
	ReportMan    string    `json:"reportMan" gorm:"size:256"`
	Phone        string    `json:"phone" gorm:"size:256"`
	Repairer     string    `json:"repairer" gorm:"size:256;default:无"`
	SubmitDate   time.Time `json:"submitDate"`
	FinishDate   time.Time `json:"finishDate" gorm:"default:null"`
	Dorm         dorm.Dorm `json:"dorm"`
	DormId       uint      `json:"dormId"`
}
