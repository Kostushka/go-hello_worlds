package main

import (
	"fmt"
)

func main() {
	s := "Ф"
	for _, v := range s {
		fmt.Printf("Code: %T (%d)\n", v, v)
	}

	fmt.Print("Code: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%T ", s[i])
	}
	fmt.Println()
}
