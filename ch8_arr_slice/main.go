package main

import "fmt"

func main() {
	// in GO array has fixed size baked into the type altho we can change values at a index
	// but we can't change the size of the array
	// so [3]int != [4]int they are both different types
	arr := [3]int{1, 2, 3}
	arr[1] = 10
	fmt.Println(arr)

	// GO slices are dynamic and has pointer
	// they are struct of 24 bytes has
	// type slices struct{
	// 	len int        8 bytes   number of elements currently in the slice
	// 	cap int        8 bytes   number of elements the slice can hold
	// 	ptr *int       8 bytes   ptr to the underlying array
	// }

	// there are 2 scenerios when a slice is passed to a function
	// 1. if orignal slice still has capacity, the slice passed to function will be get a copy of sturct of main slice and
	// the underlying array will be shared so if we append to func slice, the original slice sturct won't be updated i mean the length part but
	// if we get a new window lookup of original slice, the updated length will be reflected
	// so either we can return a new slice as append returns a new slice with updated length
	// or we can use reslicing

	//  2. if original slice doesn't have capacity, the slice passed to function will be getting a copy of orginal slice
	// so if we append to func slice, the original slice sturct and func slice will have different underlying arrays as
	// the capacity was also copied and it was full so we need to allocate a new underlying array for func slice
	// in this case, we also need to return the new slice from the function

	modifyArr(arr)
	fmt.Println(arr)
	modifySlice(arr[:])
	fmt.Println(arr)
	newArr := appSlice(arr[:])
	fmt.Println(arr)
	fmt.Println(newArr)

	a := []string{
		"jason",
		"brock",
		"rowan",
	}
	modArray(a)
	fmt.Println(a)

	modifyingSlice()
}

// modify array takes a copy of the array and modifies it but original is unchanged
func modifyArr(arr [3]int) {
	arr[0] = 1000
	fmt.Println(arr)
}

func modifySlice(sl []int) {
	sl[0] = 1000000
}

func appSlice(sl []int) []int {
	sl[1] = 3030
	fmt.Printf("the address of main array is %p", sl)
	fmt.Printf("the address of main array is %p", &sl[0])
	sl = append(sl, 8080)
	return sl
}
