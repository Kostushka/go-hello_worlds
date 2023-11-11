package main

import "fmt"

// указатель на массив вместо среза
func main() {
	// массив
	arr := [32]byte{'H', 'e', 'l', 'l', 'o'}
	fmt.Printf("%s\n", arr)
	reverse(&arr)
	fmt.Printf("%s\n", arr)
}

func reverse(ptr *[32]byte) {
	for i, j := 0, len(ptr) - 1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}
