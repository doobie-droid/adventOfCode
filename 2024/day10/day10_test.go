package main

import (
	"testing"
)

func Test_day2(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day10Sample_test.txt", 36},
		{"day10Sample.txt", 754},
	}
	for _, testCase := range testCases {
		actualOutcome := findSumOfTrailHeads(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 2 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}

func Test_day10PartTwo(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day10Sample_test.txt", 81},
		{"day10Sample.txt", 1609},
	}
	for _, testCase := range testCases {
		actualOutcome := findSumOfTrailHeadsPartTwo(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 10 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
