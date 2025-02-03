package main

import (
	"fmt"
	"regexp"
)

type Matrix[T int | float64] [][]T

func (matrix Matrix[int]) isSquare() bool {
	noOfRows := len(matrix)
	for _, row := range matrix {
		noOfColumns := len(row)
		if noOfColumns != noOfRows {
			return false
		}
	}
	return true
}

// Gets determinant of a 2 X 2 matrix
func (input Matrix[int]) getDeterminant() (int, error) {
	if !input.isSquare() {
		return 0, fmt.Errorf("can only get determinant of square matrices")
	}
	if len(input) != 2 {
		return 0, fmt.Errorf("can only get determinant of 2 X 2 matrices")
	}
	return (input[0][0] * input[1][1]) - (input[0][1] * input[1][0]), nil
}

// Get adjunct of a 2 X 2 Matrix
func (input Matrix[int]) getAdjunct() (output Matrix[int], err error) {
	if !input.isSquare() {
		return nil, fmt.Errorf("you can only find transposes of square matrices")

	}

	matrix := [][]int{{0, 0}, {0, 0}}
	matrix[0][0] = input[1][1]
	matrix[0][1] = -1 * input[0][1]
	matrix[1][0] = -1 * input[1][0]
	matrix[1][1] = input[0][0]

	return matrix, nil
}

func (input Matrix[int]) getInverse() (outputMatrix Matrix[float64], err error) {
	if !input.isSquare() {
		return nil, fmt.Errorf("you can only find transposes of square matrices")
	}
	det, _ := input.getDeterminant()
	adj, _ := input.getAdjunct()
	for row, rowVal := range input {
		matrixRow := []float64{}
		for col := range rowVal {
			resultVal := float64(adj[row][col]) / float64(det)
			matrixRow = append(matrixRow, resultVal)
		}
		outputMatrix = append(outputMatrix, matrixRow)
	}

	return outputMatrix, err
}

func (input1 Matrix[float32]) multiply(input2 Matrix[int]) (output Matrix[float32]) {
	output = [][]float32{{0}, {0}}
	output[0][0] = (input1[0][0] * float32(input2[0][0])) + (input1[0][1] * float32(input2[1][0]))
	output[1][0] = (input1[1][0] * float32(input2[0][0])) + (input1[1][1] * float32(input2[1][0]))
	return output
}

/*
* line Button A: X+94, Y+34\nButton B: X+22, Y+67\nPrize: X=8400, Y=5400
* inputMatrix  [[94,22],[34,67]]
* outputMatrix [[8400],[5400]]
 */
func newMatrix(line string) (inputMatrix Matrix[int], outputMatrix Matrix[int]) {

	pattern := `X\+(\d+)[^X]*Y\+(\d+)[^X]*X\+(\d+)[^X]*Y\+(\d+)[^X]*X=(\d+)[^X]*Y=(\d+)`
	reg := regexp.MustCompile(pattern)
	matches := reg.FindAllStringSubmatch(line, -1)
	intMatrix := convertStringArrayToIntArray(matches[0][1:])
	inputMatrix = append(inputMatrix, []int{intMatrix[0], intMatrix[2]})
	inputMatrix = append(inputMatrix, []int{intMatrix[1], intMatrix[3]})
	outputMatrix = append(outputMatrix, []int{intMatrix[4]})
	outputMatrix = append(outputMatrix, []int{intMatrix[5]})

	return inputMatrix, outputMatrix
}
