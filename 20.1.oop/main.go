package main

import "oop/employee"

func main() {
	e := employee.New("Sam", "Abolf", 30, 20)

	e.LeavesRemaining()
}
