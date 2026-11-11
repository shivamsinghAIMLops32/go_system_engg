package main

import (
	"fmt"

	"ch3/calculation"
)

func main() {
	sum := calculation.Add(1, 2, 3, 4, 5)
	fmt.Println(sum)

	// change the value of sum back to 0
	changeTOZeroVAlue(&sum)
	fmt.Println(&sum)
	fmt.Println(sum)

	val, err := calculation.Divide(5, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}

func changeTOZeroVAlue(valptr *int) {
	*valptr = 0
}
