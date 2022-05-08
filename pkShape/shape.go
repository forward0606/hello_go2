package pkShape

import (
	"fmt"
)

type Shape struct {
	color string
}

func NewShape(col string) Shape {
	return Shape{color: col}
}

func (s Shape) Print() {
	fmt.Println("color : ", s.color)
}
