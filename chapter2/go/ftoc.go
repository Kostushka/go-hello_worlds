package main

import "fmt"

func main() {
	const freezing, boiling = 32.0, 212.0

	fmt.Printf("%gF = %gC\n", freezing, ftoc(freezing))
	fmt.Printf("%gF = %gC\n", boiling, ftoc(boiling))
}

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}
