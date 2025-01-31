package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(readValuesIntoArrayOfInts("day7Sample_test.txt"))
}

func sumOfBridgeRepairs(fileName string) int {
	arrayOfOutcomeAndInput := readValuesIntoArrayOfInts(fileName)
	solutionArray := [][]int{}
	for _, innerArray := range arrayOfOutcomeAndInput {
		if isValidArray(0, innerArray[0], 0, innerArray[1:]) {
			solutionArray = append(solutionArray, innerArray)
		}
	}
	return sumArray(solutionArray)
}
func isValidArray(currentSum, desiredOutcome, currentPointer int, array []int) bool {
	if currentSum > desiredOutcome {
		return false
	}
	if currentSum == desiredOutcome {
		return true
	}

	if currentPointer == len(array) {
		return false
	}

	return isValidArray(currentSum+array[currentPointer], desiredOutcome, currentPointer+1, array) || isValidArray(currentSum*array[currentPointer], desiredOutcome, currentPointer+1, array)
}

func sumOfBridgeRepairsWithConcatenation(fileName string) int {
	arrayOfOutcomeAndInput := readValuesIntoArrayOfInts(fileName)
	solutionArray := [][]int{}
	for _, innerArray := range arrayOfOutcomeAndInput {
		if isValidArrayWithConcatenation(0, innerArray[0], 0, innerArray[1:]) {
			solutionArray = append(solutionArray, innerArray)
		}
	}
	return sumArray(solutionArray)
}

func isValidArrayWithConcatenation(currentSum, desiredOutcome, currentPointer int, array []int) bool {
	if currentSum > desiredOutcome {
		return false
	}
	if currentSum == desiredOutcome {
		return true
	}

	if currentPointer == len(array) {
		return false
	}

	concatenatedNumberAsString := fmt.Sprintf("%d%d", currentSum, array[currentPointer])
	concatenatedNumber, err := strconv.Atoi(concatenatedNumberAsString)
	if err != nil {
		return false
	}
	return isValidArrayWithConcatenation(currentSum+array[currentPointer], desiredOutcome, currentPointer+1, array) || isValidArrayWithConcatenation(currentSum*array[currentPointer], desiredOutcome, currentPointer+1, array) || isValidArrayWithConcatenation(concatenatedNumber, desiredOutcome, currentPointer+1, array)
}
func sumArray(array [][]int) (sum int) {
	for _, innerArray := range array {
		sum += innerArray[0]
	}
	return sum
}
func readValuesIntoArrayOfInts(fileName string) [][]int {
	var solutionArray [][]int
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		innerArray := []any{}
		intInnerArray := []int{}
		outcomeAndInput := strings.Split(line, ":")
		outcome := outcomeAndInput[0]
		innerArray = append(innerArray, outcome)
		input := outcomeAndInput[1]
		inputs := strings.Split(input, " ")
		for _, digit := range inputs {
			innerArray = append(innerArray, digit)
		}
		for _, digit := range innerArray {
			actualDigit, _ := strconv.Atoi(digit.(string))
			if actualDigit == 0 {
				continue
			}
			intInnerArray = append(intInnerArray, actualDigit)
		}
		solutionArray = append(solutionArray, intInnerArray)
	}
	return solutionArray
}
