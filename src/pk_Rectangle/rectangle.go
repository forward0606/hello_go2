package pk_Rectangle

import (
	"fmt"

	pk_shape "github.com/forward0606/hello_go2/src/pk_shape"
)

type Rectangle struct {
	pk_shape.Shape
	height int
	width  int
}

func New_Rectangle(col string, hei int, wid int) Rectangle {
	s := pk_shape.New_Shape(col)
	return Rectangle{Shape: s, height: hei, width: wid}
}

func (r Rectangle) Get_area() float64 {
	return float64(r.height * r.width)
}

func (r Rectangle) Print() {
	fmt.Println("type : rectangle")
	r.Shape.Print()
	fmt.Println("area : ", r.Get_area())
}
