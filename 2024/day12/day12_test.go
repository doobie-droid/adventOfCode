package main

import (
	"testing"
)

func Test_day12(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day12Sample_test.txt", 140},
		{"day12Sample_test1.txt", 772},
		{"day12Sample_test2.txt", 1930},
		{"day12Sample.txt", 1465968},
	}
	for _, testCase := range testCases {
		actualOutcome := getPriceOfFence(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 12 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
func Test_PartTwo_day12(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day12Sample_test.txt", 140},
		{"day12Sample_test1.txt", 772},
		{"day12Sample_test2.txt", 1930},
	}
	for _, testCase := range testCases {
		actualOutcome := getPriceOfFencePartTwo(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 12 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
