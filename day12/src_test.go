package main

import (
	"slices"
	"testing"
)

func TestPermuteSentence(t *testing.T) {
	parameters := []struct {
		input    string
		expected []string
	}{
		{"?", []string{".", "#"}},
		{".", []string{"."}},
		{"#", []string{"#"}},
		{"?.", []string{"..", "#."}},
		{"???.", []string{"....", "..#.", ".#..", ".##.", "#...", "#.#.", "##..", "###."}},
	}

	for i := range parameters {

		sentence := parameters[i].input
		expected := parameters[i].expected
		actual := permuteSentence(sentence)

		// Sort to compare
		slices.Sort(expected)
		slices.Sort(actual)

		if !slices.Equal(expected, actual) {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	}
}

func TestIsSentenceValid(t *testing.T) {
	parameters := []struct {
		inputSentence string
		inputClues    []int
		expected      bool
	}{
		{"#", []int{1}, true},
		{"#", []int{0}, false},
		{".", []int{1}, false},
		{".", []int{0}, true},
		{"##", []int{1}, false},
		{"##", []int{2}, true},
		{"#.#", []int{1, 1}, true},
		{"###", []int{2}, false},
		{"###", []int{3}, true},
		{".###.##.#...", []int{3, 2, 1}, true},
		{".###.##..#..", []int{3, 2, 1}, true},
		{".###.##...#.", []int{3, 2, 1}, true},
		{".###.##....#", []int{3, 2, 1}, true},
		{".###..##.#..", []int{3, 2, 1}, true},
		{".###..##..#.", []int{3, 2, 1}, true},
		{".###..##...#", []int{3, 2, 1}, true},
		{".###...##.#.", []int{3, 2, 1}, true},
		{".###...##..#", []int{3, 2, 1}, true},
		{".###....##.#", []int{3, 2, 1}, true},
		{".###....##.#", []int{3, 3, 1}, false}}

	for i := range parameters {

		sentence := parameters[i].inputSentence
		clues := parameters[i].inputClues
		expected := parameters[i].expected
		actual := isSentenceValid(sentence, clues)

		if expected != actual {
			t.Errorf("For parameters %v Expected %v, got %v", parameters[i], expected, actual)
		}
	}

}

func TestBruteForceSolve(t *testing.T) {
	lines := []line{
		{"???.###", []int{1, 1, 3}},
		{".??..??...?##.", []int{1, 1, 3}},
		{"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}},
		{"????.#...#...", []int{4, 1, 1}},
		{"????.######..#####.", []int{1, 6, 5}},
		{"?###????????", []int{3, 2, 1}}}

	actual := bruteForceSolve(lines)
	expected := 21
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestExpandLine(t *testing.T) {
	parameters := []struct {
		line     line
		expected line
	}{
		{
			line{".#", []int{1}}, 
			line{".#?.#?.#?.#?.#", []int{1, 1, 1, 1, 1}},
		},
		{
			line{"???.###", []int{1, 1, 3}}, 
			line{"???.###????.###????.###????.###????.###", []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3}},
		},
	}

	for i := range parameters {

		line := parameters[i].line
		expected := parameters[i].expected
		actual := expandLine(line)

		if !expected.Equals(actual){
			t.Errorf("For parameters %v Expected %v, got %v", parameters[i], expected, actual)
		}
	}
}
