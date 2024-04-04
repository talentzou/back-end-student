package repair

import "time"

type Repair struct {
	Id           uint      `json:"-" gorm:"primarykey"`
	FloorsName   string    `json:"floorsName" gorm:"size:256"`
	DormNumber   string    `json:"dormNumber" gorm:"size:256"`
	Problems     string    `json:"problems" gorm:"size:256"`
	SubmitDate   time.Time `json:"submitDate"`
	RepairStatus string    `json:"repairStatus" gorm:"size:256"`
	ReportMan    string    `json:"reportMan" gorm:"size:256"`
	Phone        string    `json:"phone" gorm:"size:256"`
	Repairer     string    `json:"repairer" gorm:"size:256"`
	Remark       string    `json:"remark" gorm:"size:256"`
}
