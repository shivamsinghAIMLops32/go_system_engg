package main

// this is function that returns a function
// also it uses closure to capture the values of a and b
// also we are reutrning a anonymous function
func returnFn(a, b int) func(int) int {
	return func(x int) int {
		return a + b + x
	}
}
