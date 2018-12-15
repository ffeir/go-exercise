package variables

import "fmt"

func PrintVariables()  {
	var age int // variable declaration
	fmt.Println("my age is ", age)
	age = 29 //assignment
	fmt.Println("my age is", age)
	age = 54 //assignment
	fmt.Println("my new age is", age)

	var num int = 11
	fmt.Println("num is ", num)

	var numInferred = 22
	fmt.Println("num_inferred is ", numInferred)

	var width, height = 100, 50 //declaring multiple variables
	fmt.Println("width is", width, "height is", height)

	var w, h int
	fmt.Println("w is ", w, "h is ", h)
	w = 100
	h = 50
	fmt.Println("w is ", w, "h is ", h)

	var (
		name   	= "naveen"
		age1   	= 29
		height1 int
	)
	fmt.Println("my name is", name, ", age is", age1, "and height is", height1)
}

func ShortHandDeclaration() {
	fmt.Println("========Short Hand========")
	name, age := "naveen", 29 //short hand declaration
	fmt.Println("my name is", name, "age is", age)

	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)
}
