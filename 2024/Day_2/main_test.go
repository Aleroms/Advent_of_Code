package main

import (
	"fmt"
	"testing"
)


func TestGetInputSlice(t *testing.T) {
	expected := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	result := getInputSlice("test.txt")
	for row, x := range expected {
		res := result[row]
		for idx, v := range x {
			if res[idx] != v {
				t.Errorf("Result does not match expected\nGot: %d, Expected: %d",res[idx], v)
			}
		}
	}

}

func TestIsLevelIncreasingValid(t *testing.T) {
	// test increasing
	xi1 := []int {7, 8, 9}
	if !isLevelIncreasingValid(true, xi1) {
		t.Errorf("Expected increasing levels but got an error\n")
	}
	// test decreasing
	xi2 := []int {9,8,7,6,5,4,3,2,1}
	if !isLevelIncreasingValid(false, xi2) {
		t.Errorf("Expected decreasing levels but got an error\n")
	}

	// test expect fail
	xi3 := []int {9,8,7,8,9}
	if isLevelIncreasingValid(false, xi3){
		t.Error("Expected levels to fail")
	}
	

}

func TestIsLevelIncreasingPartTwoValid(t *testing.T){
	// test increasing
	xi := []int {1,3,2,4,5}
	result := isLevelIncreasingPartTwoValid(true, xi)

	if result == false {
		t.Error("Removing 1 level marks report as safe.")
	}
	
	// test expect false for increasing
	xi2 := []int {1,3,2,4,3}
	result = isLevelIncreasingPartTwoValid(true, xi2)
	if result == true {
		t.Error("Cannot remove more than one level")
	}

	// test decreasing
	xi3 := []int{8,6,4,4,1}
	result = isLevelIncreasingPartTwoValid(false, xi3)
	if result == false {
		t.Error("Removing 1 level marks report as safe.")
	}

	xi4 := []int{100,50,80,60,20,40,50,5}
	result = isLevelIncreasingPartTwoValid(false, xi4)
	if result == true {
		t.Error("Cannot remove more than one level")
	}
}

func TestIsLevelDifferValid(t *testing.T) {
	// test pass
	xi1 := []int{7,5,4,3,1}
	if !isLevelDifferValid(xi1) {
		t.Error("Expected valid level\n")
	}

	// test expect fail
	xi2 := []int{8, 6, 4, 4, 1}
	if isLevelDifferValid(xi2) {
		t.Error("Expected invalid level\n")
	}
}

func TestPartOne(t *testing.T) {
	xxi := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	result := partOne(xxi)
	expect := 2
	if result != expect {
		t.Errorf("Error got %d Expected %d\n", result, expect)
	}
}

func TestPartTwo(t *testing.T){
	xxi := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	result := partTwo(xxi)
	expect := 4
	if result != expect {
		t.Errorf("Error got %d Expected %d\n", result, expect)
	}
}


func TestIsReportIncreasing(t *testing.T) {
	increasing := isReportIncreasing(1,2)
	if !increasing {
		t.Error("Expected values to return true for is increasing 1,2")
	}
	decreasing := isReportIncreasing(2,1)
	if decreasing {
		t.Error("Expected values to return true for is decreasing 2,1")
	}
}

func TestIsReportPartTwoValid(t *testing.T) {
	xi := []int{8, 6, 4, 4, 1}
	result := isReportPartTwoValid(false, xi)
	fmt.Println(result)
	
}