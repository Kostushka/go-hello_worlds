package main

import "fmt"

func main() {
	// строка
	str := "Привет"
	fmt.Printf("%s\n", str)
	// срез
	a := []byte(str)
	fmt.Printf("%s\n", a)
	reverse(a)
	fmt.Printf("%s\n", a)
}

func reverse(arr []byte) {
	// символы юникода
	a := []rune(string(arr))
	for i, j := 0, len(a) - 1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	copy(arr, []byte(string(a)))
}
