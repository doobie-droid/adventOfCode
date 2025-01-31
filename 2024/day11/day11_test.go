package main

import (
	"testing"
)

func Test_day11(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day11Sample_test.txt", 55312},
		{"day11Sample.txt", 185894},
	}
	for _, testCase := range testCases {
		actualOutcome := findNumberOfStones(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 11 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}

func Test_day11PartTwo(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day11Sample_test.txt", 65601038650482},
		{"day11Sample.txt", 221632504974231},
	}
	for _, testCase := range testCases {
		actualOutcome := findNumberOfStonesPartTwo(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 11 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
