package main

import (
	"fmt"
	"os"
)

// Программа выводит индекс и значение каждого аргумента в отдельной строке
func main() {

	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
