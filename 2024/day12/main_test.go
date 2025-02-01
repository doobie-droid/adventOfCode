package main

import (
	"testing"
)

func Test_day12(t *testing.T) {
	testCases := []struct {
		input           string
		expectedOutcome int
	}{
		{"AAAA\nBBCD\nBBCC\nEEEC", 140},
		{"OOOOO\nOXOXO\nOOOOO\nOXOXO\nOOOOO", 772},
		{"RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE", 1930},
		{"EEEEE\nEXXXX\nEEEEE\nEXXXX\nEEEEE", 692},
		{"AAAAAA\nAAABBA\nAAABBA\nABBAAA\nABBAAA\nAAAAAA", 1184},
		// {"day12Sample.txt", 1465968},
	}
	for _, testCase := range testCases {
		actualOutcome := getPriceOfFence(testCase.input)
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
		{"AAAA\nBBCD\nBBCC\nEEEC", 80},
		{"OOOOO\nOXOXO\nOOOOO\nOXOXO\nOOOOO", 436},
		{"RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE", 1206},
		{"EEEEE\nEXXXX\nEEEEE\nEXXXX\nEEEEE", 236},
		{"AAAAAA\nAAABBA\nAAABBA\nABBAAA\nABBAAA\nAAAAAA", 368},
		// {"day12Sample.txt", 897702},
	}
	for _, testCase := range testCases {
		actualOutcome := getPriceOfFencePartTwo(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in day 12 of advent of Code, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
