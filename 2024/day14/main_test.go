package main

import (
	"testing"
)

func Test_day14(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
		width           int
		height          int
	}{
		{"day14Sample_test.txt", 12, 11, 7},
		{"day14Sample.txt", 226548000, 101, 103},
	}
	for _, testCase := range testCases {
		actualOutcome := findSafetyFactor(testCase.fileName, testCase.width, testCase.height)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 13 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}

func Test_day14PartTwo(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
		width           int
		height          int
	}{
		{"day14Sample_test.txt", 875318608908, 11, 7},
		{"day14Sample.txt", 73267584326867, 101, 103},
	}
	for _, testCase := range testCases {
		actualOutcome := findSafetyFactorPartTwo(testCase.fileName, testCase.width, testCase.height)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 14 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
