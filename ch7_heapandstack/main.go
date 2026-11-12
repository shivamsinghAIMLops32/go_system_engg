package main

import "fmt"

func main() {
	k := stacked()
	p := heaped()

	fmt.Println(k)
	fmt.Println(*p)
}

func stacked() int {
	k := 10
	return k
}

// go build -gcflags="-m" main.go
func heaped() *int {
	p := new(int)

	*p = 10 // so go to address stored in p and then set the value from 0 to 10 but when we are returning the p which is a pointer to the address of the value 10 the value 10 escapes to heap
	return p
}
