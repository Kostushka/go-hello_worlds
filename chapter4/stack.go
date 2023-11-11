package main

import "fmt"

func main() {
	v := 2
	var stack []int
	// добавлен в стек
	stack = append(stack, v)
	fmt.Printf("%d\n", stack)
	// вершина стека
	top := stack[len(stack)-1]
	fmt.Printf("%d\n", top)
	// удален из стека
	stack = stack[:len(stack)-1]
	fmt.Printf("%d\n", stack)

	a := []int{1, 2, 3, 4, 5}
	a = remove(a, 2)
	fmt.Printf("%d\n", a)
	a = remove2(a, 2)
	fmt.Printf("%d\n", a)
}

func remove(x []int, i int) []int {
	copy(x[i:], x[i+1:])
	return x[:len(x)-1]
}

func remove2(x []int, i int) []int {
	x[i] = x[len(x)-1]
	return x[:len(x)-1]
}
