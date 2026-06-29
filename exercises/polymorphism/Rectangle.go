package polymorphism

type Rectangle struct {
	H float32
	W float32
}

func (a Rectangle) Area() float32 {
	return a.H * a.W
}
