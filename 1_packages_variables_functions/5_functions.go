package main

import "fmt"

func add(x int, y int) int {

	return x + y
}

func main() {

	fmt.Printf("2 + 3 is equal to %d\n", add(2, 3))
}
