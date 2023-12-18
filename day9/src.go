package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jgaye/adventOfCode2023/tools"
)

// HELPERS

func parseFile(filename string) [][]int {
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

	startSlices := [][]int{}
	for _, eachline := range txtlines {
		startSlices = append(startSlices, tools.SAtoi(eachline))
	}
	return startSlices
}

func getDiffs(input []int) []int {
	diffs := []int{}
	for i := 0; i < len(input)-1; i++ {
		diffs = append(diffs, input[i+1]-input[i])
	}
	return diffs
}

func processDiffsSuffix(input []int) int {
	sum := tools.SumS(input)
	if sum == 0 && input[0] == input[len(input)-1] {
		return 0
	} else {
		diffs := getDiffs(input)
		xnplusone := processDiffsSuffix(diffs)
		return xnplusone + input[len(input)-1]
	}
}

func processDiffsPrefix(input []int) int {
	sum := tools.SumS(input)
	if sum == 0 && input[0] == input[len(input)-1] {
		return 0
	} else {
		diffs := getDiffs(input)
		xnplusone := processDiffsPrefix(diffs)
		return input[0] - xnplusone
	}
}

// END HELPERS

func part1(filePath string) int {
	startTime := time.Now()
	result := 0

	startSlices := parseFile(filePath)
	fmt.Println("startSlices: ", startSlices)

	acc := []int{}

	for _, eachSlice := range startSlices {
		acc = append(acc, processDiffsSuffix(eachSlice))
	}
	fmt.Println("acc: ", acc)

	result = tools.SumS(acc)

	fmt.Println("Part 1 results: ", result)
	fmt.Println("Part 1 execution time: ", time.Since(startTime))

	return result
}

func part2(filePath string) int {
	startTime := time.Now()
	result := 0

	startSlices := parseFile(filePath)
	fmt.Println("startSlices: ", startSlices)

	acc := []int{}

	for _, eachSlice := range startSlices {
		acc = append(acc, processDiffsPrefix(eachSlice))
	}
	fmt.Println("acc: ", acc)

	result = tools.SumS(acc)


	fmt.Println("Part 2 result: ", result)
	fmt.Println("Part 2 execution time: ", time.Since(startTime))

	return result
}

func main() {
	filePath := "./input.txt"
	part1(filePath)
	part2(filePath)
}
