package apistudent

type StudInfo_model struct {
	Id            string `json:"id" gorm:"size:256"`
	StudentName   string `json:"studentName" gorm:"size:256"`
	StudentNumber string `json:"studentNumber" gorm:"size:256"`
	Sex           string `json:"sex" gorm:"size:256"`
	Major         string `json:"major" gorm:"size:256"`
	Phone         string `json:"phone" gorm:"size:256"`
	DormNumber    string `json:"dormNumber" gorm:"size:256"`
}
