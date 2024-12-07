package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {

	inputSlice := getInputSlice("test.txt")
	fmt.Printf("The answer for Day 4 Part One is %d\n", partOne(inputSlice))
	// fmt.Printf("The answer for Day 4 Part Two is %d\n", partTwo(inputSlice))
}

func partOne(inputSlice []string) (xmas int){ 
	
	xmas += horizontal(inputSlice)
	xmas += vertical(inputSlice)
	xmas += diagonal(inputSlice)
	
	
	return xmas
}
// diagonal returns the count of 'XMAS' and 'SAMX' in the word
// search for each column
func diagonal(inputSlice []string) (total int) {
	xmas, samx := "XMAS", "SAMX"

	// Helper function to count matches in a diagonal
	countMatches := func(startRow, startCol, rowStep, colStep int, word string) (sum int) {
		row, col := startRow, startCol
		word_inc := 0

		for row >= 0 && row < len(inputSlice) && col >= 0 && col < len(inputSlice[0]) {
			// Check character match
			if inputSlice[row][col] == word[word_inc] {
				word_inc++
				if word_inc == 4 { // Match found
					sum++
					word_inc = 0
				}
			} else { // Reset on mismatch
				word_inc = 0
			}

			// Move diagonally
			row += rowStep
			col += colStep
		}
		return sum
	}

	// Top-left to bottom-right (↘)
	for row := 0; row < len(inputSlice); row++ {
		total += countMatches(row, 0, 1, 1, xmas)
		total += countMatches(row, 0, 1, 1, samx)
	}
	for col := 1; col < len(inputSlice[0]); col++ {
		total += countMatches(0, col, 1, 1, xmas)
		total += countMatches(0, col, 1, 1, samx)
	}

	// Top-right to bottom-left (↙)
	for row := 0; row < len(inputSlice); row++ {
		total += countMatches(row, len(inputSlice[0])-1, 1, -1, xmas)
		total += countMatches(row, len(inputSlice[0])-1, 1, -1, samx)
	}
	for col := len(inputSlice[0]) - 2; col >= 0; col-- {
		total += countMatches(0, col, 1, -1, xmas)
		total += countMatches(0, col, 1, -1, samx)
	}

	return total
}


// vertical returns the count of 'XMAS' and 'SAMX' in the word
// search for each column
func vertical(inputSlice []string) (total int) {
	xmas, samx := "XMAS", "SAMX"
	xs, sx := 0,0

	// Loop vertically (column by column)
	for col := 0; col < len(inputSlice[0]); col++ { 
		for row := 0; row < len(inputSlice); row++ { 
			// check if XMAS or SAMX in column

			//XMAS
			if inputSlice[row][col] == xmas[xs] {
				xs++
			} else{
				xs = 0
			}

			//SAMX
			if inputSlice[row][col] == samx[sx] {
				sx++
			} else{
				sx = 0
			}


			// match
			if xs == 4 {
				total++
				xs = 0
			}
			if sx == 4 {
				total++
				sx = 0
			}
		}
		xs = 0
		xs = 0
	}
	return total
}
// horizontal returns the count of 'XMAS' in the word search for
// each row.
func horizontal(inputSlice []string) (total int) {
	//use regex to search for 'XMAS' and 'SAMX'
	re_xmas := regexp.MustCompile(`XMAS`)
	re_samx := regexp.MustCompile(`SAMX`)

	for _, search := range inputSlice {
		xmas_match := re_xmas.FindAllStringSubmatch(search,-1)
		samx_match := re_samx.FindAllStringSubmatch(search, -1)

		total += len(xmas_match) + len(samx_match)
	}

	return total
}

func getInputSlice(filename string)(inputSlice []string){
	fi, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer fi.Close()

	s := bufio.NewScanner(fi)
	for s.Scan(){
		ln := s.Text()
		inputSlice = append(inputSlice, ln)
	}
	return inputSlice	
}