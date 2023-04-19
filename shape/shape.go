package shape

import "fmt"

type Shape interface {
	getArea() float64
}

type Triangle struct {
	Height float64
	Base   float64
}

type Square struct {
	SideLength float64
}

func (s Square) getArea() float64{
	return s.SideLength * s.SideLength
}
func (s Triangle) getArea() float64{
	return 0.5 * s.Base * s.Height
}

func PrintArea(s Shape) {
	fmt.Printf("Area of the shape is : %f " , s.getArea())
}

