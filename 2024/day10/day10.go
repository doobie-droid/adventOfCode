package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	Row                      int
	Col                      int
	RowBound                 int
	ColBound                 int
	Value                    int
	Matrix                   [][]*Coordinate
	CountOfVisitsToTrailHead map[string]int
}

func (point *Coordinate) hasPassedEnd(rowPosition, colPosition int) bool {
	if colPosition < 0 || rowPosition < 0 {
		return true
	}

	if rowPosition > point.RowBound-1 || colPosition > point.ColBound-1 {
		return true
	}

	return false
}

func (point *Coordinate) CalculateNoOfTrailHeads(rowPosition, colPosition, previousValue int) int {
	var directions = []struct {
		rowMovement int
		colMovement int
	}{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	if point.hasPassedEnd(rowPosition, colPosition) {
		return 0
	}

	currentValue := point.Matrix[rowPosition][colPosition].Value
	if currentValue-previousValue != 1 {
		return 0
	}

	if currentValue == 9 {
		address := fmt.Sprintf("row_%d_col_%d", rowPosition, colPosition)
		if point.CountOfVisitsToTrailHead[address] > 0 {
			point.CountOfVisitsToTrailHead[address]++
			return 0
		} else {
			point.CountOfVisitsToTrailHead[address]++
			return 1
		}
	}

	trailHeads := 0
	for _, direction := range directions {
		trailHeads += point.CalculateNoOfTrailHeads(
			rowPosition+direction.rowMovement,
			colPosition+direction.colMovement,
			currentValue,
		)
	}

	return trailHeads
}

func findSumOfTrailHeads(fileName string) int {
	matrix := readStringsIntoRuneMatrix(fileName)
	for row, rowValue := range matrix {
		for col, colValue := range rowValue {
			if colValue.Value != 0 {
				continue
			}
			colValue.CalculateNoOfTrailHeads(row, col, -1)
		}
	}
	return sumTrailHeads(matrix)
}

func findSumOfTrailHeadsPartTwo(fileName string) int {
	matrix := readStringsIntoRuneMatrix(fileName)
	for row, rowValue := range matrix {
		for col, colValue := range rowValue {
			if colValue.Value != 0 {
				continue
			}
			colValue.CalculateNoOfTrailHeads(row, col, -1)
		}
	}
	return sumTrailHeadsRatings(matrix)
}

func sumTrailHeadsRatings(matrix [][]*Coordinate) int {
	sumOfRatings := 0
	for _, rowValue := range matrix {
		for _, colValue := range rowValue {
			ratings := 0
			for _, frequency := range colValue.CountOfVisitsToTrailHead {
				ratings += frequency
			}
			sumOfRatings += ratings
		}
	}
	return sumOfRatings

}
func sumTrailHeads(matrix [][]*Coordinate) int {
	sumOfTrailHeads := 0
	for _, rowValue := range matrix {
		for _, colValue := range rowValue {
			sumOfTrailHeads += len(colValue.CountOfVisitsToTrailHead)
		}
	}
	return sumOfTrailHeads
}

func readFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func readStringsIntoRuneMatrix(fileName string) [][]*Coordinate {
	var planeOfCoordinates [][]*Coordinate
	rowBounds, colBounds := 0, 0

	contents := readFile(fileName)
	lines := strings.Split(contents, "\n")
	for row, line := range lines {
		rowBounds = len(lines)
		rowOfCoordinates := []*Coordinate{}
		for column, letter := range line {
			colBounds = len(line)
			integer, err := strconv.Atoi(string(letter))
			if err != nil {
				panic(err)
			}
			coordinate := &Coordinate{
				Row:                      row,
				Col:                      column,
				Value:                    integer,
				RowBound:                 rowBounds,
				ColBound:                 colBounds,
				Matrix:                   nil,
				CountOfVisitsToTrailHead: map[string]int{},
			}
			rowOfCoordinates = append(rowOfCoordinates, coordinate)
		}
		planeOfCoordinates = append(planeOfCoordinates, rowOfCoordinates)
	}

	for row, line := range lines {
		for column := range line {
			planeOfCoordinates[row][column].Matrix = planeOfCoordinates
		}
	}

	return planeOfCoordinates
}
