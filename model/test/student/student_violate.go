package student

import (
	"back-end/model/test/dorm"
	"time"
)

type StudentViolate struct {
	Id         uint          `json:"id"  gorm:"primarykey"`
	Violate    string        `json:"violate" gorm:"size:256"`
	Resolve    string        `json:"resolve" gorm:"size:256"`
	RecordDate time.Time     `json:"recordDate"`
	StudInfoId uint          `json:"studInfoId"`
	DormId     uint          `json:"dormId"`
	StudInfo   dorm.StudInfo `json:"studInfo"`
	Dorm       dorm.Dorm     `json:"dorm"`
}
