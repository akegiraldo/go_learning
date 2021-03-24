package main

import (

	"fmt"
	"math"
)

func main() {

	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)

	fmt.Printf("Type %T value %v \n", x, x)
	fmt.Printf("Type %T value %v \n", y, y)
	fmt.Printf("Type %T value %v \n", f, f)
	fmt.Printf("Type %T value %v \n", z, z)
}
