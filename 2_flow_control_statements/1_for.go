package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i++ {

		sum += i
	}

	fmt.Println(sum)
	//Error, i just exists in the scope of the loop
	//fmt.Println(i)

	sum = 1
	for ; sum < 1000 ; {

		sum += sum
	}

	fmt.Println(sum)

	sum = 1
	for ;  ; {

		sum += sum
		if sum >= 1000 {

			break
		}
	}

	fmt.Println(sum)
}
