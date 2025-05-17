package main

import "fmt"

func main() {
	// arrays
	// array is a collection of elements of same type
	// array is a fixed size data structure
	// array is a value type
	// array is a contiguous block of memory
	// array is a collection of elements of same type
	// array is a fixed size data structure
	// array is a value type
	// array is a contiguous block of memory

	var arr1 [5]int // declare an array of size 5
	arr1[0] = 1     // assign value to first element
	arr1[1] = 2     // assign value to second element
	arr1[2] = 3     // assign value to third element
	arr1[3] = 4     // assign value to fourth element
	arr1[4] = 5     // assign value to fifth element

	var arr2 [5]int = [5]int{1, 2, 3, 4, 5} // declare and initialize an array

	fmt.Println(arr1[0]) // print first element of arr1
	fmt.Println(arr2[0]) // print first element of arr2

	fmt.Println(len(arr1)) // print length of arr1
	fmt.Println(len(arr2)) // print length of arr2

	fmt.Println(arr1) // print arr1
	fmt.Println(arr2) // print arr2

	var vals [4]bool // default false values will be populated
	vals[2] = true
	fmt.Println(vals) // print vals

	// for string the zeroed values is empty string
	var vals2 [4]string // default empty string values will be populated
	fmt.Println(vals2)  // print vals2
	// for float the zeroed values is 0.0, int -> 0, bool -> false

	//nums := [3]int{1, 2, 3} // declare and initialize an array with custom val
	//fmt.Println(nums)       // print nums

	// 2D arrays in go
	nums := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums)

	// arayys are used  normal without slice where size is predictable
	// this helps in memory optimization
	// constant time access to elements

}
