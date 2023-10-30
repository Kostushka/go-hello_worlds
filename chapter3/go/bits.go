package main

import "fmt"

func main() {

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)
	
	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)

	for i := 0; i < 8; i++ {
		if x & (1<<i) != 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
