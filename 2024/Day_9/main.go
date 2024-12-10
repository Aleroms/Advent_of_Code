package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
	inputPuzzle := getInputPuzzle("input.txt")
	fmt.Printf("The answer to part one is %d\n",partOne(inputPuzzle))

}

// partOne returns the solution to partOne of this day
// The disk map uses a dense format to represent the layout of 
// files and free space on the disk.
func partOne(diskMap string) int {

	sparseFormat := denseToSparseFormat(diskMap)
	compact := compactFiles(sparseFormat)
	return checksum(compact)
}
// checksum returns the aggregate result of
// multiplying each of these blocks' position with the file ID number it contains. 
// The leftmost block is in position 0. If a block contains free space, skip it instead.
func checksum(compact string) (checksum int) {
	res := strings.SplitN(compact,".",2)
	for fileID, file := range res[0] {

		checksum += fileID * runeToInt(file)
	}
	return checksum
}
// compactFiles move file blocks one at a time from the end of the disk
// to the leftmost free space block (until there are no gaps remaining between file blocks).
//
// For the disk map 12345, the process looks like this:
// 0..111....22222
// 02.111....2222.
// 022111....222..
// 0221112...22...
// 02211122..2....
// 022111222......
func compactFiles(sparseFormat string) string {
	
	period := '.'
	runes := []rune(sparseFormat)
	f, r := 0, len(runes) - 1
	for f < r {
		front, rear := runes[f], runes[r]

		// move to empty spot
		if front != period {
			f++
		} else if rear != period {
			// move rear file to empty spot
			runes[f] = rear
			runes[r] = period
			r--
			f++
		}else {
			// move rear ptr to file spot
			r--
		}

	}
	return string(runes)
}


// denseToSparseFormat converts the diskMap from dense format
// to sparseFormat. It returns the sparseFormat and the total
// number of files.
//
// The digits alternate between indicating the length of a file
// and the length of free space.
//
// ex: diskMap=12345 -> sparseFormat 0..111....22222
func denseToSparseFormat(denseFormat string) string {

	fileID := 0
	sparseFormat := strings.Builder{}
	for i := 1; i < len(denseFormat); i += 2 {
		freeSpaceLength, fileLength := byteToInt(denseFormat[i]), byteToInt(denseFormat[i-1])
		file := generateFile(fileID, fileLength)
		space := generateSpace(freeSpaceLength)

		sparseFormat.WriteString(file)
		sparseFormat.WriteString(space)
		fileID++
	}

	// last file
	fileLength := int(denseFormat[len(denseFormat)-1] - '0')
	file := generateFile(fileID, fileLength)
	sparseFormat.WriteString(file)


	return sparseFormat.String()
}
// runeToInt returns an integer conversion of a rune.
func runeToInt(r rune) int {
	return int(r - '0')
}
// byteToInt returns an integer conversion of a byte
func byteToInt(b byte) int {
	return int(b - '0')
}

// generateFile returns a string given the fileID and length of that file
//
// fileID=0, fileLength=9 would return '000000000'
func generateFile(fileID, fileLength int) string{
	var buffer strings.Builder
	id := strconv.Itoa(fileID)
	for f := 0; f < fileLength; f++ {
		buffer.WriteString(id)
	}
	return buffer.String()
}
// generateSpace returns a string given the length of that designated space
//
// freeSpaceLength=3 would return '...'
func generateSpace(spaceLength int) string {
	var buffer strings.Builder
	for s := 0; s < spaceLength; s++ {
		buffer.WriteString(".")
	}
	return buffer.String()
}

// getInputPuzzle opens the file given and returns a string
// to be used as puzzle input.
func getInputPuzzle(filename string) (s string) {
	// Read the file
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Convert content to a string
	s = string(content)

	return s
}