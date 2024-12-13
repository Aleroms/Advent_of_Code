package main

import "testing"

func TestGetTrailheads(t *testing.T) {
	input := []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}
	result := getTrailheads(input)
	expected := []trailhead{
		{point{0 ,2},0}, 
		{point{0 ,4},0} ,
		{point{2 ,4}, 0} ,
		{point{4, 6},0},
		{point{5,2},0},
		{point{5 ,5},0},
		{point {6 ,0},0},
		 {point{6, 6},0}, 
		 {point{7, 1},0},
	}

	for idx, th := range result {
		if th.position.row != result[idx].position.row || th.position.col != result[idx].position.col {
			t.Errorf("Expected %v, Got %v\n",th, expected[idx])
		}
	}
}