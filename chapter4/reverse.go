package main

import "fmt"

func main() {

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	b := []int{6, 7, 8, 9, 10}
	reverse(b)
	fmt.Println(a)
	fmt.Println(b)
}

func reverse(str []int) {
	for i, j := 0, len(str) - 1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}
