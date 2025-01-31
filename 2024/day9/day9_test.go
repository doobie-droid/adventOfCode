package main

import (
	"testing"
)

func Test_day9(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day9Sample_test.txt", 1928},
		{"day9Sample.txt", 6330095022244},
	}
	for _, testCase := range testCases {
		actualOutcome := findCheckSum(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 9 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}

func Test_day9PartTwo(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day9Sample_test.txt", 2858},
		{"day9Sample.txt", 6359491814941},
	}
	for _, testCase := range testCases {
		actualOutcome := findCheckSumPartTwo(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 9 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
