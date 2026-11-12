package main

import "fmt"

// we can also declare a fnc type too
var operation func(int, int) int // operation has a type of fun which takes 2 int as parameter and return

func work(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func subtract(a, b int) int {
	return a + b
}

func main() {
	// in go functions are first class citizens so we can treat a function as an value too or pass in another function
	// Here we assigned sum function to s
	// we didn't used sum() or else it would invoke the func immediately
	s := sum
	fmt.Println(s(10, 20)) // we passing valued here to s

	fmt.Println(work(10, 30, subtract))

	rtfn := returnFn(1000, 3000)
	fmt.Println(rtfn(6000))

	// anonymous function
	// we assingned the anonymous func to a variable so we can use it later
	// we can invoke the anonymous func just like any other function
	anonFN := func(x int) int {
		return x * 2
	}

	fmt.Println(anonFN(5))

	// another anonymous function which does small task
	func() {
		fmt.Println("did some anonymous work")
	}() // immediately invoked function

	// lexical closure
	count := 0

	icr := func() {
		count++
	}

	icr()
	icr()
	fmt.Println("count is", count)
}

func sum(a, b int) int {
	return a + b
}
