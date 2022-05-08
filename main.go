package main

import (
	pkCircle "github.com/forward0606/hello_go2/pkCircle"
	pkRec "github.com/forward0606/hello_go2/pkRectangle"
	pkTri "github.com/forward0606/hello_go2/pkTriangle"
)

func main() {
	circle := pkCircle.NewCircle("green", 5)
	circle.Print()

	rec := pkRec.NewRectangle("red", 5, 8)
	rec.Print()

	tri := pkTri.NewTriangle("red", 5, 8)
	tri.Print()

}
