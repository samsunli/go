package main

import "fmt"

func main() {
    fmt.Println("hello world")
	cities := [...]string {"Kolkata", "Chennai", "Blore"}
	fmt.Println(cities[2]) // Blore
	fmt.Println(len(cities)) // 3
		
	fmt.Println("Slice..............................")	
	slice := []string {"Kolkata", "Bangalore", "Mumbai"}
	// slice are declared with syntax similar to array. Only the length is not specified.
	// slice := [3]string {"Kolkata", "Bangalore", "Mumbai"} // Array declaration with length
	fmt.Println(slice) // [Kolkata Bangalore Mumbai]
	fmt.Println(slice[1]) // Bangalore
	slice = append(slice,"Amsterdam")
	slice = append(slice,"Den Haag")
	fmt.Println(slice) // [Kolkata Bangalore Mumbai Amsterdam Den Haag]

	fmt.Println(slice) // [Kolkata Bangalore Mumbai Amsterdam Den Haag]
	duplicateSlice := make([]string, len(slice))
	copy(duplicateSlice, slice)
	fmt.Println(duplicateSlice) // [Kolkata Bangalore Mumbai Amsterdam Den Haag]
	fmt.Println(slice[0:2]) // [Kolkata Bangalore]
	fmt.Println(slice[:3]) // [Kolkata Bangalore Mumbai]
	fmt.Println(slice[2:]) // [Mumbai Amsterdam Den Haag]

	//for init; condition; post
	for i := 0; i<10; i++{
		 fmt.Println(i)
	}
	
	var i int = 1
	var pInt *int = &i
	//Print：i=1     pInt=0xf8400371b0       *pInt=1
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)
	*pInt = 2
	//Print：i=2     pInt=0xf8400371b0       *pInt=2
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)
	i = 3
	//Print：i=3     pInt=0xf8400371b0       *pInt=3
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)

}