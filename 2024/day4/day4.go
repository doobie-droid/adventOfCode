package main

import (
	"os"
	"strings"
)

type Coordinate struct {
	row       int
	col       int
	Value     rune
	rowBounds int
	colBounds int
}

func (point *Coordinate) isValidIndex(rowPosition, colPosition int) bool {
	if colPosition < 0 || rowPosition < 0 {
		return false
	}

	if rowPosition >= point.rowBounds || colPosition >= point.colBounds {
		return false
	}

	return true
}

func (point *Coordinate) FindXmasOccurrences(board [][]*Coordinate) int {
	word := "XMAS"
	actualOccurrences := 0

	movementPatterns := []struct {
		rowMovement int
		colMovement int
	}{
		{0, 1},
		{-1, 0},
		{1, 0},
		{0, -1},
		{1, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
	}

	for _, move := range movementPatterns {
		for i, char := range word {
			rowPosition := point.row + move.rowMovement*i
			colPosition := point.col + move.colMovement*i

			if !point.isValidIndex(rowPosition, colPosition) {
				break
			}
			if board[rowPosition][colPosition].Value != char {
				break
			}
			if char == 'S' {
				actualOccurrences++
			}

		}
	}
	return actualOccurrences
}

func findXmasInFile(fileName string) int {
	xmasCount := 0
	matrix := readStringsIntoRuneMatrix(fileName)
	for _, rowValue := range matrix {
		for _, colValue := range rowValue {
			if colValue.Value != 'X' {
				continue
			}
			xmasCount += colValue.FindXmasOccurrences(matrix)
		}
	}
	return xmasCount
}

func readStringsIntoRuneMatrix(fileName string) [][]*Coordinate {
	var planeOfCoordinates [][]*Coordinate
	rowBounds, colBounds := 0, 0

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
				row:       row,
				col:       column,
				Value:     letter,
				rowBounds: rowBounds,
				colBounds: colBounds,
			}
			rowOfCoordinates = append(rowOfCoordinates, coordinate)
		}
		planeOfCoordinates = append(planeOfCoordinates, rowOfCoordinates)
	}
	return planeOfCoordinates
}
