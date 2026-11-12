package main

import "fmt"

func main() {
	// does new keyword always created pointer on heap

	ptr := new([]byte)
	// no this could be on heap or either on stack defends if pointer outlives or need after function call
	name := "luca"

	// Here we defernceed a ptr cause new return a zero value of type intialzed with like *[]byte here
	// so we need to dereference it to assign a value at the address stored in ptr
	*ptr = []byte(name)

	r_ptr := createNumber()
	fmt.Println(*r_ptr)
}

func createNumber() *int {
	// here we gonna create a ptr to an integer and then return it to main
	// in this case the ptr will be on the heap
	// num := 42
	// ptr := &num
	// return ptr

	// alternate would be using new keyword which intializes a zero value of type and returns a pointer to it
	// here we created a value in createnumber function scope but we gave access to main also via ptr which outlives the function scope
	// so the value is still accessible even after the function returns
	var pitr = new(int)
	a := 10
	pitr = &a
	return pitr

	// // we can also use declaration method which return a nil pointer
	// var nilPtr *int
	// return nilPtr
}

// creating a fnc which returns a actual number so gabage coll of go will optimize and put it on stack
func stackCreateNmbr() int {
	x := new(int)
	*x = 200  // go to x stored address and put 200 init
	return *x // return the value at x's address
}
