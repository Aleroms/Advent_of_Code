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
	position point
	value int
}
func main(){
	topographyMap := getInputPuzzle("test.txt")
	trailheads := getTrailheads(topographyMap)
	fmt.Printf("The answer to part one is %d\n",partOne(trailheads, topographyMap))

}
// partOne returns all scores for trailheads
func partOne(trailheads []trailhead, topographyMap []string) (sum int) {

	for _, trail := range trailheads {

		// get neighbors using current trail
		neighbors := getNeighbors(trail, topographyMap)
		fmt.Printf("valid neighbors for trail %v are %v\n", trail, neighbors)
		
		var visited = make(map[trailhead]bool)
		visited[trail] = true
		
		q := Queue{}
		q.Enqueue(trail)

		// visit all neighbors
		for !q.IsEmpty() {
			visit, err := q.Dequeue()
			if err != nil {
				log.Fatal(err)
			}

			pt, ok := visit.(trailhead)
			if !ok {
				log.Fatalf("unexpected type %T, expected point", visit)
			}
			
			fmt.Printf(" - %v\n", pt)
			visited[pt] = true

			// check if reached goal
			if pt.value == 9 {
				sum++
			}

			newNeighbors := getNeighbors(pt, topographyMap)
			fmt.Printf("- %v\n", newNeighbors)
			for _, nn := range newNeighbors {
				if _, ok := visited[nn]; !ok {
					q.Enqueue(nn)
				}
			}
		}

	}
	return sum
}

// getNeighbors returns all neighbors of a trail only if that 
// neightbor is value + 1 of current value. Checks horizontal
// and vertical neighbors only.
func getNeighbors(trail trailhead, topographyMap []string) []trailhead{
	var neighbors []trailhead
	row, col := trail.position.row, trail.position.col
	
	// directions
	n,s,e,w := row-1, row+1, col+1, col-1
	north := point{n, col}
	south := point{s, col}
	east := point{row, e}
	west := point{row, w}

	// bounds check and valid neighbor check
	if inBounds(north, topographyMap) && isValidNeighbor(trail.value, topographyMap[n][col]){
		value := byteToInt(topographyMap[n][col])
		neighbors = append(neighbors, trailhead{north,value})
	}
	if inBounds(south, topographyMap) && isValidNeighbor(trail.value, topographyMap[s][col]){
		value := byteToInt(topographyMap[s][col])
		neighbors = append(neighbors, trailhead{south, value})
	}
	if inBounds(east, topographyMap) && isValidNeighbor(trail.value, topographyMap[row][e]){
		value := byteToInt(topographyMap[row][e])
		neighbors = append(neighbors, trailhead{east, value})
	}
	if inBounds(west, topographyMap) && isValidNeighbor(trail.value, topographyMap[row][w]){
		value := byteToInt(topographyMap[row][w])
		neighbors = append(neighbors, trailhead{west,value })
	}
	
	return neighbors

}
// inBounds checks if the point is inbounds given topographyMap
func inBounds(point point, topographyMap []string) bool{
	return point.row >= 0 && point.row < len(topographyMap) && point.col >= 0 && point.col < len(topographyMap[point.row])
}
func isValidNeighbor(currentSpot int, value byte)bool{
	return currentSpot + 1 == byteToInt(value)
}
// byteToInt returns an integer conversion of a byte
func byteToInt(b byte) int {
	return int(b - '0')
}
// getAdjacentPoints
// func getAdjacentPoints(){}

// getTrailheads returns all trailhead starting points
// found in the topography map.
func getTrailheads(topographyMap []string) []trailhead {
	var th []trailhead
	for i, v := range topographyMap {
		for j, r := range v {
			// ASCII code for '0'
			if r == 48 {
				th = append(th, trailhead{ point{i,j}, 0})
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