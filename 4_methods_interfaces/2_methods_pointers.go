package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {

	X, Y float64
}

func (v Vertex2) Abs() float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex2) float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale(f float64) {

	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale (v Vertex2, f float64) {

	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	v1 := Vertex2{3, 4}
	v2 := Vertex2{3, 4}

	fmt.Println("v1:", v1, "v2:", v2)

	v1.Scale(10)
	Scale(v2,10)

	fmt.Println("v1:", v1, "v2:", v2)
	fmt.Println("v1:", v1.Abs(), "v2:", Abs(v2))
	fmt.Println("v1:", v1, "v2:", v2)
}
