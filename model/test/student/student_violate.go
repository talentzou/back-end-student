package student

import "time"

type StudentViolate struct {
	Id            uint    `json:"-"  gorm:"primarykey"`
	StudentNumber string    `json:"studentNumber" gorm:"size:256"`
	StudentName   string    `json:"studentName" gorm:"size:256"`
	DormNumber    string    `json:"dormNumber" gorm:"size:256"`
	Violate       string    `json:"violate" gorm:"size:256"`
	Resolve       string    `json:"resolve" gorm:"size:256"`
	RecordDate    time.Time `json:"recordDate"`
}
