package main

import (
	"testing"
)

func Test_day3PartTwo(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day3Sample_test.txt", 48},
		{"day3Sample.txt", 92082041},
	}
	for _, testCase := range testCases {
		actualOutcome := dayOfAdventCode3PartTwo(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 3 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
func Test_day3(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day3Sample_test.txt", 161},
		{"day3Sample.txt", 191183308},
	}
	for _, testCase := range testCases {
		actualOutcome := dayOfAdventCode3(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 3 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
