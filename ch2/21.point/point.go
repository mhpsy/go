package main

import "fmt"

var x, y int

var p *int

func main() {
	fmt.Println(&x == &y, &x == &x, &x == nil)
	// F T F

	// &x == nil, 因为x有初始值，所以不可能是nil

	fmt.Println(p == nil)
	// T

	fmt.Println(foo() == foo())
	// F

}

func foo() *int {
	num := 3
	return &num
}
