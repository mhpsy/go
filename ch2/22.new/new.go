package new

import "fmt"

func main() {

}

var global *int

// new可以创建匿名变量
func foo1() *int {
	return new(int)
}

func foo2() *int {
	var num int
	return &num
}

func foo3() {
	num := 0
	// 这个函数必须在堆上分配，因为逃逸出函数的作用域
	global = &num
}

func foo4() {
	// num 分配在堆还是栈上，由编译器决定，不一定是哪里
	num := new(int)
	fmt.Println(num)
}
