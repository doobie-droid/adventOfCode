package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ans := readFilesIntoArraysOfUpdates("day5Sample_test.txt")
	red := sumMiddleElementsInArray(ans)
	fmt.Println(red)
}

func findSumOfMiddleUpdates(fileName string) int {
	hashMapOfValidPositions := readFilesIntoHashMapOfValidEntries(fileName)
	arrayOfUpdates := readFilesIntoArraysOfUpdates(fileName)
	arrayOfUpdates, _ = filterOffInvalidArrays(arrayOfUpdates, hashMapOfValidPositions)

	return sumMiddleElementsInArray(arrayOfUpdates)
}

func findSumOfMiddleUpdatesForIncorrectOrder(fileName string) int {
	hashMapOfValidPositions := readFilesIntoHashMapOfValidEntries(fileName)
	arrayOfUpdates := readFilesIntoArraysOfUpdates(fileName)
	_, incorrectArray := filterOffInvalidArrays(arrayOfUpdates, hashMapOfValidPositions)
	arrayInCorrectOrder := reorderArray(incorrectArray, hashMapOfValidPositions)
	fmt.Println(arrayInCorrectOrder)
	return sumMiddleElementsInArray(arrayInCorrectOrder)
}

func reorderArray(array [][]int, checkArrangementExists map[string]bool) (solutionArray [][]int) {

	for _, miniArray := range array {
		swapHappened := true
		for swapHappened {
			swapHappened = false
			for index := 0; index < len(miniArray)-1; index++ {
				arrangement := fmt.Sprintf("%d|%d", miniArray[index], miniArray[index+1])
				if !checkArrangementExists[arrangement] {
					swapHappened = true
					swapHolder := miniArray[index]
					miniArray[index] = miniArray[index+1]
					miniArray[index+1] = swapHolder
				}
			}
			if !swapHappened {
				solutionArray = append(solutionArray, miniArray)

			}
		}
	}
	return solutionArray
}

func filterOffInvalidArrays(array [][]int, checkArrangementExists map[string]bool) (correctlyOrderedArray, filteredOffArray [][]int) {
	correctlyOrderedArray = [][]int{}
	filteredOffArray = [][]int{}
	for _, miniArray := range array {
		for index := 0; index < len(miniArray)-1; index++ {
			arrangement := fmt.Sprintf("%d|%d", miniArray[index], miniArray[index+1])
			if !checkArrangementExists[arrangement] {
				filteredOffArray = append(filteredOffArray, miniArray)
				break
			}
			if index == len(miniArray)-2 {
				correctlyOrderedArray = append(correctlyOrderedArray, miniArray)
			}

		}
	}
	return correctlyOrderedArray, filteredOffArray
}
func sumMiddleElementsInArray(array [][]int) int {
	sumOfMiddleUpdates := 0
	for _, miniArray := range array {
		middleIndex := len(miniArray) / 2
		sumOfMiddleUpdates += miniArray[middleIndex]
	}
	return sumOfMiddleUpdates
}
func readFilesIntoHashMapOfValidEntries(fileName string) map[string]bool {
	solutionMap := make(map[string]bool)
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	keyForArrangingUpdates := strings.Split(string(file), "\n\n")[0]
	lines := strings.Split(keyForArrangingUpdates, "\n")
	for _, line := range lines {
		solutionMap[line] = true
	}
	return solutionMap
}

func readFilesIntoArraysOfUpdates(fileName string) [][]int {
	solutionArray := make([][]int, 0)
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	actualUpdates := strings.Split(string(file), "\n\n")[1]
	lines := strings.Split(actualUpdates, "\n")
	for _, line := range lines {
		updateNumbers := strings.Split(line, ",")
		innerArray := []int{}
		for _, updateNumber := range updateNumbers {
			integer, err := strconv.Atoi(updateNumber)
			if err != nil {
				panic(err)
			}
			innerArray = append(innerArray, integer)
		}
		solutionArray = append(solutionArray, innerArray)
	}
	return solutionArray
}
