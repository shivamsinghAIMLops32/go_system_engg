package main

import (
	"fmt"
)

func main() {
	// creating a counter genreator independent
	f1 := countGenerator()
	f2 := countGenerator()

	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f2())

	testFirst()

	// use multiplier
	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	fmt.Println(double(5))
	fmt.Println(triple(5))
}

func countGenerator() func() int {
	// Each closure gets it's own counter so they all start at same as counter = 0
	counter := 0

	return func() int {
		counter++
		return counter
	}
}
