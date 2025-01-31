package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/3

func dayOfAdventCode3(filename string) int {
	outputArray := readTextInputIntoMulStrings(filename)
	return multiplyThenSumArray(outputArray)
}

func dayOfAdventCode3PartTwo(filename string) int {
	outputArray := readAllowedTextFromInputIntoMulStrings(filename)
	return multiplyThenSumArray(outputArray)
}

func readAllowedTextFromInputIntoMulStrings(fileName string) [][2]int {
	var solutionArray [][2]int
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "do()")
	for _, line := range lines {
		doableLines := strings.Split(line, "don't()")
		actualDoableLine := doableLines[0]
		solutionArray = append(solutionArray, getMultiplicationValuesFromString(actualDoableLine)...)
	}
	return solutionArray
}
func readTextInputIntoMulStrings(fileName string) [][2]int {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return getMultiplicationValuesFromString(string(file))
}

func getMultiplicationValuesFromString(str string) [][2]int {
	var solutionArray [][2]int
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)
	lines := re.FindAllStringSubmatch(str, -1)
	for _, line := range lines {
		if len(line) != 3 {
			continue
		}
		digit1, err1 := strconv.Atoi(line[1])
		digit2, err2 := strconv.Atoi(line[2])
		if err1 != nil || err2 != nil {
			fmt.Println("Err found in string to int conversion" + err1.Error())
			continue
		}
		solutionArray = append(solutionArray, [2]int{digit1, digit2})
	}
	return solutionArray
}

func multiplyThenSumArray(input [][2]int) int {
	solutionSum := 0
	for _, value := range input {
		solutionSum += value[0] * value[1]
	}
	return solutionSum
}
