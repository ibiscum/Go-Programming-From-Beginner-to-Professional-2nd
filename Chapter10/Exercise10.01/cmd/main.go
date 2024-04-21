package main

import "github.com/ibiscum/Go-Programming-From-Beginner-to-Professional-2nd/Chapter10/Exercise10.01/pkg/shape"

func main() {
	t := shape.Triangle{Base: 15.5, Height: 20.1}
	r := shape.Rectangle{Length: 20, Width: 10}
	s := shape.Square{Side: 10}
	shape.PrintShapeDetails(t, r, s)
}
