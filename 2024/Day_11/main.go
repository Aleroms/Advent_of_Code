package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var memoization = make(map[string] string)
func main(){
	stones := strings.Fields(getInputPuzzle("input.txt")) 
	blinks := 75
	fmt.Printf("The answer to part one is %d\n",partOne(stones, 25))
	fmt.Printf("The answer to part Two is %d\n",partTwo(stones, blinks))

}
// partOne returns how many stones will you have after blinking 'X' times?
func partOne(stones []string, blinks int) (sum int) {

	
	for i:=0; i < blinks; i++ {
		newStone := getNewStone(stones)
		stones = strings.Fields(newStone)

	}

	
	return len(stones)
}

// partTwo returns how many stones will you have after blinking 'X' times?
// but uses memoization.
func partTwo(stones []string, blinks int) int {

	a := 1
	for i:=0; i < blinks; i++ {
		newStone := getNewStoneMemo(stones)
		stones = strings.Fields(newStone)
		a++
		fmt.Println(a)
	}

	
	return len(stones)
}

// getNewStone is the helper function for part One that receives a slice of strings and returns a string
// follows the three main requirements:
//
// 1 - If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
//
// 2 - If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. 
//		- The left half of the digits are engraved on the new left stone
//		- The right half of the digits are engraved on the new right stone. 
//		- New numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.
//
// 3 - If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
func getNewStone(stones []string) string {

	var stoneBuilder strings.Builder
	for _, stone := range stones {

		stoneLength := len(stone)
		stoneValue, _ := strconv.Atoi(stone)

		if stone == "0" {
			stoneBuilder.WriteString("1 ")
		} else if stoneLength % 2 == 0 {
			pivot := stoneLength / 2

			left, right := stone[:pivot], stone[pivot:]

			// remove trailing zeros
			leftValue, _ := strconv.Atoi(left)
			rightValue, _ := strconv.Atoi(right)

			stoneBuilder.WriteString(strconv.Itoa(leftValue) + " ") 
			stoneBuilder.WriteString(strconv.Itoa(rightValue) + " ") 

		}else {
			newValue := stoneValue * 2024
			stoneBuilder.WriteString(strconv.Itoa(newValue) + " ") 
		}
	}
	return stoneBuilder.String()
}


// getNewStoneMemo is the same as getNewStone but uses memoization
// techniques for larger inputs
func getNewStoneMemo(stones []string) string {

	var stoneBuilder strings.Builder
	for _, stone := range stones {

		// check memoization 🤔
		if memoValue, ok := memoization[stone]; ok {
			stoneBuilder.WriteString(memoValue + " ")
			continue
		}

		stoneLength := len(stone)
		stoneValue, _ := strconv.Atoi(stone)
		var result string

		if stone == "0" {
			result = "1"
		} else if stoneLength % 2 == 0 {

			pivot := stoneLength / 2

			left, right := stone[:pivot], stone[pivot:]

			// remove trailing zeros
			// leftValue, _ := strconv.Atoi(left)
			rightValue, _ := strconv.Atoi(right)

			result = fmt.Sprintf("%v %d", left, rightValue)

		}else {
			newValue := stoneValue * 2024
			result = strconv.Itoa(newValue)			
		}
		memoization[stone] = result
		stoneBuilder.WriteString(result + " ")
	}
	return stoneBuilder.String()
}

// getInputPuzzle opens the file given and returns a string
// to be used as puzzle input.
func getInputPuzzle(filename string) string {

	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	return string(bytes) 
}