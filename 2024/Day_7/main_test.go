package main

import (
	"testing"
)

func TestParseEquation(t *testing.T) {
	expectedKeys := []int{123,456,789,1011}
	expectedValues := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
		{1,0,11},
	}
	inputs := []string{
		"123: 1 2 3",
		"456: 4 5 6",
		"789: 7 8 9",
		"1011: 1 0 11",
	}
	for i, input := range inputs {
		resultKey, resultValues := parseEquation(input)
		if resultKey != expectedKeys[i]{
			t.Error("error in parsing the equation")
		}
		for i2, value := range resultValues {
			if value != expectedValues[i][i2] {
				t.Error("error in parsing the equation")
			}
		}
	}
}
func TestPartOneHelper(t *testing.T){

	if !partOneHelper(190, []int{10, 19}){
		t.Error("Expected true but got false")
	}
}