package main

import (
	"fmt"
	"os"
	"strings"
)

// Выводит аргументы командной строки
func main() {
	fmt.Println(strings.Join(os.Args[1:], " ")) // функция Join из пакета strings
	// или вывод для отладки
	fmt.Println(os.Args[1:])
}
