package main

import "fmt"

type tree struct {
	value int
	left, right *tree
}

func main() {
	arr := []int{2, 2, 5, 1, 3, 4, 12, 8}
	fmt.Println(arr)
	Sort(arr)
	fmt.Println(arr)
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		// values = appendValues(values, t.left)
		// values = append(values, t.value)
		// fmt.Printf("append: %p %d\n", t, t.value)
		// values = appendValues(values, t.right)
		values = appendValues(values, t.right)
		values = append(values, t.value)
		values = appendValues(values, t.left)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// адрес tree
		t = new(tree)
		t.value = value
		fmt.Printf("%p %d\n", t, t.value)
		return t
	}
	if value < t.value {
		fmt.Printf("left ")
		t.left = add(t.left, value)
	} else {
		fmt.Printf("right ")
		t.right = add(t.right, value)
	}
	return t
}
