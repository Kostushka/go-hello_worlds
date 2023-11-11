package main

import "fmt"

func main() {
	// нельзя менять элементы
	str := "string"
	// str[0] = 'r'
	// срез с указателем на копию байт строки 
	b := []byte(str)
	b[0] = 'r'
	fmt.Printf("%s %s\n", str, b)

	// можно менять элементы
	arr := [5]int{1, 2, 3, 4, 5}
	// срез с указатель на исходный массив
	a := arr[1:3]
	a[0] = 13
	fmt.Printf("%d %d\n", arr, a)
}
