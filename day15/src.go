package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jgaye/adventOfCode2023/tools"
)

// HELPERS

func parseFile(filename string) []string {
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

	return strings.Split(txtlines[0], ",")
}

func hash(input string) int {
	value := 0
	for _, char := range input {
		value += int(char)
		value = value * 17
		value = value % 256
	}
	return value
}

// END HELPERS

func part1(filePath string) int {
	startTime := time.Now()
	result := 0

	instructions := parseFile(filePath)
	fmt.Println("instructions: ", instructions)

	hashes := []int{}
	for _, instruction := range instructions {
		hashes = append(hashes, hash(instruction))
	}

	fmt.Println("hashes: ", hashes)
	result = tools.SumS(hashes)

	fmt.Println("Part 1 results: ", result)
	fmt.Println("Part 1 execution time: ", time.Since(startTime))

	return result
}

func part2(filePath string) int {
	startTime := time.Now()
	result := 0

	fmt.Println("Part 2 result: ", result)
	fmt.Println("Part 2 execution time: ", time.Since(startTime))

	return result
}

func main() {
	filePath := "./input.txt"
	part1(filePath)
	part2(filePath)
}
