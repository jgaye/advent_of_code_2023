package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	// "strconv"
	"strings"
)

func readFile(filename string) []string {
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

func inputToGrid(entry []string) [][]string {
	grid := make([][]string, len(entry))
	for i, line := range entry {
		grid[i] = strings.Split(line, "")
	}
	return grid
}

type Coord struct {
	row      int
	colStart int
	colEnd   int
}

//	func (c Coord) distance(other Coord) int {
//		return math.Sqrt2(c.row-other.row) + min(min(abs(c.colStart-other.colStart), abs(c.colStart-other.colEnd)), abs(c.colStart-((other.colStart+other.colEnd)/2)))
//	}
func (c Coord) distance(other Coord) float64 {
	rowDiff := float64(c.row - other.row)
	colStartDiff := float64(c.colStart - other.colStart)
	colEndDiff := float64(c.colStart - other.colEnd)
	avgColDiff := float64(c.colStart - ((other.colStart + other.colEnd) / 2))

	return math.Sqrt(math.Pow(rowDiff, 2) + math.Min(math.Min(math.Pow(colStartDiff, 2), math.Pow(colEndDiff, 2)), math.Pow(avgColDiff, 2)))
}

func abs(i int) int {
	return int(math.Abs(float64(i)))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func puzzle1() {
	// searchable lines and searchable grid forms
	fileLines := readFile("./input1.txt")
	maxRow := len(fileLines) - 1
	maxCol := len(fileLines[0]) - 1
	grid := inputToGrid(fileLines)
	fmt.Println(grid)

	// find all numbers indices in the file
	re := regexp.MustCompile("[0-9]+")

	indices := []Coord{}
	for row, line := range fileLines {
		foundIndices := re.FindAllStringIndex(line, -1)
		for _, foundIndice := range foundIndices {
			indices = append(indices, Coord{row, foundIndice[0], foundIndice[1]})
		}
	}

	fmt.Println(indices)

	acc := 0
	for _, indice := range indices {
		num, _ := strconv.Atoi(fileLines[indice.row][indice.colStart:indice.colEnd])
		arounds := ""

		arounds += fileLines[max(0, indice.row-1)][max(0, indice.colStart-1):min(maxCol, indice.colEnd+1)]
		arounds += fileLines[indice.row][max(0, indice.colStart-1):min(maxCol, indice.colEnd+1)]
		arounds += fileLines[min(maxRow, indice.row+1)][max(0, indice.colStart-1):min(maxCol, indice.colEnd+1)]
		fmt.Println(num)
		fmt.Println(arounds)

		re := regexp.MustCompile("[^.[0-9]]*")
		if re.FindStringIndex(arounds) != nil {
			acc += num
		}
	}

	fmt.Println(acc)
}

func main() {
	// searchable lines and searchable grid forms
	fileLines := readFile("./input1.txt")
	// fmt.Println(fileLines)

	// find all num indices in the file
	re := regexp.MustCompile("[0-9]+")
	numIndices := []Coord{}
	for row, line := range fileLines {
		foundIndices := re.FindAllStringIndex(line, -1)
		for _, foundIndice := range foundIndices {
			numIndices = append(numIndices, Coord{row, foundIndice[0], foundIndice[1]-1})
		}
	}
	fmt.Println(numIndices)

	// find all gears indices in the file
	re2 := regexp.MustCompile("[*]")
	gearIndices := []Coord{}
	for row, line := range fileLines {
		foundIndices := re2.FindAllStringIndex(line, -1)
		for _, foundIndice := range foundIndices {
			gearIndices = append(gearIndices, Coord{row, foundIndice[0], foundIndice[1]-1})
		}
	}
	fmt.Println(len(gearIndices))

	realGearIndices := []Coord{}
	acc := 0
	for _, gear := range gearIndices {
		arounds := []int{}
		for _, num := range numIndices {
			res, _ := strconv.Atoi(fileLines[num.row][num.colStart:num.colEnd+1])
			if gear.distance(num) < 2 {
				arounds = append(arounds, res)
			}
		}
		if len(arounds) == 2 {
			fmt.Printf("gear at %d,%d is surrounded by %d and %d", gear.row, gear.colStart, arounds[0], arounds[1])
			realGearIndices = append(realGearIndices, gear)
			acc += arounds[0] * arounds[1]
		}
	}
	fmt.Println(realGearIndices)

	fmt.Println(acc)
}
