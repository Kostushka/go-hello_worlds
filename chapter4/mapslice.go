package main

import "fmt"

var h = make(map[string]int)

func main() {
	str := []string{"Moscow", "Zaozersk", "Sevastopol", "Murmansk"}
	Add(str[1:3])
	Add(str[1:3])
	fmt.Println(Count(str[1:3]))
	for name := range h {
		fmt.Println(name)
	}
}

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	h[k(list)]++
}

func Count(list []string) int {
	return h[k(list)]
}
