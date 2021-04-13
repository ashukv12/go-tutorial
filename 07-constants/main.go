package main

import (
	"fmt"
)

func main() {
	const country, code = "India", 91

	const (
		employeeId string  = "E101"
		salary     float64 = 50000.0
	)

	fmt.Println(country, code, employeeId, salary)
}
