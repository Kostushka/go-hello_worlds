package main

import "fmt"

// Функция вставляет запятые через каждые три цифры в строке
func main() {
	str := "123456789"

	fmt.Printf("%s\t%s\n", str, comma(str))
}

func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
