package main

import (
	"construct"
	"fmt"
)

func main() {

	// Declare 3D array structure
	x := 3
	y := 2
	z := 6

	// Create a 3D array
	triArray := construct.Build3DArray(x, y, z)

	// Flatten the 3D array (FYI: Below is my own declared class, its not a built in framework method)
	flatten3DArray := construct.Flatten3DArray(triArray, x, y, z)

	// Build a standard array
	memHoldNumbers := make([]int, 0)

	// Add some array elements
	memHoldNumbers = append(memHoldNumbers, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	// Flatten the 3D flat array with the standard array above (FYI: Below is my own declared class, its not a built in framework method)
	standardArray := construct.FlattenStandardArrays(flatten3DArray, memHoldNumbers)

	// Ass some space for clean logs
	fmt.Println()

	// Print the final standard results of the flattening
	fmt.Println("Standard flattened array: ", standardArray)

	// NOTE: I had some fun with the project, so I added a 3D structure

}
