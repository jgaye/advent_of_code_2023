package main

import "testing"

func TestPart1(t *testing.T) {
	expectedResult := 811
	result := part1("./example.txt")
	if result != expectedResult {
		t.Errorf("wrong result for part 1 %d, expected %d", result, expectedResult)
	}
}