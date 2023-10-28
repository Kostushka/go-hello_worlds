package main

import (
	"fmt"
	"os"
)

func main() {

	var s, sep string // объявление двух переменных типа string (неявно инициализируются "")
	for i := 1; i < len(os.Args); i++ { // := краткое объявление переменной (тип переменной определяет инициализатор)
		s += sep + os.Args[i] // конкатенация строк и присвоение их в переменную
		sep = " "
	}
	fmt.Println(s)
}
