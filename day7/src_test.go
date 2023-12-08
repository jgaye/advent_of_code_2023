package main

import (
	"slices"
	"testing"
)

func TestIdentifyHandType(t *testing.T) {
	// PART 1 test, DEPRECATED
	// handsAndResults := map[string]string{"12345": "112345",
	// 	"12315": "212315",
	// 	"12321": "312321",
	// 	"12223": "412223",
	// 	"12221": "512221",
	// 	"22122": "622122",
	// 	"22222": "722222"}

	handsAndResults := map[string]string{"32T3K": "232T3K",
		"T55J5": "6T55J5",
		"KK677": "3KK677",
		"KTJJT": "6KTJJT",
		"QQQJA": "6QQQJA"}
	for k, v := range handsAndResults {
		result := identifyHandType(k)
		if result != v {
			t.Errorf("for:%s result:%s expected:%s", k, result, v)
		}
	}
	
}

func TestReformatHand(t *testing.T) {
	// PART 1 test, DEPRECATED
	// handsAndResults := map[string]string{
	// 	"123456789TJQKA": "abcdefghijklmn",
	// 	"232T3K":         "bcbjcm",
	// 	"4T55J5":         "djeeke",
	// 	"3KK677":         "cmmfgg",
	// 	"5KTJJT":         "emjkkj",
	// 	"6QQQJA":         "flllkn"}

	handsAndResults := map[string]string{
		"123456789TJQKA": "abcdefghij0lmn",
		"232T3K":         "bcbjcm",
		"4T55J5":         "djee0e",
		"3KK677":         "cmmfgg",
		"5KTJJT":         "emj00j",
		"6QQQJA":         "flll0n"}
	for k, v := range handsAndResults {
		result := reformatHand(k)
		if result != v {
			t.Errorf("for:%s result:%s expected:%s", k, result, v)
		}
	}
}

func TestSortGames(t *testing.T) {
	games := []Game{{"32T3K", "bcbjcm", 1},
		{"T55J5", "djeeke", 4},
		{"KK677", "cmmfgg", 3},
		{"KTJJT", "emjkkj", 2},
		{"QQQJA", "flllkn", 5}}
	expectedResult := []Game{{"32T3K", "bcbjcm", 1},
		{"KTJJT", "emjkkj", 2},
		{"KK677", "cmmfgg", 3},
		{"T55J5", "djeeke", 4},
		{"QQQJA", "flllkn", 5}}
	result := sortGames(games)
	if slices.Equal(result, expectedResult) {
		t.Errorf("for:%v result:%v expected:%v", games, result, expectedResult)
	}
}

// func TestPart1(t *testing.T) {
// 	result := part1("./example1.txt")
// 	expectedResult := 6440

// 	if result != expectedResult {
// 		t.Errorf("wrong result %d, expected %d", result, expectedResult)
// 	}
// }

func TestPart2(t *testing.T) {
	expectedResult := 5905
	result := part2("./example1.txt")

	if result != expectedResult {
		t.Errorf("wrong result %d, expected %d", result, expectedResult)
	}
}
