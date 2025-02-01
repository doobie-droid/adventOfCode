package main

import (
	"fmt"
	"strings"
)

type Matrix [][]*Plot

func (farm Matrix) locateSides() {
	// Iterate Through Each Plot in the Matrix
	for rowIdx, rowValue := range farm {
		for colIdx, plot := range rowValue {
			// For each plot, check all four directions (Top, Right, Bottom, Left).
			// Check if there is a side in any direction
			// There is a side if
			// a. there is a boundary in the next direction
			// b. The plot in the next direction does not have the same plant as the current plot
			for _, direction := range directions {
				nextRowPosition := rowIdx + direction.rowOffset
				nextColPosition := colIdx + direction.colOffset

				if plot.hasBoundary(direction) || plot.Value != farm[nextRowPosition][nextColPosition].Value {
					plot.Sides[direction.Name] = true
				}

			}
		}
	}
}

func copyMap(original map[string]bool) map[string]bool {
	newMap := make(map[string]bool)
	for key, value := range original {
		newMap[key] = value
	}
	return newMap
}

func (farm Matrix) locateDistinctSides() {
	// you have to locate all the available sides first so you can remove the ones that don't have corners
	farm.locateSides()
	// you have to locate all the regions so that you only compare plots in the same region
	farm.setDistinctRegionName()
	for rowIdx, rowValue := range farm {
		for colIdx, plot := range rowValue {
			// you need to initialize a new map to keep track of the distinct sides
			// so that the old map that records the actual sides of a plot is not changed
			plot.DistinctSides = copyMap(plot.Sides)
			// for each plot, we check the previous plot in the left direction and in the up direction
			// if there is no plot in either direction, all the sides are distinct
			// if there is a plot in either direction, the sides that appear in the current plot, but do not appear in the previous plot
			// ARE DISTINCT
			for _, direction := range []Direction{leftDirection, upDirection} {
				if plot.hasBoundary(direction) {
					continue
				}
				prevRow := rowIdx + direction.rowOffset
				prevCol := colIdx + direction.colOffset
				prevPlot := farm[prevRow][prevCol]
				if plot.Region != prevPlot.Region {
					continue
				}
				// if a side in the current plot exists in the previous plot,
				// it means that that side is not distinct but rather a continuation of the previous side
				for _, side := range []string{UP, RIGHT, DOWN, LEFT} {
					if plot.Sides[side] && prevPlot.Sides[side] {
						plot.DistinctSides[side] = false
					}
				}

			}

		}
	}

}

func NewMatrix(fileContent string) Matrix {
	// Split the content into new lines based on the new line separator
	lines := strings.Split(fileContent, "\n")
	// create a matrix of plot
	var matrix [][]*Plot
	// loop through the rows
	for rowIdx, line := range lines {
		// loop through the columns
		var row []*Plot
		// for each individual plant, create a plot and attach it to the farm
		for colIdx, char := range line {
			plant := string(char)
			coordinate := &Plot{
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
	// after populating the matrix with plots
	// loop through the matrix once again and put a pointer to the farm in each plot
	// so that each plot can interact with the entirety of the farm
	for rowIdx, rowValue := range matrix {
		for colIdx := range rowValue {
			matrix[rowIdx][colIdx].Matrix = matrix
		}
	}
	return matrix
}

func (farm Matrix) setDistinctRegionName() {
	counter := 0
	for _, rowValue := range farm {
		for _, plot := range rowValue {
			if plot.Region != "" {
				continue
			}
			regionName := fmt.Sprintf("%s-%d", plot.Value, counter)
			plot.AddPlotToRegion(regionName)
			counter++
		}
	}
}

func (farm Matrix) calculatePrice() int {
	regionToShapeMap := map[string]*Shape{}

	// loop through every row and column
	// update a map, that has matches a the name of a region to a shape [update the perimeter and area part]
	for _, row := range farm {
		for _, point := range row {
			shape, exists := regionToShapeMap[point.Region]
			if exists {
				shape.Area++
				shape.Perimeter += point.CalculatePerimeter()
			} else {
				shape = &Shape{
					Perimeter: point.CalculatePerimeter(),
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
	return sum
}
