package main

import "testing"

func TestHash(t *testing.T) {
	result := hash("HASH")
	expectedResult := 52
	if result != expectedResult {
		t.Errorf("wrong result for part 1 %d, expected %d", result, expectedResult)
	}
}

func TestPart1(t *testing.T) {
	expectedResult := 1320
	result := part1("./example.txt")
	if result != expectedResult {
		t.Errorf("wrong result for part 1 %d, expected %d", result, expectedResult)
	}
}