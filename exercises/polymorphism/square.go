package polymorphism

type Square struct {
	L float32
}

func (a Square) Area() float32 {
	return a.L * a.L
}
