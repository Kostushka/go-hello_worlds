package main
 
import (
    "fmt"
)
 
// Square object
type square struct {
    name string
    sideLen int
}
func (sq *square) Area() int {
    return sq.sideLen * sq.sideLen
}
func (sq *square) Name() string {
    return sq.name
}
 
// Cube object
type cube struct {
    square
    name string
}
func (c *cube) Area() int {
    return c.square.Area() * 6
}
func (c *cube) Facet() int {
    return c.square.Area()
}
func (sq *cube) Name() string {
    return sq.name
}
 
// Circle object
type circle struct {
    name string
}
func (c *circle) Name() string {
    return c.name
}
 
// Interfaces
type geomArea interface {
    Area() int
}
 
type geomName interface {
    Name() string
}
 
func main() {
    sq := &square{name: "square", sideLen: 5}
 
    c := &cube{
        square: square{sideLen: 5},
        name: "cube",
    }
 
    circ := &circle{name: "circle"}
 
    fmt.Println("Names:")
 
    names := []geomName{ sq, c, circ }
    for _, v := range names {
        fmt.Println(v.Name())
    }
}
