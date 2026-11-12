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

	// using new(type) intialzied a ptr varibale with 0 value and returns ptr
	// ptr := new(int)
	var ptr *string // here ptr has a nil pointer
	// if ptr == nil {
	// 	fmt.Println(*ptr) // this will result in panic as we dereference a nil pointer
	// }
	name := "jason"
	ptr = &name
	fmt.Println(*ptr)

	fmt.Println("started doing some work")
	somework()
}

// even pointer is copied using value but those 2 values has same address initself so when we dereference it we get the value at address
func changeTOZeroVAlue(valptr *int) {
	*valptr = 0
}

func somework() {
	x := 10
	// this will be executed last as it enters first
	// also x value captured in defer the  moment it comes but it's just that it will execute when func exists
	// so x here is 10 captured in defer
	defer fmt.Println(x)

	// this will be executed third cause defer works as stack last in first out
	// same as previous here x is captured as 20 in defer
	x = 20
	defer fmt.Println(x)

	// this will be executed second cause defer works as stack last in first out
	x = 100
	defer fmt.Println(x)

	// this will be executed first cause defer works as stack last in first out
	defer fmt.Println("i am done working")
}
