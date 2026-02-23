package interfaces

import "fmt"


type Shape interface {
	getArea() uint
}

type Square struct{
	Width uint
}

type Rectangle struct{
	Length uint
	Breadth uint
}

func (s Square) getArea() uint {
	return s.Width * s.Width
}

func (r Rectangle) getArea() uint {
	return r.Breadth * r.Length
}

func InterfaceShapes() {
	var shapes []Shape = []Shape{Square{5}, Rectangle{4,4}}

	area := uint(0)

	for _, shape := range shapes{
		area += shape.getArea()
	}

	fmt.Println(area)
}