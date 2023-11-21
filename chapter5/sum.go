package main

import "fmt"

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(4, 5))
	arr := []int{23, 45}
	fmt.Println(sum(arr...))
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
