package main

import "fmt"

type Human interface {
	Display()
}
type Employee struct {
	name string; designation string
}
func (emp Employee) Display(){
	fmt.Println("Name - ", emp.name, ", Designation - ", emp.designation)
}
type Contractor struct {
	name string; weeklyHours int
}
func (cont Contractor) Display(){
	fmt.Println("Name - ", cont.name, ", Weekly Hours - ", cont.weeklyHours)
}
func main(){
	var emp Human = Employee{name:"Rob Pike", designation:"Engineer"}
	emp.Display()
	var cont Human = Contractor{name:"XYZ", weeklyHours:35}
	cont.Display()
}
