package example

type ExampleGroup struct {
	FloorService
	DormService
	RateService
	StayService
	ExpenseService
	RepairService
	StudentService
	ViolateService
}

var Example = new(ExampleGroup)
