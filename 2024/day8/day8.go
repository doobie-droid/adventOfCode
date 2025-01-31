package main

import (
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	Row             int
	Col             int
	RowBound        int
	ColBound        int
	InitRow         int
	InitCol         int
	Value           rune
	VisitedLocation map[string]bool
}
type Displacement struct {
	RowDistance int
	ColDistance int
}

func (point *Coordinate) isValidIndex(rowPosition, colPosition int) bool {
	if colPosition < 0 || rowPosition < 0 {
		return false
	}

	if rowPosition >= point.RowBound || colPosition >= point.ColBound {
		return false
	}

	return true
}

func (point *Coordinate) hasVisited(rowPosition, colPosition int) bool {
	positionKey := fmt.Sprintf("row_%d_col_%d", rowPosition, colPosition)
	return point.VisitedLocation[positionKey]
}

func (point *Coordinate) Visit(rowPosition, colPosition int) {
	positionKey := fmt.Sprintf("row_%d_col_%d", rowPosition, colPosition)
	point.VisitedLocation[positionKey] = true
}

func (point *Coordinate) findDistance(otherPoint *Coordinate) *Displacement {
	rowDistance := point.Row - otherPoint.Row
	colDistance := point.Col - otherPoint.Col
	return &Displacement{
		RowDistance: rowDistance * -1,
		ColDistance: colDistance * -1,
	}
}

func (displacement Displacement) MultiplyBy2() Displacement {
	return Displacement{
		RowDistance: displacement.RowDistance * 2,
		ColDistance: displacement.ColDistance * 2,
	}
}
func (displacement Displacement) Invert() Displacement {
	return Displacement{
		RowDistance: displacement.RowDistance * -1,
		ColDistance: displacement.ColDistance * -1,
	}
}

func (point *Coordinate) ResetDistance() *Coordinate {
	point.Row = point.InitRow
	point.Col = point.InitCol
	return point
}

func findNumberOfAntinodes(matrix [][]*Coordinate) int {
	hashMapOfAntennaeToLocations := map[rune][]*Coordinate{}
	validAntinodes := 0
	for _, rowValue := range matrix {
		for _, point := range rowValue {
			if point.Value == '.' {
				continue
			}
			coordinateArray, exists := hashMapOfAntennaeToLocations[point.Value]
			if !exists {
				hashMapOfAntennaeToLocations[point.Value] = []*Coordinate{point}
				continue
			}
			for _, otherPoint := range coordinateArray {
				displacement := point.findDistance(otherPoint)
				proposedRowPosition := point.Row + displacement.MultiplyBy2().RowDistance
				proposedColPosition := point.Col + displacement.MultiplyBy2().ColDistance

				if point.isValidIndex(proposedRowPosition, proposedColPosition) {
					if !point.hasVisited(proposedRowPosition, proposedColPosition) {
						validAntinodes++
						point.Visit(proposedRowPosition, proposedColPosition)
					}
				}
				proposedRowPosition = point.Row + displacement.Invert().RowDistance
				proposedColPosition = point.Col + displacement.Invert().ColDistance
				if point.isValidIndex(proposedRowPosition, proposedColPosition) {
					if !point.hasVisited(proposedRowPosition, proposedColPosition) {
						validAntinodes++
						point.Visit(proposedRowPosition, proposedColPosition)
					}
				}
			}
			coordinateArray = append(coordinateArray, point)
			hashMapOfAntennaeToLocations[point.Value] = coordinateArray
		}
	}
	return validAntinodes
}

func findNumberOfAntinodesPartTwo(matrix [][]*Coordinate) int {
	hashMapOfAntennaeToLocations := map[rune][]*Coordinate{}
	for _, rowValue := range matrix {
		for _, point := range rowValue {
			if point.Value == '.' {
				continue
			}
			coordinateArray, exists := hashMapOfAntennaeToLocations[point.Value]
			if !exists {
				hashMapOfAntennaeToLocations[point.Value] = []*Coordinate{point}
				continue
			}
			for _, otherPoint := range coordinateArray {
				displacement := point.findDistance(otherPoint)
				for point.isValidIndex(point.Row, point.Col) {
					point.Visit(point.Row, point.Col)
					point.Row = point.Row + displacement.RowDistance
					point.Col = point.Col + displacement.ColDistance
				}
				point.ResetDistance()
				reverseDisplacement := displacement.Invert()
				for point.isValidIndex(point.Row, point.Col) {
					point.Visit(point.Row, point.Col)
					point.Row = point.Row + reverseDisplacement.RowDistance
					point.Col = point.Col + reverseDisplacement.ColDistance
				}
				point.ResetDistance()
			}
			coordinateArray = append(coordinateArray, point)
			hashMapOfAntennaeToLocations[point.Value] = coordinateArray
		}
	}
	return len(matrix[0][0].VisitedLocation)
}

func countAntinodes(fileName string) int {
	matrix := readStringsIntoRuneMatrix(fileName)
	return findNumberOfAntinodes(matrix)
}

func countAntinodesPartTwo(fileName string) int {
	matrix := readStringsIntoRuneMatrix(fileName)
	return findNumberOfAntinodesPartTwo(matrix)
}
func main() {

}

func readStringsIntoRuneMatrix(fileName string) [][]*Coordinate {
	var planeOfCoordinates [][]*Coordinate
	rowBounds, colBounds := 0, 0
	locatedAntinodePositions := map[string]bool{}

	contents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	for row, line := range lines {
		rowBounds = len(lines)
		rowOfCoordinates := []*Coordinate{}
		for column, letter := range line {
			colBounds = len(line)
			coordinate := &Coordinate{
				Row:             row,
				Col:             column,
				Value:           letter,
				RowBound:        rowBounds,
				ColBound:        colBounds,
				InitRow:         row,
				InitCol:         column,
				VisitedLocation: locatedAntinodePositions,
			}
			rowOfCoordinates = append(rowOfCoordinates, coordinate)
		}
		planeOfCoordinates = append(planeOfCoordinates, rowOfCoordinates)
	}
	return planeOfCoordinates
}
