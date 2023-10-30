package main

import "fmt"

type Celsius float64
type Fahrenheit float64

// Именованые типы позволяют определить новое поведение с помощью методов
func main() {
	c := FtoC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%s\n", c)
	fmt.Printf("%g\n", c)
	fmt.Printf("%s\n", c)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
