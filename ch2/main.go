package main

import (
	"ch2/calculator"
	"fmt"
)

func main() {
	greet("jason")
	if result := calculator.Sum(2, 3); result == 5 {
		fmt.Println("result is 5")
	}
}
