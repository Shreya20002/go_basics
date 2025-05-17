package main

import (
	"fmt"
	"slices" // inbuilt lib package
)

// slice -> dynamic
// most used construct in go
// + useful methods

func main() {
	// unitialized slice is nil (Null in other languages)
	//var nums []int
	//fmt.Println(nums)
	//fmt.Println(nums == nil)
	//fmt.Println(len(nums)) // 0

	// make inbuilt function
	// cap inbuilt fn to  check -> maximum num of elements that can fit
	//var nums = make([]int, 2, 5) // 2 is length, 5 is capacity -> zeroed vals are put in length-> do keep it 0 most times ([]int, 0, 5)
	//fmt.Println(cap(nums))
	//fmt.Println(len(nums))
	//fmt.Println(nums)
	//nums = append(nums, 1)
	//nums = append(nums, 2)
	//nums = append(nums, 3)
	//nums = append(nums, 4)
	// capacity doubled and resized to 10 (algo)

	//nums := []int{}

	//fmt.Println(nums)
	//fmt.Println(cap(nums))
	//fmt.Println(len(nums))

	// make(type, length, capacity)
	// length is what is actually present now , capacity is what can be max occupation possible

	// copy function

	//var nums = make([]int, 0, 5)
	//nums = append(nums, 2)
	//var nums2 = make([]int, len(nums))

	//nums = append(nums, 2)
	//copy(nums2, nums)
	//fmt.Println(nums, nums2)

	// slice operator
	var nums = []int{1, 2, 3}
	// same as python slice operstor
	// nums[from:to] -> from inclusive, to exclusive
	fmt.Println(nums[0:2]) // 1,2

	// slice
	var nums1 = []int{1, 2}
	var nums2 = []int{1, 2}

	fmt.Println(slices.Equal(nums1, nums2))

	//  2D slices
	var nums3 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums3)

}
