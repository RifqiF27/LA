package main

import (
	"fmt"
	"reflect"
)

type Vehicle interface {
	mileage(gas int) int
}

type Speed struct {
	speed int
}
type Motorcycle struct {
	Speed
}
type Car struct {
	Speed
}
type Bajaj struct {
	Speed
}

func (m Motorcycle) mileage(gas int) int {
	return gas * m.speed
}
func (c Car) mileage(gas int) int {
	return gas * c.speed
}
func (b Bajaj) mileage(gas int) int {
	return gas * b.speed
}

func mostEfficient(gas int, vehicles ...Vehicle) Vehicle {
	var bestVehicle Vehicle
	mileages := make([]int, len(vehicles))

	for i, v := range vehicles {
		mileages[i] = v.mileage(gas)
	}

	maxMileage := mileages[0]
	bestVehicle = vehicles[0]

	for i, mileage := range mileages {
		if mileage > maxMileage {
			maxMileage = mileage
			bestVehicle = vehicles[i]
		}
	}

	fmt.Println("Mileages:", mileages)
	return bestVehicle
}

func main() {
	motor := Motorcycle{Speed{speed: 100}}
	car := Car{Speed{speed: 80}}
	bajaj := Bajaj{Speed{speed: 60}}

	best := mostEfficient(10, motor, car, bajaj)
	// fmt.Println("Most efficient vehicle:", best,reflect.TypeOf(best))
	switch reflect.TypeOf(best) {
	case reflect.TypeOf(motor):
		fmt.Print("Most efficient vehicle Motorcycle")
	case reflect.TypeOf(car):
		fmt.Print("Most efficient vehicle Car")
	case reflect.TypeOf(bajaj):
		fmt.Print("Most efficient vehicle Bajaj")
	}

}
