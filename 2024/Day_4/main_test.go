package main

import "testing"

func TestGetInputSlice(t *testing.T) {
	expected := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	
	result := getInputSlice("test.txt")
	for i := 0; i < len(expected); i++ {
		if expected[i] != result[i] {
			t.Errorf("mismatch: expected %v but got %v", expected[i], result[i])
		}
	}

}

func TestHorizontal(t *testing.T) {
	inputSlice := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	if horizontal(inputSlice) != 5 {
		t.Errorf("incorrect number of XMAS or SAMX horizontally")
	}
}
func TestVertical(t *testing.T) {
	inputSlice := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	if vertical(inputSlice) != 3 {
		t.Errorf("incorrect number of XMAS or SAMX horizontally")
	}
}

func TestDiagonal(t *testing.T) {
	inputSlice := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	result := diagonal(inputSlice)
	if result != 10 {
		t.Errorf("incorrect number of XMAS or SAMX horizontally Expected 10 got %v",result)
	}
}