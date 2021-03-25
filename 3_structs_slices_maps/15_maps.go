package main

import "fmt"

type Vertex3 struct {

	Lat, Long float64
}

var m1 map[string]Vertex3

func main() {

	m1 = make(map[string]Vertex3)
	m1["Bell Labs"] = Vertex3{

		40.68433, -74.39967,
	}

	fmt.Println(m1["Bell Labs"])

	// Map literals
	var m2 = map[string]Vertex3{

		"Bell Labs": Vertex3{
			40.68433, -74.39967,
		},
		"Google": Vertex3{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m2)

	// Map literals continued
	var m3 = map[string]Vertex3{

		"Bell Labs": { 40.68433, -74.39967 },
		"Google": { 37.42202, -122.08408 },
	}

	fmt.Println(m3)
}
