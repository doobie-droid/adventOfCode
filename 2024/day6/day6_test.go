package main

import (
	"testing"
)

func Test_day2(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day6Sample_test.txt", 41},
		{"day6Sample.txt", 4559},
	}
	for _, testCase := range testCases {
		actualOutcome := findDistinctPositions(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 2 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}

func Test_day2PartTwo(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day6Sample_test.txt", 6},
		{"day6Sample.txt", 1604},
	}
	for _, testCase := range testCases {
		actualOutcome := findDistinctPositionForObstruction(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 2 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
