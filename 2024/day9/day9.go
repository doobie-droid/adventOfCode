package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "input.txt" // Replace with your file name
	fmt.Println("Checksum (Part 1):", findCheckSum(fileName))
	fmt.Println("Checksum (Part 2):", findCheckSumPartTwo(fileName))
}

func findCheckSum(fileName string) int {
	fileContent := readFile(fileName)
	decompressedData := decompressData(fileContent)
	rearrangedChecksum := reArrangeChecksum(decompressedData)
	return calculateCheckSum(rearrangedChecksum)
}

func findCheckSumPartTwo(fileName string) int {
	fileContent := readFile(fileName)
	decompressedData := decompressData(fileContent)
	rearrangedChecksum := reArrangeChecksumPartTwo(decompressedData, fileContent)
	return calculateCheckSum(rearrangedChecksum)
}

func readFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func decompressData(fileSystemStorageCode string) []int {
	var fileIdCounter int
	result := []int{}

	for index, runeVal := range fileSystemStorageCode {
		value := int(runeVal) - '0'
		if index%2 == 0 {
			result = append(result, createArray(fileIdCounter, value)...)
			fileIdCounter++
		} else {
			result = append(result, createArray(-1, value)...)
		}
	}
	return result
}

func reArrangeChecksumPartTwo(decompressedData []int, compressedData string) []int {
	cumulativeArray := convertToCumulativeArray(compressedData)
	copyCumulative := append([]int{}, cumulativeArray...)

	for idx := len(cumulativeArray) - 1; idx > 0; idx-- {
		if idx%2 != 0 {
			continue
		}
		fileSize := copyCumulative[idx] - copyCumulative[idx-1]
		for i := 0; i <= idx; i++ {
			if i%2 != 0 {
				continue
			}
			availableSpace := cumulativeArray[(i+1)%len(cumulativeArray)] - cumulativeArray[i]
			if availableSpace >= fileSize {
				for size := 0; size < fileSize; size++ {
					from := copyCumulative[idx] - size
					to := cumulativeArray[i] + size + 1
					decompressedData[to] = decompressedData[from]
					decompressedData[from] = -1
				}
				cumulativeArray[i] += fileSize
				cumulativeArray[idx] -= fileSize
				break
			}
		}
	}
	return decompressedData
}

func convertToCumulativeArray(data string) []int {
	values := strings.Split(data, "")
	result := []int{}

	for _, val := range values {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		result = append(result, num)
	}

	for i := 1; i < len(result); i++ {
		result[i] += result[i-1]
	}

	for i := range result {
		result[i]--
	}

	return result
}

func reArrangeChecksum(data []int) []int {
	end := len(data) - 1
	for start := 0; start < end; start++ {
		if data[start] != -1 {
			continue
		}
		for data[end] == -1 && end > start {
			end--
		}
		data[start], data[end] = data[end], data[start]
	}
	return data
}

func calculateCheckSum(data []int) int {
	sum := 0
	for index, value := range data {
		if value != -1 {
			sum += index * value
		}
	}
	return sum
}

func createArray(value, frequency int) []int {
	array := make([]int, frequency)
	for i := range array {
		array[i] = value
	}
	return array
}
