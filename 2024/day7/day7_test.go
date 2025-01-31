package main

import "testing"

func Test_sumOfBridgeRepairs(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day7Sample_test.txt", 3749},
		{"day7Sample.txt", 28730327770375},
	}
	for _, testCase := range testCases {
		actualOutcome := sumOfBridgeRepairs(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 7 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}

func Test_sumOfBridgeRepairsPartTwo(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day7Sample_test.txt", 11387},
		{"day7Sample.txt", 424977609625985},
	}
	for _, testCase := range testCases {
		actualOutcome := sumOfBridgeRepairsWithConcatenation(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 7 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
