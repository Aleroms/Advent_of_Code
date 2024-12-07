package main

import (
	"testing"
)

func TestGuardGetDirection(t *testing.T) {
	guard := guard{
		direction:        east,
		current_position: [2]int{0, 0},
	}
	for i := 1; i <= 16; i++ {
		direction := guard.getDirection()
		guard.headTowards(direction)

		if i == 4 {
			guard.direction = south
		}else if i == 8 {
			guard.direction = west
		} else if i == 12 {
			guard.direction = north
		}
	}
	if guard.current_position != [2]int{0,0} {
		t.Errorf("guard direction did not work properly, expected current_position=(%d,%d) got=(%d,%d)\n",
	0,0,guard.current_position[0],guard.current_position[1])
	}
}
func TestGenerateMap(t *testing.T){
	test_map := [][]string{
		{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "#", ".", ".", "^", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
		{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
	}
	result := getMap("test.txt")
	for i := 0; i < len(test_map); i++ {
		for j := 0; j < len(test_map[i]); j++ {
			if result[i][j] != test_map[i][j] {
				t.Error("Mismatch generated with test.txt and expected test_map")
			}
		}
	}
	
}
func TestGetGuardStartingPosition(t *testing.T) {
	test_map := [][]string{
		{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "#", ".", ".", "^", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
		{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
	}
	gsi, _ := getGuardStartingPosition(test_map)
	if gsi.direction != north || gsi.position != [2]int{6,4} {
		t.Errorf("incorrect starting position, expected (6,4)")
	}
}
func TestPartOne(t *testing.T){
	test_map := [][]string{
		{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "#", ".", ".", "^", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
		{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
	}
	guardStartingInfo, _ := getGuardStartingPosition(test_map)
	
	guard := guard{
		direction: guardStartingInfo.direction,
		current_position: guardStartingInfo.position,
	}
	result := partOne(test_map, guard)
	if result != 41 {
		t.Errorf("Expected 41 got %d\n", result)
	}
}
func TestGuardLeftMappedArea(t *testing.T) {
	test_map := [][]string{
		{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "#", ".", ".", "^", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
		{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
	}
	border_x, border_y := len(test_map), len(test_map[0])
	bounds := point{border_x, border_y}

	// Points within bounds
	points := []point{
		{0, 0},
		{5, 3},
		{9, 0},
		{0, 9},
		{9, 9},
	}
	
	for _, p := range points {
		if guardLeftMappedArea(p, bounds) {
			t.Errorf("point %v incorrectly detected as out of bounds", p)
		}
	}

	// Points outside bounds
	points = []point{
		{-1, 0},
		{5, 10},
		{11, 0},
		{-3, 3},
	}
	for _, p := range points {
		if !guardLeftMappedArea(p, bounds) {
			t.Errorf("point %v incorrectly detected as within bounds", p)
		}
	}
}
