package main

import (
	"slices"
	"testing"
)

func TestParseGalaxies(t *testing.T) {
	startUniverse := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#....."}

	expectedGalaxies := []Galaxy{
		{0, 3},
		{1, 7},
		{2, 0},
		{4, 6},
		{5, 1},
		{6, 9},
		{8, 7},
		{9, 0},
		{9, 4}}

	actualGalaxies := parseGalaxies(startUniverse)
	if !slices.EqualFunc(actualGalaxies, expectedGalaxies, galaxyEquals) {
		t.Errorf("wrong result %d, expected %d", actualGalaxies, expectedGalaxies)
	}
}

func TestGalaxyDistance(t *testing.T) {
	galaxy1 := Galaxy{6, 1}
	galaxy2 := Galaxy{11, 5}
	expectedResult := 9

	result := galaxy1.Distance(galaxy2)
	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}

	galaxy1 = Galaxy{0, 4}
	galaxy2 = Galaxy{10, 9}
	expectedResult = 15

	result = galaxy1.Distance(galaxy2)
	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}

	galaxy1 = Galaxy{0, 2}
	galaxy2 = Galaxy{7, 12}
	expectedResult = 17

	result = galaxy1.Distance(galaxy2)
	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}

	galaxy1 = Galaxy{11, 0}
	galaxy2 = Galaxy{11, 5}
	expectedResult = 5

	result = galaxy1.Distance(galaxy2)
	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}
}

func TestExpandUniverse(t *testing.T) {
	galaxies := []Galaxy{
		{0, 3},
		{1, 7},
		{2, 0},
		{4, 6},
		{5, 1},
		{6, 9},
		{8, 7},
		{9, 0},
		{9, 4}}
	expandedUniverse := expandUniverse(9, 9, galaxies)
	slices.SortFunc(expandedUniverse, xComparison)
	slices.SortFunc(expandedUniverse, yComparison)
	expectedUniverse := []Galaxy{
		{0, 4},
		{1, 9},
		{2, 0},
		{5, 8},
		{6, 1},
		{7, 12},
		{10, 9},
		{11, 0},
		{11, 5}}

	if !slices.EqualFunc(expandedUniverse, expectedUniverse, galaxyEquals) {
		t.Errorf("wrong result %d, expected %d", expandedUniverse, expectedUniverse)
	}
}

func TestPart1(t *testing.T) {
	result := part1("./example.txt")
	expectedResult := 374

	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}
}