package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/2

const INCREASING = "increasing"
const DECREASING = "decreasing"

func dayOfAdventCode2PartTwo(fileName string) int {
	outputSlice := readTextInputIntoIntSlices(fileName)
	return FindSafeLevelsWithDampener(outputSlice)
}
func dayOfAdventCode2(filename string) int {
	outputSlice := readTextInputIntoIntSlices(filename)
	return FindSafeLevels(outputSlice)
}

func FindSafeLevelsWithDampener(input [][]int) int {
	var counterOfSafeReports int
	for _, slice := range input {
		unsafeLevel := findUnsafeLevel(slice)
		if unsafeLevel == -1 {
			counterOfSafeReports++
			continue
		}
		for index := range slice {
			newSlice := removeElement(slice, index)
			newUnsafeLevel := findUnsafeLevel(newSlice)
			if newUnsafeLevel == -1 {
				counterOfSafeReports++
				break
			}
		}

	}
	return counterOfSafeReports
}

func FindSafeLevels(input [][]int) int {
	var counterOfSafeReports int

	for _, slice := range input {
		if findUnsafeLevel(slice) == -1 {
			counterOfSafeReports++
		}

	}
	return counterOfSafeReports
}

// Returns index of first unsafe spike if there is one
// Returns index of -1 if level is safe
func findUnsafeLevel(slice []int) int {
	var levelStatus string
	if slice[1]-slice[0] >= 0 {
		levelStatus = INCREASING
	} else {
		levelStatus = DECREASING
	}
	for index := range slice {
		if index < 1 {
			continue
		}
		if levelStatus == INCREASING && slice[index]-slice[index-1] < 0 {
			return index
		}
		if levelStatus == DECREASING && slice[index]-slice[index-1] >= 0 {
			return index
		}
		incrementLevel := math.Abs(float64(slice[index] - slice[index-1]))
		if incrementLevel > 3 || incrementLevel < 1 {
			return index
		}
	}
	return -1
}

func readTextInputIntoIntSlices(fileName string) [][]int {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	var outputSlice [][]int
	for _, line := range lines {
		digits := strings.Fields(line)
		innerSlice := []int{}
		for _, digit := range digits {
			integer, _ := strconv.Atoi(digit)
			innerSlice = append(innerSlice, integer)
		}
		outputSlice = append(outputSlice, innerSlice)
	}
	return outputSlice
}

func removeElement(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
