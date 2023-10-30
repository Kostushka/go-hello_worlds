package main

import (
	"fmt"
	"flag"
	"strings"
)

// переменные-флаги
var n = flag.Bool("n", false, "пропуск символа новой строки")
var s = flag.String("s", " ", "разделитель")

// Выводит аргументы командной строки с возможностью использовать ключи
func main() {
	// переопределение переменных-флагов
	flag.Parse()

	// применение разделителя
	fmt.Print(strings.Join(flag.Args(), *s))
	// применение символа \n
	if !*n {
		fmt.Println()
	}
	
}
