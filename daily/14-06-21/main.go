package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge-2021/604/week-2-june-8th-june-14th/3778/

func main() {

	boxes := [][]int{{1, 3}, {2, 2}, {3, 1}}

	result := maximumUnits(boxes, 4)

	fmt.Println(result)
}

func maximumUnits(boxTypes [][]int, truckSize int) int {

	// This is an example of a greedy algorithm making the optimised choice
	// at each stage to reach (hopefully) an optimal solution.

	sortedBoxes := sortBoxes(boxTypes)

	totalUnits := 0

	for _, box := range sortedBoxes {

		availableBoxQuantity := box[0]
		unitsPerBox := box[1]

		if truckSize > availableBoxQuantity {
			// there is space on the truck for all boxes of this type
			units := availableBoxQuantity * unitsPerBox
			totalUnits += units
			truckSize -= availableBoxQuantity
			continue

		} else {
			// not all boxes of this type will fit on the truck
			units := truckSize * unitsPerBox
			totalUnits += units
			return totalUnits
		}
	}

	return totalUnits
}

func sortBoxes(boxes [][]int) [][]int {
	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i][1] > boxes[j][1]
	})
	return boxes
}
