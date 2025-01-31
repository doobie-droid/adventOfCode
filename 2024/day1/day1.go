package main

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/1

func historianHysteria(filename string) int {
	slice1, slice2 := sortTextInputIntoTwoSortedSlices(filename)
	differenceBetweenSlices := differenceBetweenTwoSlices(slice1, slice2)
	return sumSlice(differenceBetweenSlices)
}

func similarityIndex(fileNAME string) int {
	slice1, slice2 := sortTextInputIntoTwoSortedSlices(fileNAME)
	hashMapOfNumberFrequency := make(map[int]int)
	solutionSum := 0
	for _, value := range slice2 {
		hashMapOfNumberFrequency[value] = hashMapOfNumberFrequency[value] + 1
	}
	for _, value := range slice1 {
		solutionSum += value * hashMapOfNumberFrequency[value]
	}
	return solutionSum
}

func sortTextInputIntoTwoSortedSlices(fileName string) ([]int, []int) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	var slice1, slice2 []int
	for _, line := range lines {
		digits := strings.Fields(line)
		digit1, _ := strconv.Atoi(strings.TrimSpace(digits[0]))
		digit2, _ := strconv.Atoi(strings.TrimSpace(digits[1]))
		slice1 = append(slice1, digit1)
		slice2 = append(slice2, digit2)
	}

	sort.Slice(slice1, func(i, j int) bool {
		return slice1[i] < slice1[j]
	})
	sort.Slice(slice2, func(i, j int) bool {
		return slice2[i] < slice2[j]
	})

	return slice1, slice2
}

func differenceBetweenTwoSlices(slice1 []int, slice2 []int) []int {
	if len(slice1) != len(slice2) {
		panic("The two slices must have the same length")
	}
	solutionSlice := make([]int, len(slice1))
	for index := range slice1 {
		solutionSlice[index] = int(math.Abs(float64(slice1[index]) - float64(slice2[index])))
	}
	return solutionSlice

}
func sumSlice(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}
