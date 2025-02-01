package main

import "fmt"

func getPriceOfFence(input string) int {

	//read the string content into a matrix of rows and columns
	farm := NewMatrix(input)

	// loop through every row and column in the matrix
	// for each cell in the matrix, record the valid sides , up, right, down, left
	farm.locateSides()

	// loop through every row and column in the matrix
	// assign a distinct name to same shapes that are together
	farm.setDistinctRegionName()

	return farm.calculatePrice()
}

func getPriceOfFencePartTwo(input string) int {
	farm := NewMatrix(input)

	farm.locateDistinctSides()

	// assign the distinct sides to sides so the area can be calculated easily
	for _, rowValue := range farm {
		for _, plot := range rowValue {
			plot.Sides = plot.DistinctSides
		}
	}

	return farm.calculatePrice()
}

func main() {
	input := "AAAA\nBBCD\nBBCC\nEEEC"

	fmt.Println(getPriceOfFence(input))
	fmt.Println(getPriceOfFencePartTwo(input))
}
