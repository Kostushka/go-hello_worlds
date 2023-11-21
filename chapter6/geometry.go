package main

import "math"
import "fmt"

func main() {
	perim := Path{
		{1, 1},
		{5, 1},
		{1, 4},
		{1, 1},
	}
	// метод среза
	fmt.Println(perim.Distance())
}

// срез координат
type Path []Point

// координаты точки
type Point struct {
	X float64
	Y float64
}

// метод точки
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

// метод среза
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		fmt.Printf("i: %d\n", i)
		if i > 0 {
			fmt.Printf("%p %p\n", path[i-1], path[i])
			// метод точки
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
