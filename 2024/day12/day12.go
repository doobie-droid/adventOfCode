package main

import (
	"os"
	"strings"
)

func getPriceOfFence(fileName string) int {
	// get string content from file
	fileContent := readFile(fileName)

	//read the string content into a matrix of rows and columns
	farm := getMatrix(fileContent)

	// loop through every row and column in the matrix
	// for each cell in the matrix, record the valid sides , up, right, down, left
	farm.locateSides()

	// loop through every row and column in the matrix
	// assign a distinct name to same shapes that are together
	farm.setDistinctRegionName()

	area := farm.calculateArea()
	return area
}

func getMatrix(fileContent string) Matrix {
	lines := strings.Split(fileContent, "\n")
	var matrix [][]*Coordinate
	for rowIdx, line := range lines {
		var row []*Coordinate
		for colIdx, char := range line {
			plant := string(char)
			coordinate := &Coordinate{
				Row:       rowIdx,
				Col:       colIdx,
				RowBounds: len(lines),
				ColBounds: len(line),
				Value:     plant,
				Matrix:    nil,
				Sides:     map[string]bool{},
			}
			row = append(row, coordinate)
		}
		matrix = append(matrix, row)
	}
	for rowIdx, rowValue := range matrix {
		for colIdx := range rowValue {
			matrix[rowIdx][colIdx].Matrix = matrix
		}
	}
	return matrix
}

func readFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(file)
}
