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
