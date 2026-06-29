package polymorphism

type Triangle struct {
	B float32
	H float32
}

func (a Triangle) Area() float32 {
	return 0.5 * a.B * a.H
}
