package apistudent

type Violate_model struct {
	Id            string `json:"id" gorm:"size:256"`
	StudentNumber string `json:"studentNumber" gorm:"size:256"`
	StudentName   string `json:"studentName" gorm:"size:256"`
	DormNumber    string `json:"dormNumber" gorm:"size:256"`
	Violate       string `json:"violate" gorm:"size:256"`
	Resolve       string `json:"resolve" gorm:"size:256"`
	RecordDate    string `json:"recordDate" gorm:"type:date"`
}
