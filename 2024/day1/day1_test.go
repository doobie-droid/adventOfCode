package main

import (
	"reflect"
	"testing"
)

func Test_historianHysteria(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day1Sample_test.txt", 11},
		{"day1Sample.txt", 1938424},
	}
	for _, testCase := range testCases {
		actualOutcome := historianHysteria(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in historian Hysteria, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}

func Test_similarityIndex(t *testing.T) {
	testCases := []struct {
		fileName        string
		expectedOutcome int
	}{
		{"day1Sample_test.txt", 31},
		{"day1Sample.txt", 22014209},
	}
	for _, testCase := range testCases {
		actualOutcome := similarityIndex(testCase.fileName)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Error in Similarity Index, expected %v, Got %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
func Test_sortTextInputIntoTwoSortedSlices(t *testing.T) {
	testCases := []struct {
		fileName             string
		expectedOutputSlice1 []int
		expectedOutputSlice2 []int
	}{
		{"day1Sample_test.txt", []int{1, 2, 3, 3, 3, 4}, []int{3, 3, 3, 4, 5, 9}},
	}

	for _, testCase := range testCases {
		actualOutputSlice1, actualOutputSlice2 := sortTextInputIntoTwoSortedSlices(testCase.fileName)
		if !reflect.DeepEqual(actualOutputSlice1, testCase.expectedOutputSlice1) {
			t.Errorf("The sorted slices are not equal: Expected %v, Actual %v", testCase.expectedOutputSlice1, actualOutputSlice1)
		}
		if !reflect.DeepEqual(actualOutputSlice2, testCase.expectedOutputSlice2) {
			t.Errorf("The sorted slices are not equal: Expected %v, Actual %v", testCase.expectedOutputSlice2, actualOutputSlice2)
		}
	}
}

func Test_differenceBetweenTwoSlices(t *testing.T) {
	testCases := []struct {
		slice1          []int
		slice2          []int
		expectedOutcome []int
	}{
		{[]int{1, 2, 3, 3, 3, 4}, []int{3, 3, 3, 4, 5, 9}, []int{2, 1, 0, 1, 2, 5}},
		{[]int{0, 1, 1, 2, 3, 4}, []int{2, 3, 3, 3, 3, 3}, []int{2, 2, 2, 1, 0, 1}},
	}

	for _, testCase := range testCases {
		actualOutcome := differenceBetweenTwoSlices(testCase.slice1, testCase.slice2)
		if !reflect.DeepEqual(actualOutcome, testCase.expectedOutcome) {
			t.Errorf("Difference between two slices failing, Expected %v Actual %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}

func Test_sumSlice(t *testing.T) {
	testCases := []struct {
		slice           []int
		expectedOutcome int
	}{
		{[]int{1, 2, 3, 3, 3, 4}, 16},
		{[]int{0, 1, 1, 2, 3, 4}, 11},
		{[]int{0, 0, 1}, 1},
	}

	for _, testCase := range testCases {
		actualOutcome := sumSlice(testCase.slice)
		if actualOutcome != testCase.expectedOutcome {
			t.Errorf("Sum of Slice failing, Expected %v Actual %v", testCase.expectedOutcome, actualOutcome)
		}
	}
}
