package dorm

import (
	// "back-end/model/test/expense"
	// "back-end/model/test/repair"
	// "back-end/model/test/student"
	// "fmt"
	"time"
	// "gorm.io/gorm"
)

// 宿舍楼	Dorms      []Dorm `json:"dorms"`
type Floor struct {
	Id         uint   `json:"id" gorm:"primarykey"`
	FloorsName string `json:"floorsName" gorm:"size:256"`
	FloorsType string `json:"floorsType" gorm:"size:256"`
	DormAmount uint   `json:"dormAmount" gorm:"size:256"`
	Dorms      []Dorm `json:"dormList,omitempty"`
}

// 宿舍
type Dorm struct {
	Id         uint       `json:"id" gorm:"primarykey"`
	DormNumber string     `json:"dormNumber"  gorm:"size:256"`
	Img        string     `json:"img" gorm:"size:256"`
	Capacity   int        `json:"Capacity"`
	FloorId    uint       `json:"floorId" ` //宿舍楼主键
	Floor      Floor      `json:"-"`
	FloorsName string     `json:"floorsName" gorm:"->"`
	Count      int64      `json:"count" gorm:"->"`
	StudInfos  []StudInfo `json:"studInfoList"`
	Rates      []Rate     `json:"rateList"`
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
	StayCause   string   `json:"stayCause" gorm:"size:256"`
	Opinions    string   `json:"opinions" gorm:"default:审核中;size:256"`
	Dorm        Dorm     `json:"dorm"`
	DormId      uint     `json:"dormId"`
}

// 学生信息
type StudInfo struct {
	Id            uint   `json:"id" gorm:"primarykey"`
	StudentName   string `json:"studentName" gorm:"size:256"`
	StudentNumber string `json:"studentNumber" gorm:"size:256"`
	Sex           string `json:"sex" gorm:"size:256"`
	Phone         string `json:"phone" gorm:"size:256"`
	Dorm          Dorm   `json:"dorm"`
	DormId        uint   `json:"dormId"`
}
type StayTime struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

// 评分
type Rate struct {
	Id         uint      `json:"id" gorm:"primarykey"`
	RateDate   time.Time `json:"rateDate"`
	Bed        uint      `json:"bedRate"`
	Ground     uint      `json:"groundRate"`
	Lavatory   uint      `json:"lavatory"`
	Goods      uint      `json:"goods"`
	TotalScore uint      `json:"totalScore"`
	Rater      string    `json:"rater"  gorm:"size:256"`
	Evaluation string    `json:"evaluation" gorm:"size:256"`
	Dorm       Dorm      `json:"dorm"`
	DormId     uint      `json:"dormId"`
}

type Expense struct {
	Id                uint      `json:"id" gorm:"primarykey"`
	PaymentTime       time.Time `json:"paymentTime" gorm:"type:date"`
	WaterCharge       float64   `json:"waterCharge"`
	ElectricityCharge float64   `json:"electricityCharge"`
	TotalCost         float64   `json:"totalCost"`
	Accountant        string    `json:"accountant" gorm:"size:256"`
	Phone             string    `json:"phone" gorm:"size:256"`
	Dorm              Dorm      `json:"dorm"`
	DormId            uint      `json:"dormId"`
}

// FloorsName string `json:"floorsName" gorm:"->"`
// DormNumber string `json:"dormNumber" gorm:"->"`
