package main

import "testing"

func Test_getSumOfMiddleUpdates(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day5Sample_test.txt", 143},
		{"day5Sample.txt", 5639},
	}
	for _, testCase := range testCases {
		actualOutcome := findSumOfMiddleUpdates(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 5 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}

func Test_getSumOfMiddleUpdatesForIncorrectArray(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day5Sample_test.txt", 123},
		{"day5Sample.txt", 5273},
	}
	for _, testCase := range testCases {
		actualOutcome := findSumOfMiddleUpdatesForIncorrectOrder(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 5 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
