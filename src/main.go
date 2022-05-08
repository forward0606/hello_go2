package main

import (
	pk_circle "github.com/forward0606/hello_go2/src/pk_Circle"
	pk_rec "github.com/forward0606/hello_go2/src/pk_Rectangle"
	pk_tri "github.com/forward0606/hello_go2/src/pk_Triangle"
)

func main() {
	circle := pk_circle.New_Circle("green", 5)
	circle.Print()

	rec := pk_rec.New_Rectangle("red", 5, 8)
	rec.Print()

	tri := pk_tri.New_Triangle("red", 5, 8)
	tri.Print()
}
