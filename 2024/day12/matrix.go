package main

import "fmt"

type Matrix [][]*Coordinate

func (farm Matrix) setPerimeters() {
	for rowIdx, rowValue := range farm {
		for colIdx, plot := range rowValue {
			perimeter := 0
			for _, direction := range directions {
				nextRowPosition := rowIdx + direction.rowPosition
				nextColPosition := colIdx + direction.colPosition
				if plot.hasBoundary(direction) || plot.Value != farm[nextRowPosition][nextColPosition].Value {
					perimeter++
				}
			}
			plot.Perimeter = perimeter
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
