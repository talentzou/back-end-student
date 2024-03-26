package apidorm

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
	BedNumber   int    `json:"bedNumber" gorm:"size:256"`
	Remark      string `json:"remark" gorm:"size:256"`
	StudentName string `json:"studentName" gorm:"size:256"`
}

// 评分
type Rate_api struct {
	Id         string `json:"id" gorm:"size:256"`
	RateDate   string `json:"rateDate" gorm:"type:date"`
	FloorsName string `json:"floorsName" gorm:"size:256"`
	DormNumber string `json:"dormNumber" gorm:"size:256"`
	BedRate    uint   `json:"bedRate"`
	GroundRate uint   `json:"groundRate"`
	Lavatory   uint   `json:"lavatory"`
	Goods      uint   `json:"goods"`
	TotalScore uint   `json:"totalScore"`
	Rater      string `json:"rater"`
	Evaluation string `json:"evaluation" gorm:"size:256"`
	Remark     string `json:"remark" gorm:"size:256"`
}
type Stay_api struct {
	Id          string   `json:"id" gorm:"size:256"`
	StayTime    StayTime `json:"stayTime"  gorm:"embedded"`
	StudentName string   `json:"studentName" gorm:"size:256"`
	FloorsName  string   `json:"floorsName" gorm:"size:256"`
	DormNumber  string   `json:"dormNumber" gorm:"size:256"`
	StayCause   string   `json:"stayCause" gorm:"size:256"`
	Instructor  string   `json:"instructor" gorm:"size:256"`
	Opinions    string   `json:"opinions" gorm:"size:256"`
}

type StayTime struct {
	StartTime string `json:"startTime" gorm:"type:date"`
	EndTime   string `json:"endTime" gorm:"type:date"`
}
// 宿舍
type Dorm_api struct {
	Id           string `json:"id" gorm:"size:256"`
	FloorsName   string `json:"floorsName" gorm:"size:256"`
	DormNumber   string `json:"dormNumber" gorm:"size:256"`
	Img          string `json:"img" gorm:"size:256"`
	DormCapacity int    `json:"dormCapacity" gorm:"size:256"`
	DormStatus   string `json:"dormStatus" gorm:"size:256"`
}
