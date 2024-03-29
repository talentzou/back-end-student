package apiexpense

type Expense_dorm struct {
	Id                     string  `json:"id" gorm:"size:256"`
	FloorsName             string  `json:"floorsName" gorm:"size:256"`
	DormNumber             string  `json:"dormNumber" gorm:"size:256"`
	PaymentTime            string  `json:"paymentTime" gorm:"type:date"`
	WaterConsumption       float64 `json:"waterConsumption" `
	WaterCharge            float64 `json:"waterCharge"`
	ElectricityConsumption float64 `json:"electricityConsumption" `
	ElectricityCharge      float64 `json:"electricityCharge"`
	TotalCost              float64 `json:"totalCost"`
	Accountant             string  `json:"accountant" gorm:"size:256"`
	Phone                  string  `json:"phone" gorm:"size:256"`
	Remark                 string  `json:"remark" gorm:"size:256"`
}
