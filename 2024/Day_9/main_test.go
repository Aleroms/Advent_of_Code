package main

import "testing"

func TestGenerateFile(t *testing.T) {
	expect := "000000000"
	result := generateFile(0, 9)
	if result != expect {
		t.Errorf("Expected %v but got %v\n",expect, result)
	}
}

func TestGenerateSpace(t *testing.T) {
	expect := "..."
	result := generateSpace(3)
	if result != expect {
		t.Errorf("Expected %v but got %v\n",expect, result)
	}
}

func TestDenseToSparseFormat(t *testing.T) {
	input := "12345"
	expect := "0..111....22222"
	result := denseToSparseFormat(input)
	if result != expect {
		t.Errorf("Expected %v but got %v\n",expect, result)
	}
}

func TestCompactFiles(t *testing.T) {

	input := "0..111....22222"
	expect := "022111222......"
	result := compactFiles(input)
	if result != expect {
		t.Errorf("Expected %v but got %v\n",expect, result)
	}
}

func TestChecksum(t *testing.T) {
	input := "022111222......"
	expect := 60
	result := checksum(input)
	if result != expect {
		t.Errorf("Expected %v but got %v\n",expect, result)
	}
}