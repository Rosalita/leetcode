package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortBoxes(t *testing.T) {

	tests := map[string]struct {
		input  [][]int // represents each element of slice {quantity, units}
		result [][]int
	}{
		"can sort a slice from largest units to smallest": {
			input:  [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}},
			result: [][]int{{5, 10}, {3, 9}, {4, 7}, {2, 5}},
		},
		"sorted slice remains sorted": {
			input:  [][]int{{1, 3}, {2, 2}, {3, 1}},
			result: [][]int{{1, 3}, {2, 2}, {3, 1}},
		},
	}

	for _, test := range tests {
		result := sortBoxes(test.input)
		assert.Equal(t, test.result, result)
	}
}

func TestMaximumUnits(t *testing.T) {

	tests := map[string]struct {
		inputBoxTypes  [][]int
		inputTruckSize int
		output         int
	}{
		"provided test case 1": {
			inputBoxTypes:  [][]int{{1, 3}, {2, 2}, {3, 1}},
			inputTruckSize: 4,
			output:         8,
		},
		"provided test case 2": {
			inputBoxTypes:  [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}},
			inputTruckSize: 10,
			output:         91,
		},
	}

	for _, test := range tests {
		result := maximumUnits(test.inputBoxTypes, test.inputTruckSize)
		assert.Equal(t, test.output, result)
	}
}
