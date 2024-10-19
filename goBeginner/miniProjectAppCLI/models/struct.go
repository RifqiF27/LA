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
	food []Food
}

func (m *Menu) Init() {
	defaultFoods := []Food{
		{Name: "Steak", Kinds: "Main", Price: 200000, Qty: 10},
		{Name: "Chicken", Kinds: "Main", Price: 25000, Qty: 0},
		{Name: "Pudding", Kinds: "Dessert", Price: 15000, Qty: 20},
		{Name: "Salad", Kinds: "Appetizer", Price: 25000, Qty: 30},
		{Name: "Soup", Kinds: "Appetizer", Price: 10000, Qty: 25},
	}

	for _, food := range defaultFoods {
		m.AddMenu(&food)
	}
}

func (m *Menu) GetFoods() []Food {
	return m.food
}

func (m *Menu) AddMenu(f *Food) {
	if f.Qty <= 0 {
		f.Status = false
	} else {
		f.Status = true
	}
	m.food = append(m.food, *f)

}

func (m *Menu) SearchMenu(name string) (*Food, error) {
	for i := range m.food {
		if m.food[i].Name == name {
			return &m.food[i], nil
		}

	}
	return nil, fmt.Errorf("Menu dengan nama '%s' tidak ditemukan", name)
}

func (m *Menu) DeleteMenu(name string) error {
	for i, f := range m.food {
		if f.Name == name {
			m.food = append(m.food[:i], m.food[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Menu dengan nama '%s' tidak ditemukan", name)

}

func (m *Menu) UpdateMenu(name, kinds string, price float64, qty int, status bool) error {
	for i, f := range m.food {
		if f.Name == name {
			m.food[i].Name = name
			m.food[i].Kinds = kinds
			m.food[i].Price = price
			m.food[i].Qty = qty
			if qty <= 0 {
				m.food[i].Status = false
			} else {
				m.food[i].Status = true
			}
			return nil
		}
	}
	return fmt.Errorf("Menu dengan nama '%s' tidak ditemukan", name)
}
