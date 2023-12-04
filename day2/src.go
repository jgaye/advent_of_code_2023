package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// puzzle1()
    puzzle2()
}

func puzzle1() {
	var games = parseInput("./day2/input1.txt")
	fmt.Println(games)

	var acc = 0

	for _, game := range games {
		acc += nbFromCondition(game)
	}

	fmt.Println(acc)
}

func puzzle2() {
	var games = parseInput("./day2/input1.txt")
	fmt.Println(games)

    acc := 0

    // for each color, when does it become possible?
    // at the max of each draw
	for _, game := range games {
		gameNb, draws := parseGame(game)
        maxBlue, maxRed, maxGreen := maxFromMaps(draws)
        fmt.Println("gameNb :", gameNb)
        fmt.Println("maxBlue, maxRed, maxGreen :", maxBlue, maxRed, maxGreen)
        acc += maxBlue * maxRed * maxGreen
	}

    fmt.Println("solution :", acc)
	
}

func maxFromMaps(maps []map[string]int) (int, int, int) {
    maxBlue, maxRed, maxGreen := 0, 0, 0
    for _, m := range maps {
        if m["blue"] > maxBlue {
            maxBlue = m["blue"]
        }
        if m["red"] > maxRed {
            maxRed = m["red"]
        }
        if m["green"] > maxGreen {
            maxGreen = m["green"]
        }
    }
    return maxBlue, maxRed, maxGreen
}

func parseInput(filename string) []string {
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

	return txtlines
}

func nbFromCondition(game string) int {
	fmt.Println(game)

	gameNb, draws := parseGame(game)

	for _, draw := range draws {
		if !isPossible(draw) {
			fmt.Printf("game %d impossible", gameNb)
			gameNb = 0
		}
	}

	return gameNb
}

func parseGame(game string) (int, []map[string]int) {
	parts := strings.Split(game, ": ")
	gameInfo, draws := parts[0], parts[1]

	// Game nb
	gameNbStr, _ := strings.CutPrefix(gameInfo, "Game ")
	fmt.Println("gameNb :" + gameNbStr)
	gameNb, _ := strconv.Atoi(gameNbStr)

	// Draws
	// red, blue, green results per draw
	// default value for a draw
	drawsParsed := []map[string]int{}
	for _, draw := range strings.Split(draws, "; ") {
		drawParsed := defaultDraw()
		for _, color := range strings.Split(draw, ", ") {
			red, found := strings.CutSuffix(color, " red")
			if found {
				drawParsed["red"], _ = strconv.Atoi(red)
			}

			blue, found := strings.CutSuffix(color, " blue")
			if found {
				drawParsed["blue"], _ = strconv.Atoi(blue)
			}

			green, found := strings.CutSuffix(color, " green")
			if found {
				drawParsed["green"], _ = strconv.Atoi(green)
			}
		}
		fmt.Println("drawParsed :", fmt.Sprintf("%v", drawParsed))
		drawsParsed = append(drawsParsed, drawParsed)
	}
	fmt.Println("drawsParsed :", fmt.Sprintf("%v", drawsParsed))

	return gameNb, drawsParsed
}

func defaultDraw() map[string]int {
	var defaultDraw = make(map[string]int)
	defaultDraw["red"] = 0
	defaultDraw["blue"] = 0
	defaultDraw["green"] = 0
	return defaultDraw
}

func isPossible(draw map[string]int) bool {
	conditionMap := make(map[string]int)
	conditionMap["red"] = 12
	conditionMap["blue"] = 14
	conditionMap["green"] = 13

	if draw["red"] > conditionMap["red"] {
		return false
	}
	if draw["blue"] > conditionMap["blue"] {
		return false
	}
	if draw["green"] > conditionMap["green"] {
		return false
	}
	return true
}
