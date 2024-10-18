package models

import "fmt"

type Food struct {
	Name   string
	Kinds  string
	Price  float64
	Qty    int
	Status bool
}
type Menu struct {
	Food []Food
}

func (m *Menu) AddMenu(f *Food) {
	m.Food = append(m.Food, *f)

}

func (m *Menu) SearchMenu(name string) (*Food, error) {
	for i := range m.Food {
		if m.Food[i].Name == name {
			return &m.Food[i], nil
		}

	}
	return nil, fmt.Errorf("Menu dengan nama '%s' tidak ditemukan", name)
}

func (m *Menu) DeleteMenu(name string) error {
	for i, f := range m.Food {
		if f.Name == name {
			m.Food = append(m.Food[:i], m.Food[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Menu dengan nama '%s' tidak ditemukan", name)

}

func (m *Menu) UpdateMenu(name, kinds string, price float64, qty int, status bool) error {
	m.SearchMenu(name)
	for i, f := range m.Food {
		if f.Name == name {
			m.Food[i].Name = name
			m.Food[i].Kinds = kinds
			m.Food[i].Price = price
			m.Food[i].Qty = qty
			m.Food[i].Status = status
			return nil
		}
	}
	return fmt.Errorf("Menu dengan nama '%s' tidak ditemukan", name)
}
