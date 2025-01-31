package main

import "fmt"

type Matrix [][]*Coordinate

func (farm Matrix) locateSides() {
	for rowIdx, rowValue := range farm {
		for colIdx, plot := range rowValue {
			for _, direction := range directions {
				nextRowPosition := rowIdx + direction.rowPosition
				nextColPosition := colIdx + direction.colPosition
				if plot.hasBoundary(direction) || plot.Value != farm[nextRowPosition][nextColPosition].Value {
					plot.Sides[direction.Name] = true
				}
			}
		}
	}
}

func (farm Matrix) locateDistinctSides() {
	// you have to locate all the available sides first so you can remove the ones that don't have corners
	farm.locateSides()
	// you have to locate all the regions so that you only compare plots in the same region
	farm.setDistinctRegionName()
	for rowIdx, rowValue := range farm {
		for colIdx, plot := range rowValue {
			plot.DistinctSides = copyMap(plot.Sides)
			// you need to initialise a new map to keep track of the distinct sides
			// if you change the
			// checks the previous plots, left and up direction
			for _, direction := range []Direction{leftDirection, upDirection} {
				if plot.hasBoundary(direction) {
					continue
				}

				prevRow := rowIdx + direction.rowPosition
				prevCol := colIdx + direction.colPosition
				prevPlot := farm[prevRow][prevCol]
				if plot.Region != prevPlot.Region {
					continue
				}
				// if a side in the current plot exists in the previous plot,
				// it means that that side does not have a corner
				for _, side := range []string{UP, RIGHT, DOWN, LEFT} {
					if plot.Sides[side] && prevPlot.Sides[side] {
						plot.DistinctSides[side] = false
					}
				}

			}

		}
	}

}

func (farm Matrix) setDistinctRegionName() {
	counter := 0
	for _, rowValue := range farm {
		for _, plot := range rowValue {
			if plot.Region != "" {
				continue
			}
			regionName := fmt.Sprintf("%s-%d", plot.Value, counter)
			plot.NameAllValidNeighbors(regionName)
			counter++
		}
	}
}

func (farm Matrix) calculateArea() int {
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

func copyMap(original map[string]bool) map[string]bool {
	newMap := make(map[string]bool)
	for key, value := range original {
		newMap[key] = value
	}
	return newMap
}
