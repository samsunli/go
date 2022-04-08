package main

import "fmt"

func main() {
	defer println("World")
	println("Hello")
}

func println(str string) {
	fmt.Println(str)
}