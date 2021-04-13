package main

import "fmt"

func main() {
	var mp map[string]int
	fmt.Println(mp)
	if mp == nil {
		fmt.Println("m is nil")
	}

	// Initialize
	var m = make(map[string]int)

	fmt.Println(m)

	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println("m is not nil")
	}

	// make() function returns an initialized and ready to use map.
	// Since it is initialized, you can add new keys to it.
	m["one hundred"] = 100
	fmt.Println(m)

	// Checking if a key exists
	var employees = map[int]string{
		1001: "John",
		1002: "Steve",
		1003: "Maria",
	}

	printEmployee(employees, 1001)
	printEmployee(employees, 1010)

	if isEmployeeExists(employees, 1002) {
		fmt.Println("EmployeeId 1002 found")
	}
}

func printEmployee(employees map[int]string, employeeId int) {
	if name, ok := employees[employeeId]; ok {
		fmt.Printf("name = %s, ok = %v\n", name, ok)
	} else {
		fmt.Printf("EmployeeId %d not found\n", employeeId)
	}
}

func isEmployeeExists(employees map[int]string, employeeId int) bool {
	_, ok := employees[employeeId]
	return ok
}
