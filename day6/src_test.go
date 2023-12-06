package main

import (
	"testing"
)

func equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func TestParseLine(t *testing.T) {
	timeStr := "Time: 0 1  2 3 4 5      6 7 8 9 10                                          11 12"
	expectedResult := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	result := parseLine(timeStr)
	if !equal(result, expectedResult) {
		t.Errorf("Did not provide the expected result: %d instead of %d", result, expectedResult)
	}
}

func TestMergeIntoOneInt(t *testing.T) {
	expectedResult := 71530
	result := mergeIntoOneInt([]int{7, 15, 30})
	if result != expectedResult {
		t.Errorf("Did not provide the expected result: %d instead of %d", result, expectedResult)
	}
}

func TestPart1(t *testing.T) {
	expectedResult := 288
	result := part1("./example1.txt")
	if result != expectedResult {
		t.Errorf("wrong result for part 1 %d, expected %d", result, expectedResult)
	}
}


func TestPart2(t *testing.T) {
	expectedResult := 71503
	result := part2("./example1.txt")
	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}
}