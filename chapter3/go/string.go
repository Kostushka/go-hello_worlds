package main

import "fmt"
import "unicode/utf8"

func main() {

	str := "Hello Привет"

	fmt.Printf("%s\n", str)

	// обычный вывод всех символов строки
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	fmt.Println()

	// вывод с декодированием рун (символов) юникода
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%2d%6c\n", i, r)
		i += size	
	}

	// вывод с декодированием рун (символов) юникода
	for i, r := range str {
		fmt.Printf("%2d%6q%6d\n", i, r, r)
	}

	// кол-во символов
	n := 0
	for range str {
		n++	
	}
	fmt.Printf("Кол-во символов: %d\n", n)

	s := "Программа"

	fmt.Printf("последовательность байт: % x\n", s)
	// последовательность символов Unicode
	r := []rune(s)
	fmt.Printf("последовательность символов: %x\n", r)
	fmt.Println(string(r))
	fmt.Println(string(65))
	
	
	
	// длина в байтах и кол-во рун (символов)
	fmt.Printf("%d %d\n", len(str), utf8.RuneCountInString(str))
}
