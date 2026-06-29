package polymorphism

import "fmt"

type Shape interface {
	Area() float32
}

func PrintArea(a Shape) {
	fmt.Printf("area: %v\n", a.Area())
}
