package main

import "fmt"

func main() {
	a := map[string]int{
		"root": 1,
	}
	b := map[string]int{
		"root": 1,
		"tmp": 6,
	}
	fmt.Println(equal(a, b))
}

func equal(x, y map[string]int) bool {

	if len(x) != len(y) {
		return false
	}

	for k, vx := range x {
		if vy, ok := y[k]; !ok || vx != vy {
			return false
		}
	}

	return true
}
