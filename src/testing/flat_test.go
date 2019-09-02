package testing__test

import (
	"construct"
	"fmt"
	"testing"
)

func TestFlattenStandardArrays(t *testing.T) {
	// Create a 3D array
	tryArray := construct.Build3DArray(3, 2, 6)

	// Flatten the 3D array
	flatten3DArray := construct.Flatten3DArray(tryArray, 3, 2, 6)

	// Build a standard array
	memHoldNumbers := make([]int, 0)

	// Add some array elements
	memHoldNumbers = append(memHoldNumbers, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	// Test array compare
	compare := make([]int, 0)
	compare = append(compare, 0, 0, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 6, 1, 1, 2, 3, 4, 5, 6, 2, 3, 4, 5, 6, 7, 2, 2, 3, 4, 5, 6, 7, 3, 4, 5, 6, 7, 8, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	// Flatten the 3D flat array with the standard array above
	standardArray := construct.FlattenStandardArrays(flatten3DArray, memHoldNumbers)

	arraysEqual := ArraysEqual(compare, standardArray)

	// Ass some space for clean logs
	fmt.Println()

	// Print the final standard results of the flattening
	fmt.Println("Standard flattened array: ", standardArray)

	// Ass some space for clean logs
	fmt.Println()

	if !arraysEqual {
		t.Errorf("Expected result ** %v ** however the recieved result is: %v", compare, standardArray)
	}

}

func ArraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
