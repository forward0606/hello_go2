package pkCircle

import (
	"fmt"
	"math"

	pkShape "github.com/forward0606/hello_go2/pkShape"
)

type Circle struct {
	pkShape.Shape
	radius int
}

func NewCircle(col string, rad int) Circle {
	s := pkShape.NewShape(col)
	return Circle{Shape: s, radius: rad}
}

func (c Circle) GetArea() float64 {
	return float64(c.radius*c.radius) * math.Pi
}

func (c Circle) Print() {
	fmt.Println("type : circle")
	c.Shape.Print()
	fmt.Println("area : ", c.GetArea())
}
