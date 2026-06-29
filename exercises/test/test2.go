package test

import "fmt"

type AREA struct {
	Width  float32
	Height float32
}

type Shape interface {
	Test() float32
}

func AreaPrint(a Shape) {
	fmt.Printf("area: %v\n", a.Test())
}

func (area AREA) Test() float32 {

	return area.Height * area.Width
}
