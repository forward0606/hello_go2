package pkRectangle

import (
	"fmt"

	pkShape "github.com/forward0606/hello_go2/pkShape"
)

type Rectangle struct {
	pkShape.Shape
	height int
	width  int
}

func NewRectangle(col string, hei int, wid int) Rectangle {
	s := pkShape.NewShape(col)
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
