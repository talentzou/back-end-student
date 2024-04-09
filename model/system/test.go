package system

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	ID        int
	Name      string
	Addresses []Address `gorm:"many2many:person_addresses;"`
}

type Address struct {
	ID   uint
	Name string
}

type PersonAddress struct {
	PersonID  int `gorm:"primaryKey"`
	AddressID int `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}