package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main(){
	corruptedMemory := getCorruptedMemory("input.txt")
	fmt.Printf("The answer for Day 3 Part One is %d\n", partOne(corruptedMemory))
	fmt.Printf("The answer for Day 3 Part Two is %d\n", partTwo(corruptedMemory))
}

func partOne(corruptedMemory []string) (res int) {

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for _, v := range corruptedMemory {
		matches := re.FindAllStringSubmatch(v, -1)
		for _, match := range matches {
			res += calculate(match)
		}
	}
	return res
}
func partTwo(corruptedMemory []string) (res int) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don\'t\(\)|do\(\)`)
	for _, v := range corruptedMemory {
		matches := re.FindAllStringSubmatch(v,-1)
		skip := false

		for _, match := range matches {

			if match[0] == "do()" {
				skip = false
			} else if match[0] == "don't()"{
				skip = true
			}

			if !skip {
				res += calculate(match)
			}

		}
	}
	return res
}

func calculate(xs []string) int {
	n1, _ := strconv.Atoi(xs[1])
	n2, _ := strconv.Atoi(xs[2])
	return n1 * n2
}

func getCorruptedMemory(filename string) []string {
	fi, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer fi.Close()

	var xs []string
	s := bufio.NewScanner(fi)
	for s.Scan() {
		memory := s.Text()
		xs = append(xs, memory)
	}
	return xs
}
