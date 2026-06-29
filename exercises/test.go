package main

import (
	"corola/polymorphism"
)

func main() {

	polymorphism.PrintArea(polymorphism.Rectangle{H: 2, W: 3})
	polymorphism.PrintArea(polymorphism.Square{L: 3})
	polymorphism.PrintArea(polymorphism.Triangle{B: 2.2, H: 3.3})

}
