package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// HELPERS

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

func countWins(line string) float64 {
	fmt.Print(line)
	var cardWins float64 = 0

	line = strings.Split(line, ": ")[1]
	winningNums := strings.Split(strings.Split(line, " | ")[0], " ")
	// fmt.Println("winningNums: ", winningNums)
	numsIHave := strings.Split(strings.Split(line, " | ")[1], " ")
	// fmt.Println("numsIHave: ", numsIHave)

	for _, num := range numsIHave {
		if num == "" {
			continue
		}
		for _, wNum := range winningNums {
			if num == wNum {
				cardWins += 1
			}
		}
	}

	fmt.Println(" ||| cardWins: ", cardWins)
	return cardWins
}

func sumS(slice []int) int{
	r := 0
	for _, a := range slice{
		r += a
	}
	return r
}

// END HELPERS

func part1(fileLines []string) {
	startTime := time.Now()
	var result float64 = 0

	for _, line := range fileLines {
		cardWins := countWins(line)
		if cardWins >0 {
			result += math.Pow(2, cardWins-1)
		}
	}

	fmt.Println("Part 1 result: ", result)
	fmt.Println("Part 1 execution time: ", time.Since(startTime))
}

func part2(fileLines []string) {
	startTime := time.Now()
	var result int = 0

	lastLine := fileLines[len(fileLines)-1]
	a, _ := strings.CutPrefix(strings.Split(lastLine, ":")[0], "Card ")
	maxCard, _ := strconv.Atoi(a)
	fmt.Println("maxCard: ", maxCard)

	cardEarned := make([]int, maxCard)

	//init cardEarned
	for i, _ := range cardEarned {
		cardEarned[i] = 1
	}

	// check the wins
	for cardIndex, line := range fileLines {
		countWins := int(countWins(line))
		for i := cardIndex+1; i<=cardIndex+countWins; i++ {
			cardEarned[i] += cardEarned[cardIndex]
		}
	}
		
	fmt.Println("cardEarned: ", cardEarned)
	result = sumS(cardEarned)

	fmt.Println("Part 2 result: ", result)
	fmt.Println("Part 2 execution time: ", time.Since(startTime))
}

func main() {
	// fileLines := readFile("./example1.txt")
	fileLines := readFile("./input1.txt")
	part1(fileLines)
	part2(fileLines)
}
