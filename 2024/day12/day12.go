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
	// for each cell in the matrix, assign a perimeter
	farm.setPerimeters()

	// loop through every row and column in the matrix
	// assign a distinct name to same shapes that are together
	farm.setDistinctRegionName()

	// loop through every row and column
	// update a map, that has matches a name to a shape [update the perimeter and area part]
	regionToShapeMap := map[string]*Shape{}

	for _, row := range farm {
		for _, point := range row {
			shape, exists := regionToShapeMap[point.Region]
			if exists {
				shape.Area++
				shape.Perimeter += point.Perimeter

			} else {
				shape = &Shape{
					Perimeter: point.Perimeter,
					Area:      1,
				}
			}
			regionToShapeMap[point.Region] = shape
		}
	}

	//loop through map containing name of shapes
	// add sum of perimeters and areas
	sum := 0
	for _, shape := range regionToShapeMap {
		sum += (shape.Perimeter * shape.Area)
	}
	// return said sum
	return sum
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
