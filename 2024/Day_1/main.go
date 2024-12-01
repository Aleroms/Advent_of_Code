package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)
func main(){

	left, right := getInputFile("input.txt")
	
	fmt.Printf("The answer for Day 1 Part One is %d\n", partOne(left,right))
	fmt.Printf("The answer for Day 1 Part Two is %d\n", partTwo(left,right))

}

// partOne returns the answer to Advent of Code Part 1
func partOne(left, right []int) (total_distance int) {
	
	slices.Sort(left)
	slices.Sort(right)

	for i := 0; i < len(left); i++ {
		total_distance += absDiff(left[i], right[i])
	}
	return total_distance
}

// partTwo returns the answer to Advent of Code Part 1
func partTwo(left, right []int) (similarity_score int) {

	m := make(map[int]int, len(right))
	for i := 0; i < len(right); i++ {
		m[right[i]]++
	}

	for i := 0; i < len(left); i++ {
		curr := left[i]
		similarity_score += curr * m[curr]
	}


	return similarity_score
}
// absDiff calculates the absolute difference between two integers.
func absDiff(a, b int) int {
	if a > b {
        return a - b
    }
    return b - a
}

// getInputFile parses Day 1's input file for the 
// Advent of Code challenge. getInputFile returns two
// slices - left and right
//
// Input:
//
// 64433 75582
//
// 87936 20843
//
// 98310 69076
// 
// Returns:
//
// l := []slice {64433 87936 98310}
//
// r := []slice {75582 20843 69076}
func getInputFile(filePath string) (l []int, r []int) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening %v\n", f.Name())
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	for s.Scan() {
		ln := strings.Fields(s.Text())
		n1, n2 := convertToInt(ln)
		l = append(l, n1)
		r = append(r, n2)
	}
	return l, r
}

// convertToInt is a helper function that converts the slice of strings
// into two integers, returning back to the user, two integers
func convertToInt(ln []string) (int,int) {
	n1, err := strconv.Atoi(ln[0])
	if err != nil {
		fmt.Printf("Failed to convert '%s' to integer: %v", ln[0], err)
	}
	n2, err := strconv.Atoi(ln[1])
	if err != nil {
		fmt.Printf("Failed to convert '%s' to integer: %v", ln[1], err)
	}
	return n1,n2 
}
