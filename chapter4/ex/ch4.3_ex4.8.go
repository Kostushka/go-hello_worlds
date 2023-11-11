package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	difc := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		// считать кол-во букв
		if unicode.IsLetter(r) {
			difc["letter"]++ 
		}
		// считать кол-во цифр
		if unicode.IsDigit(r) {
			difc["digit"]++
		}
		// считать кол-во пробельных символов
		if unicode.IsSpace(r) {
			difc["space"]++
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("len\tcount\n")
	for i, n := range utflen {
		fmt.Printf("%d\t%d\n", i, n)
	}

	for name, count := range difc {
		fmt.Printf("%s\t%d\n", name, count)
	}

	if invalid > 0 {
		fmt.Printf("\n%d неверных символов UTF-8\n", invalid)
	}
}
