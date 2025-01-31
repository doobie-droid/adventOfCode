package main

import (
	"os"
	"strconv"
	"strings"
)

func findNumberOfStones(fileName string) int {
	fileContent := readFile(fileName)
	stones := readStringIntoIntSlice(fileContent)
	return blink(stones, 25)
}

func blink(stones []int, frequency int) int {
	Count := map[int]int{}
	for _, stone := range stones {
		Count[stone]++
	}
	for i := 0; i < frequency; i++ {
		newStones := map[int]int{}
		for stone, count := range Count {
			if stone == 0 {
				newStones[1] += count
				continue
			}
			stoneDigits := strconv.Itoa(stone)
			if len(stoneDigits)%2 == 0 {
				midPoint := len(stoneDigits) / 2
				leftPart := stoneDigits[:midPoint]
				rightPart := stoneDigits[midPoint:]
				leftNumber, err1 := strconv.Atoi(leftPart)
				rightNumber, err2 := strconv.Atoi(rightPart)
				if err1 != nil {
					panic(err1)
				}
				if err2 != nil {
					panic(err2)
				}
				newStones[leftNumber] += count
				newStones[rightNumber] += count
				continue

			}
			newStones[stone*2024] += count

		}
		Count = newStones
	}
	sum := 0
	for _, count := range Count {
		sum += count
	}
	return sum
}
func findNumberOfStonesPartTwo(fileName string) int {

	fileContent := readFile(fileName)
	stones := readStringIntoIntSlice(fileContent)
	return blink(stones, 75)
}

func readStringIntoIntSlice(str string) []int {
	intSlice := []int{}
	str = strings.TrimSpace(str)
	strArray := strings.Split(str, " ")
	for _, strVal := range strArray {
		integer, err := strconv.Atoi(strVal)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, integer)

	}
	return intSlice
}
func readFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(file)
}
