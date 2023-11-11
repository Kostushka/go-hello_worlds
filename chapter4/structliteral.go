package main

import "fmt"

type Point struct {
	x int
	y int
}

type Point2 struct {
	x int
	y int
}

type Circle struct {
	Point
	radius int
}

func main() {
	p := Point{1, 3}
	fmt.Printf("%d %d\n", p.x, p.y)

	p2 := Point2{y: 2, x: 4}
	fmt.Printf("%d %d\n", p2.x, p2.y)

	c := Circle{Point{4, 4}, 5}
	fmt.Printf("%#v\n", c)
	c.x = 6
	fmt.Printf("%v\n", c)
}
