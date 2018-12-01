package main

import "fmt"

func main()  {
	var arr [2]int
	fmt.Println("arr", arr)

	//arrayDot := [...]float64{7.0, 8.5, 9.1}
	//array := []float64{7.0, 8.5, 9.1}
	// fmt.Println("array", array, arrayDot, array == arrayDot)

	a := [3]int{1:1, 2:3}
	a1 := [...]int{1:1, 2:3}
	a2 := []int{1:1, 2:3}
	fmt.Println("a", a, a1, a2, len(a), len(a1), len(a2))
	fmt.Printf("%T, |%T|\n", a1, a2)
}