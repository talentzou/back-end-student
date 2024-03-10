package apidorm

import "time"

//
type Floors_api struct {
	Id         string `json:"-" gorm:"size:256;"`
	FloorsName string `json:"floorsName" gorm:"size:256"`
	Floors     uint `json:"floors" gorm:"size:256"`
	FloorsType string `json:"floorsType" gorm:"size:256"`
	DormAmount uint   `json:"dormAmount" gorm:"size:256"`
}

// 床位表格
type Bed_api struct {
	BedStatus   string `json:"BedStatus" gorm:"size:256"`
	DormNumber  uint   `json:"DormNumber" gorm:"size:256"`
	BedNumber   uint   `json:"BedNumber" gorm:"size:256"`
	Message     string `json:"Message" gorm:"size:256"`
	StudentName string `json:"StudentName" gorm:"size:256"`
}

// 评分
type Rate_api struct {
	Id         string    `json:"Id" gorm:"size:256"`
	RateDate   time.Time `json:"RateDate" gorm:"size:256"`
	FloorsName string    `json:"FloorsName" gorm:"size:256"`
	DormNumber uint      `json:"DormNumber" gorm:"size:256"`
	BedRate    uint      `json:"BedRate" gorm:"size:256"`
	GroundRate uint      `json:"GroundRate" gorm:"size:256"`
	Lavatory   uint      `json:"Lavatory" gorm:"size:256"`
	Goods      uint      `json:"Goods" gorm:"size:256"`
	TotalScore uint      `json:"TotalScore" gorm:"size:256"`
	Rater      string    `json:"Rater" gorm:"size:256"`
	Evaluation string    `json:"Evaluation" gorm:"size:256"`
	Remark     string    `json:"Remark" gorm:"size:256"`
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
	FloorsName string `json:"floorsName" gorm:"size:256"`
	DormNumber uint   `json:"dormNumber" gorm:"size:256"`
	Img        []byte
	DormType   string `json:"dormType" gorm:"size:256"`
	DormStatus string `json:"dormStatus" gorm:"size:256"`
}
