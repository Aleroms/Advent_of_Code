package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)



func main(){
	manufacturingLab := getMap("input.txt")
	guardStartingInfo, err := getGuardStartingPosition(manufacturingLab)
	if err != nil {
		log.Fatalln(err)
	}
	guard := guard{
		direction: guardStartingInfo.direction,
		current_position: guardStartingInfo.position,
	}
	fmt.Printf("Part One: The answer is %d\n", partOne(manufacturingLab, guard))

}

// guardLeftMappedArea returns true if the guard has left the manufacturing lab.
// bounds check.
func guardLeftMappedArea(guardPos point, bounds point) bool {
	return guardPos[0] < 0 || guardPos[0] >= bounds[0] || guardPos[1] < 0 || guardPos[1] >= bounds[1]
}

// partOne returns the distinct positions that the guard will visit 
// before leaving the mapped area
func partOne(manufacturingLab [][]string, guard guard) int{
	
	obstacles := "#"
	
	// always 1 distinct position
	distinct_positions := 1
	manufacturingLab[guard.current_position[0]][guard.current_position[1]] = "X"

	border_x, border_y := len(manufacturingLab), len(manufacturingLab[0])
	bounds := point{border_x,border_y}

	// while guard still in mapped area
	for !guardLeftMappedArea(guard.current_position, bounds) {
		direction := guard.getDirection()
		newPos := point{
			guard.current_position[0] + direction[0],
			guard.current_position[1] + direction[1],
		}
		
		// if new spot is within bounds
		if !guardLeftMappedArea(newPos,bounds){
			new_position_on_map := manufacturingLab[newPos[0]][newPos[1]]
			// check if new spot has an obstacle
			if new_position_on_map == obstacles {
				guard.turnDirection()
			}else {
				guard.headTowards(newPos)
				// mark with 'X' after checking if not present
				if new_position_on_map != "X" {
					distinct_positions++
				}

				manufacturingLab[newPos[0]][newPos[1]] = "X"
			}
		}else {
			// left mapped area
			break;
		}
	}
	
	return distinct_positions
}
// getGuardStartingPosition takes in a [][]string and checks if
// a guard is located in the manufacturing lab. It returns the
// starting position if found; else an error
func getGuardStartingPosition(manufacturingLab [][]string) (guardStartingInfo,  error) {
	var gsi guardStartingInfo
	gsiDir := map[string]dir {
		"^":north,
		">": east,
		"v": south,
		"<": west,
	}
	for i, row := range manufacturingLab {
		for j, char := range row {
			if strings.Contains("^<>v", char) {
				gsi.direction = gsiDir[char]
				gsi.position = point{i,j}
				return gsi, nil
			}
		}
	}
	return guardStartingInfo{}, errors.New("no guard found in the manufacturing lab")
}
// getMap returns a slice of slice string which is 
// the manufacturing Lab map
func getMap(filename string) (xxs [][]string) {
	fi, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer fi.Close()

	s := bufio.NewScanner(fi)
	for s.Scan() {
		ln := s.Text()
		xs := strings.Split(ln,"")
		xxs = append(xxs, xs)
	}
	return xxs
}