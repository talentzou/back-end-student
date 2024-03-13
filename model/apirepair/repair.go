package apirepair

type Repair_dorm struct {
	Id           string `json:"id" gorm:"size:256"`
	FloorsName   string `json:"floorsName" gorm:"size:256"`
	DormNumber   string `json:"dormNumber" gorm:"size:256"`
	Problems     string `json:"problems" gorm:"size:256"`
	SubmitDate   string `json:"submitDate" gorm:"type:date"`
	RepairStatus string `json:"repairStatus" gorm:"size:256"`
	ReportMan    string `json:"reportMan" gorm:"size:256"`
	Phone        string `json:"phone" gorm:"size:256"`
	Repairer     string `json:"repairer" gorm:"size:256"`
	Remark       string `json:"remark" gorm:"size:256"`
}
