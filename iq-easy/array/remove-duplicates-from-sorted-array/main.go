package main

import (
	"fmt"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/

func main() {

	// Write a function to remove duplicates from an incrementally sorted array by modifying it in place.
	// Then return the number of unique values in the array

	// The trick here is to filter the array without reallocating it
	// by chopping an element out of the middle of the slice, changing the size of the backing array
	// all elements after the chop will need to move to a new index, and while this can be
	// done by tracking the length of the slice as it shrinks, this isn't an efficient
	// way to filter the slice.

	// The way append works is that if all the elements fit into the current capacity, it
	// will use the same backing array as the original.
	// if the elements to be appended do not fit into the current capactity, append() will
	// allocate a new backing array and copy the contents over to it.
	// So to efficiently use append, it should be used in a way where all the elements fit
	// so that the slice can share the same backing array as the original and storage is reused.

	array := []int{1, 1, 2, 2, 2, 2, 2, 3, 3}
	result := removeDuplicates(array)
	fmt.Println(result)
	fmt.Println(array)
}

func removeDuplicates(nums []int) int {

	// create a slice which will share the same backing array as nums
	slice := nums[:0] // creates a 0 length slice retaining capacity of nums
	printSlice(slice) // len=0 cap=6 []

	counter := 0     // used to count unique numbers in the slice
	var previous int // used to store the previous value in the slice

	for i, v := range nums {
		if i == 0 { // there is nothing to compare the first value against
			// append the value, this will reuses the storage of nums
			// and modify its original contents
			slice = append(slice, v)
			counter++
			previous = v
		}

		if v == previous {
			// the item in the array is a duplicate of the previous value, so ignore it
			continue
		}

		if v != previous {
			// the value in the array is unique
			slice = append(slice, v)
			counter++
			previous = v
			continue
		}
	}

	return counter
}

// printSlice is a helper function to check slice information
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
