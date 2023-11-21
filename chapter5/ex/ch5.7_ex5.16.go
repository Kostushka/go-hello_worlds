package main

import "fmt"

func main() {
	fmt.Println(join("Hello", "world", "!", "\\"))
}

func join(str ...string) string {
	// итоговая строка
	res := ""
	// последний элемент - разделитель
	sep := str[len(str) - 1]
	// срез без разделителя
	str = str[:len(str) - 1]
	
	for _, s := range str {
		res += s
		res += sep
	}
	result := []byte(res)
	result = result[:len(result) - 1]
	
	return string(result)
}
