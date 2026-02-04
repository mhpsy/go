package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))         // 输出字节长度
	fmt.Println(len([]rune(s))) // 输出字符长度

	for _, ss := range s {
		fmt.Print(ss)
		fmt.Print(" ", string(ss), " ")
		sss := utf8.RuneLen(ss)
		// 根据字符的字节长度打印对应数量的下划线
		for i := 1; i < sss; i++ {
			fmt.Print("1")
		}
		fmt.Print("\t")
	}

	fmt.Println()
	fmt.Println("====================")

	i := 0

	for range s {
		i++
	}
	fmt.Println(i)

	t := s
	s += "!"
	fmt.Println(t) //  t是不会变的，字符串是不可变的
	fmt.Println(s)

	// s[0] = 'h' // 编译错误，字符串是不可变的

	s2 := `
	----
This is a raw string literal.
It can span multiple lines.`

	fmt.Println(s2)

	s3 := "世界"
	s4 := "\xe4\xb8\x96\xe7\x95\x8c"
	s5 := "\u4e16\u754c"
	s6 := "\U00004e16\U0000754c"

	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)

	fmt.Println(s3 == s4, s4 == s5, s5 == s6)

}
