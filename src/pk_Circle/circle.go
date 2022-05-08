package pk_Circle

import (
	"fmt"
	"math"

	shape "github.com/forward0606/hello_go2/src/pk_shape"
)

type Circle struct {
	shape.Shape
	radius int
}

func New_Circle(col string, rad int) Circle {
	s := shape.New_Shape(col)
	return Circle{Shape: s, radius: rad}
}

func (c Circle) Get_area() float64 {
	return float64(c.radius*c.radius) * math.Pi
}

func (c Circle) Print() {
	fmt.Println("type : circle")
	c.Shape.Print()
	fmt.Println("area : ", c.Get_area())
}
