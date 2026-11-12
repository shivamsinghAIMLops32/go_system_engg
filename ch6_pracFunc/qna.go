package main

import "fmt"

func testFirst() {
	x := 10
	fn := func() {
		fmt.Println("value of x is :", x)
	}
	x = 50
	// here when fn is called, it will print the updated value of x
	fn()
}

// Returns a closure which multiplies it's input by factor
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// count++ is not just 1 operation
// it has 3 stages
// READ count from memory
// add 1 to count
// write count back to memory
// so it is not atomic operation or thread-safe
