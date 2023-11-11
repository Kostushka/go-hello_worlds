package main

import "fmt"

func main() {
	a := []string{"one", "", "three"}
	// изменяет исходный срез
	b := nonempty(a)
	fmt.Printf("%s %s\n", a, b)
	c := []string{"", "", "two", "", "hi"}
	fmt.Printf("%s\n", c)
	fmt.Printf("%s\n", nonempty2(c))
	fmt.Printf("%s\n", c)
}

func nonempty(x []string) []string {
	i := 0
	for _, r := range x {
		if r != "" {
			x[i] = r
			i++
		}
	}
	return x[:i]
}

func nonempty2(x []string) []string {
	var z []string
	for _, r := range x {
		if r != "" {
			z = append(z, r)
		}
	}
	return z
}
