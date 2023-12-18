package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// HELPERS

func parseFile(filename string) (Coord, []string) {
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

	startCoord := Coord{0, 0, ""}
	for row, line := range txtlines {
		column := strings.Index(line, "S")
		if column != -1 {
			startCoord = Coord{row, column, ""}
			break
		}
	}

	return startCoord, txtlines
}

type Coord struct {
	row  int
	col  int
	from string
}

var terrain []string
var terrain2 []string
var firstPipe Coord
var secondPipe Coord

func (c Coord) Equals(other Coord) bool {
	return c.row == other.row && c.col == other.col
}

func (c Coord) GoNorth() Coord {
	return Coord{c.row - 1, c.col, "S"}
}

func (c Coord) GoSouth() Coord {
	return Coord{c.row + 1, c.col, "N"}
}

func (c Coord) GoEast() Coord {
	return Coord{c.row, c.col + 1, "W"}
}

func (c Coord) GoWest() Coord {
	return Coord{c.row, c.col - 1, "E"}
}

func (c Coord) Next() Coord {
	symbol := string(terrain[c.row][c.col])

	switch symbol {
	case "|":
		if c.from == "N" {
			return c.GoSouth()
		} else if c.from == "S" {
			return c.GoNorth()
		} else {
			errors.New("invalid from direction")
		}
	case "-":
		if c.from == "E" {
			return c.GoWest()
		} else if c.from == "W" {
			return c.GoEast()
		} else {
			errors.New("invalid from direction")
		}
	case "L":
		if c.from == "N" {
			return c.GoEast()
		} else if c.from == "W" {
			return c.GoNorth()
		} else {
			errors.New("invalid from direction")
		}
	case "F":
		if c.from == "S" {
			return c.GoEast()
		} else if c.from == "E" {
			return c.GoSouth()
		} else {
			errors.New("invalid from direction")
		}
	case "7":
		if c.from == "S" {
			return c.GoWest()
		} else if c.from == "W" {
			return c.GoSouth()
		} else {
			errors.New("invalid from direction")
		}
	case "J":
		if c.from == "N" {
			return c.GoWest()
		} else if c.from == "W" {
			return c.GoNorth()
		} else {
			errors.New("invalid from direction")
		}
	default:
		errors.New("invalid from direction")
	}

	return c.GoNorth()
}

func nextPipes(firstPipe Coord, secondPipe Coord) (Coord, Coord) {
	firstPipe = firstPipe.Next()
	secondPipe = secondPipe.Next()

	return firstPipe, secondPipe
}

func replaceAtIndex(str string, replacement rune, index int) string {
    runes := []rune(str)
    runes[index] = replacement
    return string(runes)
}


func ioMap(terrain []string) []string {
	for row, line := range terrain {
		for 

	terrain[0] = replaceAtIndex(terrain[0], 'O', 0)
	
}

// func isInsidePath(coord Coord, path []Coord) bool {
// 	crossing := 0

// 	return coord.row >= 0 && coord.row < len(terrain) && coord.col >= 0 && coord.col < len(terrain[coord.row])
// }

// END HELPERS

// func part1(filePath string) int {
// 	startTime := time.Now()
// 	result := 0
// 	var startCoord Coord

// 	startCoord, terrain = parseFile(filePath)
// 	fmt.Println("startCoord, terrain: ", startCoord, terrain)

// 	firstPipe = Coord{startCoord.row, startCoord.col - 1, "E"}
// 	secondPipe = Coord{startCoord.row, startCoord.col + 1, "W"}
// 	distance := 1

// 	for !firstPipe.Equals(secondPipe) {
// 		firstPipe, secondPipe = nextPipes(firstPipe, secondPipe)
// 		distance += 1
// 	}

// 	result = distance

// 	fmt.Println("Part 1 results: ", result)
// 	fmt.Println("Part 1 execution time: ", time.Since(startTime))

// 	return result
// }

func part2(filePath string) int {
	startTime := time.Now()
	result := 0
	var startCoord Coord

	startCoord, terrain = parseFile(filePath)
	terrain2 := make([]string, len(terrain))
	copy(terrain2, terrain)
	terrain2[startCoord.row][startCoord.col] = "X"

	fmt.Println("startCoord, terrain: ", startCoord, terrain)

	firstPipe = Coord{startCoord.row, startCoord.col - 1, "E"}
	secondPipe = Coord{startCoord.row, startCoord.col + 1, "W"}
	distance := 1

	for !firstPipe.Equals(secondPipe) {
		a := firstPipe
		b := secondPipe
		terrain2[a.row] = "X"
		terrain2[b.row] = "X"
		firstPipe, secondPipe = nextPipes(firstPipe, secondPipe)
		distance += 1
	}

	result = distance

	fmt.Println("Part 2 result: ", result)
	fmt.Println("Part 2 execution time: ", time.Since(startTime))

	return result
}

func main() {
	filePath := "./input.txt"
	// part1(filePath)
	part2(filePath)
}
