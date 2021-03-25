package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	list := strings.Split(s, " ")
	m := make(map[string]int)

	for _, v := range(list) {

		m[v] += 1
	}

	return m
}

func main() {

	phrase1 := "Cry Over Spilled Milk"
	phrase2 := "The greatest glory in living lies not in never falling, but in rising every time we fall"
	phrase3 := "The way to get started is to quit talking and begin doing"
	phrase4 := "If you look at what you have in life, you'll always have more. If you look at what you don't have in" +
		" life, you'll never have enough"


	fmt.Println(WordCount(phrase1))
	fmt.Println(WordCount(phrase2))
	fmt.Println(WordCount(phrase3))
	fmt.Println(WordCount(phrase4))
}