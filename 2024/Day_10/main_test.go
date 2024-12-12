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
		{0 ,2}, 
		{0 ,4} ,
		{2 ,4} ,
		{4, 6} ,
		{5,2} ,
		{5 ,5},
		 {6 ,0} ,
		 {6, 6}, 
		 {7, 1},
	}
	for idx, th := range result {
		if th.row != result[idx].row || th.col != result[idx].col {
			t.Errorf("Expected %v, Got %v\n",th, expected[idx])
		}
	}
}