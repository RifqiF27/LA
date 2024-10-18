package models

type Item interface {
	AddMenu(Food *Food)
	SearchMenu(name string) error
	DeleteMenu(name string) error
	UpdateMenu(name, kinds string, price float64, qty int, status bool) error
}
