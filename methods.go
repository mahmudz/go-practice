package main

import "fmt"

type Rect struct {
	height int
	width int
}

func (r Rect) area() int  {
	return r.height * r.width;
}

func main()  {
	box := Rect{5,6}

	fmt.Println(box.area())
}
