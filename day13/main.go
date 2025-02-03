package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isInteger(float float64) bool {
	value1 := float
	value2 := math.Round(float)
	return math.Abs(value1-value2) < 0.001
}

func findMinimumNumberOfTokens(fileName string) int {
	var tokenCount int
	fileContent := readFile(fileName)
	lines := strings.Split(fileContent, "\n\n")

	for _, line := range lines {
		inputMatrix, outputMatrix := newMatrix(line)
		inv, _ := inputMatrix.getInverse()
		output := inv.multiply(outputMatrix)
		tokenACount, tokenBCount := output[0][0], output[1][0]
		if isInteger(tokenACount) && isInteger(tokenBCount) {
			tokenCount += int(math.Round(tokenACount*3) + math.Round(tokenBCount*1))
		}

	}
	return tokenCount
}

func findMinimumNumberOfTokensPartTwo(fileName string) int {
	var tokenCount int
	fileContent := readFile(fileName)
	lines := strings.Split(fileContent, "\n\n")

	for _, line := range lines {
		inputMatrix, outputMatrix := newMatrix(line)
		inv, _ := inputMatrix.getInverse()
		// update to accommodate part 2
		outputMatrix[0][0] = outputMatrix[0][0] + 10000000000000
		outputMatrix[1][0] = outputMatrix[1][0] + 10000000000000
		output := inv.multiply(outputMatrix)
		tokenACount, tokenBCount := output[0][0], output[1][0]
		if isInteger(tokenACount) && isInteger(tokenBCount) {
			tokenCount += int(math.Round(tokenACount*3) + math.Round(tokenBCount*1))
		}

	}
	return tokenCount
}

func main() {
	tokens := findMinimumNumberOfTokens("day13Sample_test.txt")
	fmt.Println(tokens)
}

func readFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(file)
}

/*
* input  [2,3,4,5]
* output [2,3,4,5]
 */
func convertStringArrayToIntArray(input []string) (output []int) {
	for _, value := range input {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		output = append(output, intValue)

	}

	return output
}
