package main

import "fmt"

func main() {
	fmt.Printf("%d\n", triple(5))
}

func triple(x int) (result int) {
	defer func() {
		result += x
	}()
	return x + x
}
