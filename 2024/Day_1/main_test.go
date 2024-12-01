package main

import "testing"

func TestPartOne(t *testing.T) {
	left := []int{3,4,2,1,3,3}
	right := []int{4,3,5,3,9,3}
	
	result := partOne(left,right)
	expect := 11

	if result != expect {
		t.Errorf("Expected %d Got %d", expect, result)
	}


}

func TestPartTwo(t *testing.T) {
	left := []int{3,4,2,1,3,3}
	right := []int{4,3,5,3,9,3}

	result := partTwo(left,right)
	expect := 31

	if result != expect {
		t.Errorf("Expected %d Got %d", expect, result)
	}
}

func TestAbsDiff(t *testing.T) {
	a, b := 10, 12

	result := absDiff(a,b)
	expect := 2

	if result != expect {
		t.Errorf("Expected %d Got %d", expect, result)
	}
}

func TestGetInputFile(t *testing.T) {
	left := []int{3,4,2,1,3,3}
	right := []int{4,3,5,3,9,3}

	l,r := getInputFile("test.txt")

	for i, v := range left {
		if l[i] != v {
			t.Errorf("Left slice did not match result from getInputFile\nExpected:%d, got %d",l[i], v)
		}
	}

	for i, v := range right {
		if r[i] != v {
			t.Errorf("Right slice did not match result from getInputFile\nExpected:%#v, got %#v",r[i], v)
		}
	}
}

func TestConvertToInt(t *testing.T) {
	p := []string {"1","2"}

	r1, r2 := convertToInt(p)
	if r1 != 1 {
		t.Errorf("Did not properly convert string to int.")
	}
	if r2 != 2 {
		t.Errorf("Did not properly convert string to int.")
	}
}