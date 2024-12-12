package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
type point struct {
	row, col int
}
type trailhead struct {
	row, col int
}
func main(){
	topographyMap := getInputPuzzle("test.txt")
	trailheads := getTrailheads(topographyMap)
	fmt.Printf("The answer to part one is %d\n",partOne(trailheads, topographyMap))

}
// partOne returns all scores for trailheads
func partOne(trailheads []trailhead, topographyMap []string) (sum int) {

	for _, trail := range trailheads {
		

	}
	return sum
}
// getAdjacentPoints
func getAdjacentPoints()
// getTrailheads returns all trailhead starting points
// found in the topography map.
func getTrailheads(topographyMap []string) []trailhead {
	var th []trailhead
	for i, v := range topographyMap {
		for j, r := range v {
			if r == 48 {
				th = append(th, trailhead{i, j})
			}
		} 
	}
	return th
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