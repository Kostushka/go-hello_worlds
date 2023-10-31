package main

import "fmt"

// Функция убирает компоненты каталога и суффикс файла
func main() {
	str := "hello/hi/c.go"

	fmt.Printf("basename: %s\tstr: %s\n", basename(str), str)
}

func basename(str string) string {
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '/' {
			str = str[i+1:]
			break
		}
	}

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '.' {
			str = str[:i]
			break
		}
	}

	return str
}
