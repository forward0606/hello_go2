package pk_triangle

import (
	"fmt"

	"github.com/forward0606/hello_go2/src/pk_shape"
)

type Triangle struct {
	pk_shape.Shape

	base   int
	height int
}

func New_Triangle(col string, _base int, _height int) Triangle {
	s := pk_shape.New_Shape(col)
	return Triangle{Shape: s, base: _base, height: _height}
}

func (t Triangle) Get_area() float64 {
	return float64(t.height*t.base) / 2
}

func (t Triangle) Print() {
	fmt.Println("type : triangle")
	t.Shape.Print()
	fmt.Println("area : ", t.Get_area())
}
