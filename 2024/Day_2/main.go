package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
	inputSlice := getInputSlice("input.txt")
	fmt.Printf("The answer for Day 2 Part One is %d\n", partOne(inputSlice))
	fmt.Printf("The answer for Day 2 Part Two is %d\n", partTwo(inputSlice))
}

 
// partTwo returns the answer to Advent of Code Day 2 - Part Two.
func partTwo(inputSlice [][]int) (total_safe int) {

	for _, report := range inputSlice {

		increasing := isReportIncreasing(report[0], report[1])
		if isReportValid(increasing, report){
			total_safe++
		} else if isReportPartTwoValid(increasing, report){
			// part two - check if removing a level would mark 
			// report as safe.
			total_safe++
		} else {
			fmt.Printf("unsafe %v\n",report)
		}
	}
	return total_safe
}
// partOne returns the answer to Advent of Code Day 2 - Part One.
func partOne(inputSlice [][]int) (total_safe int){

	// check edge case if []int len < 2

	for _, report := range inputSlice {

		increasing := isReportIncreasing(report[0], report[1])
		if isReportValid(increasing, report){
			total_safe++
		}
	}
	return total_safe
}

// isReportIncreasing takes two numbers, a & b, and returns if a is increasing
// compared to b.
//
// return a < b
func isReportIncreasing(a, b int) bool {
	return a < b
}

// isLevelValid checks a slice of int and determines if all levels are in
// all decreasing or all increasing.
func isLevelIncreasingValid(increasing bool, report []int) bool {
	for i := 1; i < len(report); i++ {
		if increasing {
			if report[i-1] > report[i] {
				return false
			}
		} else {
			if report[i-1] < report[i] {
				return false
			}
		}
	}
	return true
}

// isLevelIncreasingPartTwoValid determines if removing a single level from the report would 
// mark the report as safe.
func isLevelIncreasingPartTwoValid(increasing bool, report []int) bool {

	var rm_lvls int
	for i := 1; i < len(report); i++ {
		if increasing && !isReportIncreasing(report[i-1], report[i]) {
			rm_lvls++
		}else if !increasing && isReportIncreasing(report[i-1],report[i]){
			rm_lvls++
		}

		if rm_lvls > 1 {
			// must only be one single level
			return false
		}
	
	}
	return true
}

// isLevelDifferValid checks a slice of int and determines if all levels are valid.
// A level is valid if:
//
// Any two adjacent levels differ by at least one and at most three.
func isLevelDifferValid(report []int) bool {
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff < 0 {
			diff = -diff
		}
		
		if diff < 1 || diff > 3 {
			return false
		}

	}
	return true
}

// isLevelDifferValid checks a slice of int and determines if all levels are valid with a one bad level tolerance.
// A level is valid if:
//
// Any two adjacent levels differ by at least one and at most three.
func isLevelDifferPartTwoValid(report []int) bool {
	var rm_lvls int
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff < 0 {
			diff = -diff
		}
		
		if diff < 1 || diff > 3 {
			rm_lvls++

		
			if rm_lvls > 1 {
				// must only be one single level
				return false
			}
		}

	}
	return true
}

// isReportValid validates if the report is valid given the levels.
// 
// A report is valid if:
//
// 1. The levels are either all increasing or all decreasing.
//
// 2. Any two adjacent levels differ by at least one and at most three.
func isReportValid(increasing bool, report []int) bool {
	return isLevelIncreasingValid(increasing, report) && isLevelDifferValid(report)
}

func isReportPartTwoValid(increasing bool, report []int)bool {
	// return isLevelDifferPartTwoValid(report) && isLevelIncreasingPartTwoValid(increasing, report)
	return isLevelIncreasingValid(increasing, report)
}

// getInputSlice receives a string filename and returns a
// slice of slice ints to be used in part one and two of
// advent of code.
func getInputSlice(filename string)(inputSlice [][]int){
	fi, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer fi.Close()

	s := bufio.NewScanner(fi)
	for s.Scan(){
		report := s.Text()
		xs := strings.Fields(report)
		xi := make([]int, len(xs))

		for i, s := range xs {
			xi[i], err = strconv.Atoi(s)
			if err != nil {
				log.Fatalln(err)
			}
		}
		inputSlice = append(inputSlice, xi)
	}
	return inputSlice	
}