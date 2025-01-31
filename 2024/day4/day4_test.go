package main

import (
	"testing"
)

func Test_day4PartTwo(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day4Sample_test.txt", 9},
		{"day4Sample.txt", 1945},
	}
	for _, testCase := range testCases {
		actualOutcome := findTwoMASInXShapeInFile(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 4 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
func Test_day4(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day4Sample_test.txt", 18},
		{"day4Sample.txt", 2458},
	}
	for _, testCase := range testCases {
		actualOutcome := findXmasInFile(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 4 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
