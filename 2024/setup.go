package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	year, day := time.Now().Year(), time.Now().Day()

	if day > 25 {
		fmt.Println("Advent of Code is now over! 🎄🦌")
	} else {
		fmt.Printf("Fetching Advent of Code - Day %d 🎅\n", day)

		// Fetch and save the input file
		inputFilePath := getInputFile(year, day)
		saveInputFile(day, inputFilePath)
		createMainGO(day)

		fmt.Printf("Fetch complete! Happy Hacking! ✨🎄\n")
	}
}

// createMainGO creates main.go in the current day's directory.
// main.go uses boiler plate code to fill main.go
//
// package main func main(){ }
func createMainGO(day int){
	destDir := fmt.Sprintf("Day_%d", day)
	destFile := fmt.Sprintf("%s/main.go", destDir)

	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create directory %s: %v\n", destDir, err)
	}

	content := "package main\n\nfunc main(){}"
	err = os.WriteFile(destFile, []byte(content), 0644)
	if err != nil {
		log.Fatalf("Failed to create main.go file in %s: %v\n", destDir, err)
	}
}

// saveInputFile says the input.txt into the current day's directory
//
// IE Day_6/input.txt
func saveInputFile(day int, inputFilePath string) {
	destDir := fmt.Sprintf("Day_%d/", day)
	destFile := destDir + "input.txt" // Destination filename is hardcoded

	// Create destination directory
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create directory %s: %v\n", destDir, err)
	}

	// Open the source file for reading
	src, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Failed to open source file %s: %v\n", inputFilePath, err)
	}
	defer src.Close()

	// Create the destination file
	dst, err := os.Create(destFile)
	if err != nil {
		log.Fatalf("Failed to create destination file %s: %v\n", destFile, err)
	}
	defer dst.Close()

	// Copy contents from source file to destination file
	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatalf("Failed to copy file contents: %v\n", err)
	}

	log.Printf("File successfully saved to %s\n", destFile)
}

// getInputFile submits a GET request to adventofcode.com and requests
// the current day's puzzle input.
func getInputFile(year, day int) string {
	envMap := getEnvMap()
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: envMap["SESSION"],
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	// Write the response body to a temporary file
	tempFilePath := "input.txt"
	err = os.WriteFile(tempFilePath, body, 0644)
	if err != nil {
		log.Fatalf("Error writing to file %s: %v", tempFilePath, err)
	}

	return tempFilePath
}

// getEnvMap returns a map data structure of all environment variables
// saved in /2024/.env
//
// Specifically, this methods returns the session key used in the GET request
// submitted to Advent of Code for authentication
func getEnvMap() map[string]string {
	env := make(map[string]string)
	fi, err := os.Open(".env")
	if err != nil {
		log.Fatalln(err)
	}
	defer fi.Close()

	s := bufio.NewScanner(fi)
	for s.Scan() {
		variable := strings.SplitN(s.Text(), "=", 2) // Use SplitN to handle "=" in values
		if len(variable) == 2 {
			k, v := strings.TrimSpace(variable[0]), strings.TrimSpace(variable[1])
			env[k] = v
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}
	return env
}
