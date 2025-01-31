package main

import (
	"testing"
)

func Test_day2PartTwo(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day2Sample_test.txt", 4},
		{"day2Sample.txt", 398},
	}
	for _, testCase := range testCases {
		actualOutcome := dayOfAdventCode2PartTwo(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 2 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
func Test_day2(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day2Sample_test.txt", 2},
		{"day2Sample.txt", 332},
	}
	for _, testCase := range testCases {
		actualOutcome := dayOfAdventCode2(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 2 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
