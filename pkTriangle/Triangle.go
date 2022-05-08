package pkTriangle

import (
	"fmt"

	pkShape "github.com/forward0606/hello_go2/pkShape"
)

type Triangle struct {
	pkShape.Shape

	base   int
	height int
}

func NewTriangle(col string, _base int, _height int) Triangle {
	s := pkShape.NewShape(col)
	return Triangle{Shape: s, base: _base, height: _height}
}

func (t Triangle) GetArea() float64 {
	return float64(t.height*t.base) / 2
}

func (t Triangle) Print() {
	fmt.Println("type : triangle")
	t.Shape.Print()
	fmt.Println("area : ", t.GetArea())
}
