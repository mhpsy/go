package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	for i, arg := range os.Args {
		fmt.Println(i, "     ", arg)
	}

	fmt.Println(s)
}
