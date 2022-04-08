package main

import "fmt"

type Employee struct {
	name string // field name of type string
	age int // field age of type int
	salary float32 // field salary of type float32
	designation string // field designation of type string
}

func (emp Employee) Display(){
	fmt.Println("Hello from Employee")
}
func main() {
	emp := Employee{}
	// method invokation
	emp.Display() // displays "Hello from Employee"
}
