package pk_shape

import (
	"fmt"
)

type Shape struct {
	color string
}

func New_Shape(col string) Shape {
	return Shape{color: col}
}

func (s Shape) Print() {
	fmt.Println("color : ", s.color)
}
