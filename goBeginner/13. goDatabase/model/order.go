package model

type Order struct {
	ID uint16
	AdminId int
	CustomerId int
	DriverId int
	RegionAreaId int
	DateOrder string
	Status bool
}