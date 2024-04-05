package dorm

import (
	"back-end/model/test/expense"
	"back-end/model/test/repair"
	"back-end/model/test/student"
	// "fmt"
	"time"
	// "gorm.io/gorm"
)

// 宿舍楼
type Floor struct {
	Id         uint   `json:"id" gorm:"primarykey"`
	FloorsName string `json:"floorsName" gorm:"size:256"`
	Floors     uint   `json:"floors" gorm:"size:256"`
	FloorsType string `json:"floorsType" gorm:"size:256"`
	DormAmount uint   `json:"dormAmount" gorm:"size:256"`
	Dorms      []Dorm `json:"dorms" gorm:"foreignKey:FloorsName;references:FloorsName"` 
}

// 宿舍
type Dorm struct {
	Id           uint              `json:"id" gorm:"primarykey"`
	DormNumber   string            `json:"dormNumber"  gorm:"size:256"`
	Img          string            `json:"img" gorm:"size:256"`
	DormCapacity int               `json:"dormCapacity" gorm:"size:256"`
	DormStatus   string            `json:"dormStatus" gorm:"size:256"`
	FloorsName   string            `json:"floorsName" gorm:"size:256"`
	Beds         []Bed             `json:"beds" gorm:"foreignKey:DormNumber;references:DormNumber"`
	Expenses     []expense.Expense `json:"expenses" gorm:"foreignKey:DormNumber;references:DormNumber"`
	Repairs      []repair.Repair   `json:"repairs" gorm:"foreignKey:DormNumber;references:DormNumber"`
	Rates        []Rate            `json:"rates" gorm:"foreignKey:DormNumber;references:DormNumber"`
}

// 床位
type Bed struct {
	Id          uint     `json:"id" gorm:"primarykey"`
	BedStatus   string   `json:"bedStatus" gorm:"size:256"`
	DormNumber  string   `json:"dormNumber"`
	BedNumber   int      `json:"bedNumber" gorm:"size:256"`
	Remark      string   `json:"remark" gorm:"size:256"`
	StudentName string   `json:"studentName" gorm:"size:256"`
	StudInfo    StudInfo `json:"studInfo" gorm:"foreignKey:StudentName;references:StudentName"` //属于某个学生
}

// 留宿
type Stay struct {
	Id          uint     `json:"id" gorm:"primarykey"`
	StayTime    StayTime `json:"stayTime"  gorm:"embedded"`
	StudentName string   `json:"studentName" gorm:"size:256"`
	FloorsName  string   `json:"floorsName" gorm:"size:256"`
	DormNumber  string   `json:"dormNumber" gorm:"size:256"`
	StayCause   string   `json:"stayCause" gorm:"size:256"`
	Opinions    string   `json:"opinions" gorm:"default:审核中;size:256"`
}

// 学生信息
type StudInfo struct {
	Id              uint                     `json:"id" gorm:"primarykey"`
	StudentName     string                   `json:"studentName" gorm:"size:256"`
	StudentNumber   string                   `json:"studentNumber" gorm:"size:256"`
	Sex             string                   `json:"sex" gorm:"size:256"`
	Major           string                   `json:"major" gorm:"size:256"`
	Phone           string                   `json:"phone" gorm:"size:256"`
	DormNumber      string                   `json:"dormNumber" gorm:"size:256"`
	StudentViolates []student.StudentViolate `json:"StudentViolates" gorm:"foreignKey:StudentName;references:StudentName"` //拥有多张违纪
	Stays           []Stay                   `json:"StudInfo" gorm:"foreignKey:StudentName;references:StudentName"`        //拥有多张留宿
}
type StayTime struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

// 评分
type Rate struct {
	Id         uint      `json:"id" gorm:"primarykey"`
	RateDate   time.Time `json:"rateDate"`
	FloorsName string    `json:"floorsName" gorm:"size:256"`
	DormNumber string    `json:"dormNumber" gorm:"size:256"`
	BedRate    uint      `json:"bedRate"`
	GroundRate uint      `json:"groundRate"`
	Lavatory   uint      `json:"lavatory"`
	Goods      uint      `json:"goods"`
	TotalScore uint      `json:"totalScore"`
	Rater      string    `json:"rater"`
	Evaluation string    `json:"evaluation" gorm:"size:256"`
	Remark     string    `json:"remark" gorm:"size:256"`
}



