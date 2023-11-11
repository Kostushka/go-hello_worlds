package main

import "fmt"

func double(s []int) {
	s = append(s, s...)
}

func main() {
	s := []int{1, 2, 3}
	double(s)
	fmt.Println(s, len(s)) // prints [1 2 3] 3
}
