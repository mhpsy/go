package main

import (
	"fmt"

	"github.com/mhpsy/package/tempconv"
)

func main() {

	var c tempconv.Celsius = 25.0
	var f tempconv.Fahrenheit = tempconv.CToF(c)

	fmt.Println(c.String())
	fmt.Println(f.String())

}
