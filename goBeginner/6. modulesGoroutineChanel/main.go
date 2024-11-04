package main

import (
	"fmt"
	"github.com/RifqiF27/package-add-data"
)

func main() {
	var dataInput []addData.Data
	for i := 0; i < 10; i++ {
		dataInput = append(dataInput, addData.Data{
			ID:   i + 1,
			Name: fmt.Sprintf("Data %d", i+1),
		})
	}

	addData.ProcessData(dataInput)
}
