package main

import "fmt"

// наибольший общий делитель двух целых чисел
func main() {
	fmt.Printf("%d\n", gcd(5, 30))	
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x % y
	}

	return x
}
