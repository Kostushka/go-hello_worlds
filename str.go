package main

import "fmt"

func main() {
	str := "\x30\xab\xe2\x8c\x98"
	fmt.Println("str: ", str)
	fmt.Printf("bytes: % x\n", str)
	fmt.Printf("print and non print: %+q\n", str)
	// перебор строки по символам
	for i, rune := range str {
		fmt.Printf("rune: %#U [%d]\n", rune, i)
	}
	// перебор строки по байтам
	for i := 0; i < len(str); i++ {
		fmt.Printf("byte[%d]: %x\n", i, str[i])
	}
}
