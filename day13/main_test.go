package main

import (
	"testing"
)

func Test_day13(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day13Sample_test.txt", 480},
		// {"day13Sample.txt", 6330095022244},
	}
	for _, testCase := range testCases {
		actualOutcome := findMinimumNumberOfTokens(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 13 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}

func Test_day13PartTwo(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		// {"day13Sample_test.txt", 2858},
		// {"day13Sample.txt", 6359491814941},
	}
	for _, testCase := range testCases {
		actualOutcome := findMinimumNumberOfTokensPartTwo(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 13 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
