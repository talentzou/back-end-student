package apidorm

import "time"

//
type Floors_api struct {
	Id         string `json:"id" gorm:"size:256;"`
	FloorsName string `json:"floorsName" gorm:"size:256"`
	Floors     uint   `json:"floors" gorm:"size:256"`
	FloorsType string `json:"floorsType" gorm:"size:256"`
	DormAmount uint   `json:"dormAmount" gorm:"size:256"`
}

// 床位表格
type Bed_api struct {
	Id          string `json:"id" gorm:"size:256"`
	BedStatus   string `json:"bedStatus" gorm:"size:256"`
	DormNumber  string `json:"dormNumber" gorm:"size:256"`
	BedNumber   uint   `json:"bedNumber" gorm:"size:256"`
	Remark      string `json:"remark" gorm:"size:256"`
	StudentName string `json:"studentName" gorm:"size:256"`
}

// 评分
type Rate_api struct {
	Id         string    `json:"id" gorm:"size:256"`
	RateDate   time.Time `json:"rateDate" gorm:"size:256"`
	FloorsName string    `json:"floorsName" gorm:"size:256"`
	DormNumber uint      `json:"dormNumber" gorm:"size:256"`
	BedRate    uint      `json:"bedRate" gorm:"size:256"`
	GroundRate uint      `json:"groundRate" gorm:"size:256"`
	Lavatory   uint      `json:"lavatory" gorm:"size:256"`
	Goods      uint      `json:"goods" gorm:"size:256"`
	TotalScore uint      `json:"totalScore" gorm:"size:256"`
	Rater      string    `json:"rater" gorm:"size:256"`
	Evaluation string    `json:"evaluation" gorm:"size:256"`
	Remark     string    `json:"remark" gorm:"size:256"`
}
type Stay_api struct {
	StayDate    time.Time `json:"stayDate" gorm:"size:256"`
	StudentName string    `json:"studentName" gorm:"size:256"`
	FloorsName  string    `json:"floorsName" gorm:"size:256"`
	DormNumber  string    `json:"dormNumber" gorm:"size:256"`
	StayCause   string    `json:"stayCause " gorm:"size:256"`
	Instructor  string    `json:"instructor" gorm:"size:256"`
}
type Dorm_api struct {
	Id         string `json:"id" gorm:"size:256"`
	FloorsName string `json:"floorsName" gorm:"size:256"`
	DormNumber string `json:"dormNumber" gorm:"size:256"`
	Img        string `json:"img" gorm:"size:256"`
	DormType   string `json:"dormType" gorm:"size:256"`
	DormStatus string `json:"dormStatus" gorm:"size:256"`
}
