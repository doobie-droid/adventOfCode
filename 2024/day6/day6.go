package main

import (
	"fmt"
	"os"
	"strings"
)

//Bug that took time was forgetting to reset the position of the coordinate to its initial when using it multiple times

type Coordinate struct {
	row       int
	col       int
	Value     rune
	rowBounds int
	colBounds int
}

func (point *Coordinate) hasReachedEnd() bool {
	if point.col <= 0 || point.row <= 0 {
		return true
	}

	if point.row >= point.rowBounds-1 || point.col >= point.colBounds-1 {
		return true
	}

	return false
}

func (point *Coordinate) hasObstacle(rowPosition, colPosition int, matrix [][]*Coordinate) bool {
	foundPoint := matrix[rowPosition][colPosition]
	return foundPoint.Value == '#'
}

func (object *Coordinate) exitMaze(matrix [][]*Coordinate) (outputMatrix [][]*Coordinate) {
	currentDirection := 0
	movement := []struct {
		rowMovement int
		colMovement int
	}{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	direction := movement[currentDirection]
	for !object.hasReachedEnd() {
		rowPosition := object.row + direction.rowMovement
		colPosition := object.col + direction.colMovement
		if object.hasObstacle(rowPosition, colPosition, matrix) {
			currentDirection = (currentDirection + 1) % 4
			direction = movement[currentDirection]
			continue
		}
		previousRowPosition := object.row
		previousColPosition := object.col
		object.row = rowPosition
		object.col = colPosition
		matrix[previousRowPosition][previousColPosition].Value = 'X'
	}
	matrix[object.row][object.col].Value = 'X'
	return matrix
}

func (object *Coordinate) exitMazeWithoutTrace(matrix [][]*Coordinate) (exitStatus bool) {
	locationsVisited := 0
	hashMapOfVisitedLocationsAndBearings := map[string]bool{}
	currentDirection := 0
	movement := []struct {
		rowMovement int
		colMovement int
	}{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	direction := movement[currentDirection]
	for !object.hasReachedEnd() {
		rowPosition := object.row + direction.rowMovement
		colPosition := object.col + direction.colMovement
		if object.hasObstacle(rowPosition, colPosition, matrix) {
			locationWithBearing := fmt.Sprintf("row_%d_column_%d_%d%d", object.row, object.col, direction.rowMovement, direction.colMovement)
			if hashMapOfVisitedLocationsAndBearings[locationWithBearing] {
				return false
			}
			hashMapOfVisitedLocationsAndBearings[locationWithBearing] = true
			currentDirection = (currentDirection + 1) % 4
			direction = movement[currentDirection]
			continue
		}
		locationsVisited++
		object.row = rowPosition
		object.col = colPosition
	}
	return true
}

func (object *Coordinate) findPossibleObstructions(matrixOfTraveledPath [][]*Coordinate, matrix [][]*Coordinate) int {
	validObstructions := 0
	for row, rowValue := range matrix {
		for col, _ := range rowValue {
			if matrixOfTraveledPath[row][col].Value != 'X' {
				continue
			}
			matrix[row][col].Value = '#'
			initialRow, initialColumn := object.row, object.col
			if !object.exitMazeWithoutTrace(matrix) {
				validObstructions++
			}
			object.row, object.col = initialRow, initialColumn
			matrix[row][col].Value = '.'
		}
	}
	return validObstructions

}

func findDistinctPositionForObstruction(fileName string) int {
	matrix, obj := readStringsIntoRuneMatrix(fileName)
	matrixWithVisitedSpots := obj.exitMaze(matrix)
	freshMatrix, object := readStringsIntoRuneMatrix(fileName)
	return object.findPossibleObstructions(matrixWithVisitedSpots, freshMatrix)
}
func findDistinctPositions(fileName string) int {
	matrix, object := readStringsIntoRuneMatrix(fileName)
	matrix = object.exitMaze(matrix)

	return findDistinctVisitedSpots(matrix)
}

func findDistinctVisitedSpots(matrix [][]*Coordinate) (visitedSpots int) {
	for _, rowVal := range matrix {
		for _, colVal := range rowVal {
			fmt.Print(string(colVal.Value))
			if colVal.Value == 'X' {
				visitedSpots++
			}
		}
	}
	return visitedSpots
}

func readStringsIntoRuneMatrix(fileName string) (matrix [][]*Coordinate, object *Coordinate) {
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
			if letter == '^' {
				object = coordinate
			}
			rowOfCoordinates = append(rowOfCoordinates, coordinate)
		}
		matrix = append(matrix, rowOfCoordinates)
	}
	return matrix, object
}
