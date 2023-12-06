package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}

// HELPERS

func parseLine(line string) []int {
	parts := strings.Fields(line)[1:]
	els := make([]int, len(parts))
	for i, part := range parts {
		els[i], _ = strconv.Atoi(part)
	}
	return els
}

func readFile(filename string) ([]int, []int) {
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

	return parseLine(txtlines[0]), parseLine(txtlines[1])
}

func calcWinPresses(time int, distance int) []int {
	winPresses := []int{}
	for i := 0; i <= time; i++ {
		if (time-i)*i > distance{
			winPresses = append(winPresses, i)
		}
	}
	return winPresses
}

func optimCalcWinPresses(time int, distance int) int {
	
	firstWin := 0
	lastWin := 0

	for i := 0; i <= time; i++ {
		if firstWin == 0 && i*(time-i) > distance{
			firstWin = i
			lastWin = time-i
			break
		}
		
	}
	return lastWin - firstWin
}

func mergeIntoOneInt(els []int) int {
	mergedStr := ""
	for _, el := range els {
		mergedStr += strconv.Itoa(el)
	}
	merged, _ := strconv.Atoi(mergedStr)
	return merged
}

// END HELPERS

func part1(filePath string) int {
	startTime := time.Now()
	var result int = 1

	times, distances := readFile(filePath)
	fmt.Println("check: ", times, distances)

	for i, time := range times {
		winPresses := calcWinPresses(time, distances[i])

		result = result * len(winPresses)
	}


	fmt.Println("Part 1 result: ", result)
	fmt.Println("Part 1 execution time: ", time.Since(startTime))

	return result
}

func part2(filePath string) int {
	startTime := time.Now()
	var result int = 1

	times, distances := readFile(filePath)
	time1 := mergeIntoOneInt(times)
	distance := mergeIntoOneInt(distances)
	fmt.Println("check: ", time1, distance)

	winPresses := calcWinPresses(time1, distance)
	result = result * len(winPresses)
	
	fmt.Println("Part 2 result: ", result)
	fmt.Println("Part 2 execution time: ", time.Since(startTime))

	return result
}

func part2Optim(filePath string) int {
	startTime := time.Now()
	var result int = 1

	times, distances := readFile(filePath)
	time1 := mergeIntoOneInt(times)
	distance := mergeIntoOneInt(distances)
	fmt.Println("check: ", time1, distance)

	winPresses := optimCalcWinPresses(time1, distance)
	result = winPresses
	
	fmt.Println("Part 2 optim result: ", result)
	fmt.Println("Part 2 optim execution time: ", time.Since(startTime))

	return result
}

func main() {
	filePath := "./input1.txt"
	part1(filePath)
	part2(filePath)
	part2Optim(filePath)
}
