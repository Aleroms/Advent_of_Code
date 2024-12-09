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
	inputSlice := getInputPuzzle("test.txt")
	fmt.Printf("The answer to part one is %d\n",partOne(inputSlice))
	fmt.Printf("The answer to part two is %d\n",partTwo(inputSlice))
}

// partOne returns the solution to partOne of this day
func partOne(calibrations []string) (total int) {

	for _, equation := range calibrations {
		solution, testValues := parseEquation(equation)
		if partOneHelper(solution, testValues){
			total += solution
		}
	}
	return total
}
// partOneHelper recursively calls itself to create subproblems and determines
// if the equation can be solved using '+' or 'x' operators. Returns false if 
// cannot
func partOneHelper(target int, testValues []int) bool {

	if len(testValues) == 1 {
		return target == testValues[0]
	}
	// calculate
	mult := testValues[0] * testValues[1]
	sum := testValues[0] + testValues[1]

	// make slice
	newValues := testValues[2:]
	newValues = append([]int{mult}, newValues...)

	// recurse
	if partOneHelper(target, newValues){
		return true
	}
	newValues[0] = sum
	return partOneHelper(target, newValues)
}

// partTwo returns the total calibration result for operators
// '+', '*', and 'concat'
func partTwo(calibrations []string) (total int) {

	for _, equation := range calibrations {
		solution, testValues := parseEquation(equation)
		if partTwoHelper(solution, testValues){
			total += solution
		}
	}
	return total
}

// partOneHelper recursively calls itself to create subproblems and determines
// if the equation can be solved using '+', 'x', or 'concat' operators. 
// Returns false if cannot.
func partTwoHelper(target int, testValues []int) bool {

	if len(testValues) == 1 {
		return target == testValues[0]
	}
	// calculate
	mult := testValues[0] * testValues[1]
	sum := testValues[0] + testValues[1]
	concat, err := strconv.Atoi(fmt.Sprintf("%d%d",testValues[0], testValues[1]))
	fmt.Printf("%d || %d: %d\n",testValues[0], testValues[1], concat)
	if err != nil {
		log.Fatal(err)
	}

	// make slice
	newValues := testValues[2:]
	newValues = append([]int{mult}, newValues...)

	// recurse
	if partOneHelper(target, newValues){
		return true
	}
	newValues[0] = sum
	if partTwoHelper(target, newValues){
		return true
	}
	newValues[0] = concat
	return partTwoHelper(target, newValues)
}

// parseEquation parses the equation by deliminer ':' and returns two
// values: (target int, testValues []int)
func parseEquation(equation string) ( int,  []int) {
	xs := strings.Split(equation,":")
	
	target, err := strconv.Atoi(xs[0])
	if err != nil {
		log.Fatalln(err)
	}

	xs2 := strings.Fields(xs[1])
	
	testValues := make([]int, len(xs2))
	for i, v := range xs2 {
		tv, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		testValues[i] = tv
	}
	return target, testValues
}
// getInputPuzzle opens the file given and returns a slice of string
// to be used as puzzle input.
func getInputPuzzle(filename string) (xs []string) {
	fi, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer fi.Close()

	s := bufio.NewScanner(fi)
	for s.Scan() {
		ln := s.Text()
		xs = append(xs, ln)
	}
	return xs
}