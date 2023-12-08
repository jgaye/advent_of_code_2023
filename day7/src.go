package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	hand          string
	formattedHand string
	bet           int
}

// HELPERS

func mapValues(a map[string]int) []int {
	x := []int{}
	for _, val := range a {
		x = append(x, val)
	}

	return x
}

func countCards(hand string) ([]int, int) {
	c := make(map[string]int)
	nbJokers := 0

	for _, card := range hand {
		char := string(card)
		if char == "J" {
			nbJokers += 1
		} else if _, ok := c[char]; ok {
			c[char] += 1
		} else {
			c[char] = 1
		}
	}

	values := mapValues(c)
	slices.Sort(values)
	slices.Reverse(values)
	return values, nbJokers
}

func processJokers(counts []int, nbJokers int) []int {
	if nbJokers == 5 {
		return []int{5}
	}
	
	for i := 0; i < nbJokers; i++ {
		if counts[0] < 5 {
			counts[0] += 1
		} else {
			counts[1] += 1
		}
	}

	return counts
}

func identifyHandType(hand string) string {
	counts, nbJokers := countCards(hand)
	counts = processJokers(counts, nbJokers)

	if counts[0] == 5 { // 1. Five of a kind
		hand = "7" + hand
	} else if counts[0] == 4 { // 2. Four of a kind
		hand = "6" + hand
	} else if counts[0] == 3 && counts[1] == 2 { // 3. Full house
		hand = "5" + hand
	} else if counts[0] == 3 { // 4. Three of a kind
		hand = "4" + hand
	} else if counts[0] == 2 && counts[1] == 2 { // 5. Two pair
		hand = "3" + hand
	} else if counts[0] == 2 { // 6. One pair
		hand = "2" + hand
	} else { // 7. High card
		hand = "1" + hand
	}

	return hand
}

func reformatHand(hand string) string {
	hand = strings.Replace(hand, "J", "0", -1)
	hand = strings.Replace(hand, "1", "a", -1)
	hand = strings.Replace(hand, "2", "b", -1)
	hand = strings.Replace(hand, "3", "c", -1)
	hand = strings.Replace(hand, "4", "d", -1)
	hand = strings.Replace(hand, "5", "e", -1)
	hand = strings.Replace(hand, "6", "f", -1)
	hand = strings.Replace(hand, "7", "g", -1)
	hand = strings.Replace(hand, "8", "h", -1)
	hand = strings.Replace(hand, "9", "i", -1)
	hand = strings.Replace(hand, "T", "j", -1)
	hand = strings.Replace(hand, "Q", "l", -1)
	hand = strings.Replace(hand, "K", "m", -1)
	hand = strings.Replace(hand, "A", "n", -1)
	return hand
}

func sortGames(games []Game) []Game {
	slices.SortFunc[[]Game](games, func(a, b Game) int {
		if a.formattedHand > b.formattedHand {
			return 1
		} else if a.formattedHand < b.formattedHand {
			return -1
		}
		return 0
	})
	return games
}

func parseLine(line string) Game {
	parts := strings.Fields(line)
	bet, _ := strconv.Atoi(parts[1])
	g := Game{parts[0], "", bet}
	return g
}

func readFile(filename string) []Game {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	games := make([]Game, len(txtlines))
	for i, line := range txtlines {
		games[i] = parseLine(line)
	}

	return games
}

// END HELPERS

func part1(filePath string) int {
	// DEPRECATED with the Joker modifications
	startTime := time.Now()

	games := readFile(filePath)
	for i, game := range games {
		games[i].formattedHand = reformatHand(identifyHandType(game.hand))
	}
	games = sortGames(games)

	result := 0
	for i, game := range games {
		result += game.bet * (i + 1)
	}

	fmt.Printf("Part 1 results %d: \n", result)
	fmt.Println("Part 1 execution time: ", time.Since(startTime))

	return result
}

func part2(filePath string) int {
	startTime := time.Now()

	games := readFile(filePath)
	for i, game := range games {
		games[i].formattedHand = reformatHand(identifyHandType(game.hand))
	}
	games = sortGames(games)

	result := 0
	for i, game := range games {
		result += game.bet * (i + 1)
	}

	fmt.Println("Part 2 result: ", result)
	fmt.Println("Part 2 execution time: ", time.Since(startTime))

	return result
}

func main() {
	filePath := "./input1.txt"
	// part1(filePath)
	part2(filePath)
}
