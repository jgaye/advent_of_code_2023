package main

import (
	"testing"
)

func TestIOMap(t *testing.T) {
	startMap := []string{
		"..XX..",
		".XXX..",
		".X.XX.",
		".X..X.",
		".X..X.",
		".XXXX."}
	expectedResult := []string{
		"OOXXOO",
		"OXXXOO",
		"OXIXXO",
		"OXIIXO",
		"OXIIXO",
		"OXXXXO"}

	
	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}
}

// func TestPart1(t *testing.T) {
// 	result := part1("./example.txt")
// 	expectedResult := 9

// 	if result != expectedResult {
// 		t.Errorf("wrong result %d, expected %d", result, expectedResult)
// 	}
// }