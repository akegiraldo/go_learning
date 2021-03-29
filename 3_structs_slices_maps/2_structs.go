package main

import "fmt"

type Vertex struct {

	X int
	Y int
}

func main() {

	v := Vertex{1 , 2}
	fmt.Println(v)

	v.X = 4
	fmt.Println(v)

	p := &v
	p.X = 1e9 //it's same that (*p).X without explicit dereference
	fmt.Println(v)
}
