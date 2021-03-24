package main

import (

	"fmt"
	"math"
)

func main() {

	//Generate an error for be an unexported (private) name
	//fmt.Println(math.pi)
	//Exported (public) name is with capital letter
	fmt.Println(math.Pi)
}
