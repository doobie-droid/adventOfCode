package main

import (
	"testing"
)

func Test_findAntinodes(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day8Sample_test.txt", 14},
		{"day8Sample.txt", 423},
	}
	for _, testCase := range testCases {
		actualOutcome := countAntinodes(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 8 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}

func Test_findAntinodesPartTwo(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day8Sample_test.txt", 34},
		{"day8Sample.txt", 1287},
	}
	for _, testCase := range testCases {
		actualOutcome := countAntinodesPartTwo(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 8 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
