package main

import (
	"strconv"
	"fmt"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Printf("%s %s\n", y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 2))
	s := fmt.Sprintf("x=%b", x)
	fmt.Printf("%s\n", s)

	x, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d\n", x)

	n, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d\n", n)
}
