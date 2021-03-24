package main

import "fmt"

//Functions continued is when replace "x int, y int" to "x, y int" because is the same type
func add2(x, y int) int {

	return x + y
}

func main() {

	fmt.Printf("2 + 3 is equal to %d\n", add2(2, 3))
}
