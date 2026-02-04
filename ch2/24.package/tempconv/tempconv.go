// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import (
	"fmt"
	"os"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

var cwd string

// init函数可以有多个 这个设计有点奇怪哦
func init() {
	// 这个太坑了，因为内部声明的cwd将屏蔽外部的声明，因此上面的代码并不会正确更新包级声明的cwd变量。
	// cwd, err := os.Getwd()
	// 这里不能使用:=，否则会创建一个新的局部变量cwd，导致包级变量cwd没有被赋值
	cwd, _ = os.Getwd()
	fmt.Println("Current working directory:", cwd)
	fmt.Println("tempconv package initialized1")
}

func init() {
	fmt.Println("tempconv package initialized2")
}
