package main

import "fmt" //import语言的fmt库——用于输出

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

	//经典的for语句 init; condition; post
	for i := 0; i<10; i++{
		 fmt.Println(i)
	}
	
	var i int = 1
	var pInt *int = &i
	//输出：i=1     pInt=0xf8400371b0       *pInt=1
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)
	*pInt = 2
	//输出：i=2     pInt=0xf8400371b0       *pInt=2
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)
	i = 3
	//输出：i=3     pInt=0xf8400371b0       *pInt=3
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)


	var p *[]int = new([]int)   // 为切片结构分配内存；*p == nil；很少使用
	var v  []int = make([]int, 10) // 切片v现在是对一个新的有10个整数的数组的引用
	fmt.Println(v)
	// 不必要地使问题复杂化：
	var p2 *[]int = new([]int)
	fmt.Println(p2) //输出：&[]
	*p = make([]int, 10, 10)
	fmt.Println(p) //输出：&[0 0 0 0 0 0 0 0 0 0]
	fmt.Println((*p)[2]) //输出： 0
	// 习惯用法:
	v2 := make([]int, 10)
	fmt.Println(v2) //输出：[0 0 0 0 0 0 0 0 0 0]
}