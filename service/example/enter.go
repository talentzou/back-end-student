package example

type ExampleGroup struct {
	FloorService
	DormService
	RateService
	StayService
	ExpenseService
	RepairService
	StudentService
}

var Example = new(ExampleGroup)
