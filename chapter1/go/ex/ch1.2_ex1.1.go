package main

import (
	"fmt"
	"os"
)

// Программа echo выводит помимо аргументов еще и имя исполняемой программы
func main() {
	s, sep := "", ""
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
