package main

import (
	"testing"
)


func TestPart1Example1(t *testing.T) {
	result := part1("./example1.txt")
	expectedResult := 2

	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}
}

func TestPart1Example2(t *testing.T) {
	result := part1("./example2.txt")
	expectedResult := 6

	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}
}

func TestPart2Example3(t *testing.T) {
	result := part2("./example3.txt")
	expectedResult := 6

	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}
}
