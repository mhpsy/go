package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		return
	}
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		counts := countLines(string(data))
		for line, n := range counts {
			fmt.Printf("%s: %d\n", line, n)
		}
	}
}

func countLines(data string) map[string]int {
	counts := make(map[string]int)
	for _, line := range strings.Split(data, "\n") {
		counts[line]++
	}
	return counts
}
