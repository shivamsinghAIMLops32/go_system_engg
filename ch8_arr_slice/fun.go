package main

import "fmt"

func modArray(arr []string) []string {
	arr[0] = "2020"
	arr = append(arr, "9090")
	return arr
}

func modifyingSlice() {
	sli := make([]byte, 3)
	fmt.Printf("addr: %p\n", sli)

	// sli has 3 elements with their 0 values like 0 as byte
	newSli := modME(sli)
	fmt.Printf("addr: %p of sli\n", sli)
	fmt.Printf("addr: %p after modifying\n", newSli)

	newSliW := make([]byte, 3, 5)
	// these two has same underlying array as they still have capacity
	// it's just newSliW's struct length field won't be updated so need to receive the new slice from modMEWithCapacity
	fmt.Printf("addr: %p of newSliW before modifying\n", newSliW)
	newSliW = modMEWithCapacity(newSliW)
	fmt.Printf("addr: %p of newSliW after modifying	\n", newSliW)

	// create a new slice and get the addres of it
	spl := createNEWSLice()
	fmt.Printf("addr: %v stored in spl\n", spl)
	fmt.Printf("addr: %p of spl\n", *spl)
}

func modME(sli []byte) []byte {
	// as sli has 3 element and 3 capacity here also aslo appednding would result in new underlying array
	sli = append(sli, 99)
	fmt.Printf("addr: %p\n", sli)
	return sli
}

func modMEWithCapacity(slo []byte) []byte {
	slo = append(slo, 100)
	return slo
}

func createNEWSLice() *[]int {
	it := []int{1, 2, 3}
	fmt.Printf("addr: %p\n", &it)
	fmt.Printf("addr: %p\n", &it[0])
	return &it
}
