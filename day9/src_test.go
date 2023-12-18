package main

import (
	"testing"
)

func TestProcessDiffsSuffix(t *testing.T) {
	result := processDiffsSuffix([]int{0, 3, 6, 9, 12, 15})
	expectedResult := 18

	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}

	result = processDiffsSuffix([]int{10, 13, 16, 21, 30, 45})
	expectedResult = 68

	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}
}

func TestProcessDiffsPrefix(t *testing.T) {
	result := processDiffsPrefix([]int{0, 3, 6, 9, 12, 15})
	expectedResult := -3

	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}

	result = processDiffsPrefix([]int{10, 13, 16, 21, 30, 45})
	expectedResult = 5

	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}
}

func TestPart1(t *testing.T) {
	result := part1("./example.txt")
	expectedResult := 114

	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}
}

func TestPart2(t *testing.T) {
	result := part2("./example.txt")
	expectedResult := 2

	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}
}
