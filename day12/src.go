package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/jgaye/adventOfCode2023/tools"
)

// HELPERS

type line struct {
	sentence string
	clues    []int
}

func (l line) Equals(other line) bool {
	return l.sentence == other.sentence && slices.Equal(l.clues, other.clues)
}

func parseFile(filename string) []line {
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

	lines := []line{}
	for _, eachline := range txtlines {
		arr := strings.Split(eachline, " ")
		sentence, cluesS := arr[0], arr[1]
		clues := tools.SAtoi(cluesS, ",")
		lines = append(lines, line{sentence, clues})
	}

	return lines
}

func analyze(txtlines []string) {
	minSentenceLength := 99
	maxSentenceLength := 0
	minCluesLength := 99
	maxCluesLength := 0
	minCluesSum := 99
	maxCluesSum := 0

	for _, eachline := range txtlines {
		arr := strings.Split(eachline, " ")
		sentence, cluesS := arr[0], arr[1]
		clues := tools.SAtoi(cluesS, ",")

		sentenceLength := len(sentence)
		if sentenceLength > maxSentenceLength {
			maxSentenceLength = sentenceLength
		} else if sentenceLength < minSentenceLength {
			minSentenceLength = sentenceLength
		}

		cluesLength := len(clues)
		if cluesLength > maxCluesLength {
			maxCluesLength = cluesLength
		} else if cluesLength < minCluesLength {
			minCluesLength = cluesLength
		}

		cluesSum := tools.SumS(clues)
		if cluesSum > maxCluesSum {
			maxCluesSum = cluesSum
		} else if cluesSum < minCluesSum {
			minCluesSum = cluesSum
		}

	}

	fmt.Println("Analyzis maxSentenceLength: ", maxSentenceLength,
		" minSentenceLength: ", minSentenceLength,
		" maxCluesLength: ", maxCluesLength,
		" minCluesLength: ", minCluesLength,
		" maxCluesSum: ", maxCluesSum,
		" minCluesSum: ", minCluesSum)
}

func permuteSentence(sentence string) []string {
	permutations := []string{}

	// end of recursive loop
	if len(sentence) == 1 {
		if sentence == "?" {
			permutations = append(permutations, ".")
			permutations = append(permutations, "#")
		} else {
			permutations = append(permutations, sentence)
		}
		return permutations
	}

	head := sentence[0:1]
	tail := sentence[1:]

	for _, tailPerm := range permuteSentence(tail) {
		if head == "?" {
			permutations = append(permutations, "."+tailPerm)
			permutations = append(permutations, "#"+tailPerm)
		} else {
			permutations = append(permutations, head+tailPerm)
		}
	}
	return permutations
}

func isSentenceValid(sentence string, clues []int) bool {
	f := func(c rune) bool {
		return c == '.'
	}

	splittedS := strings.FieldsFunc(sentence, f)
	splittedSLen := []int{}
	for _, s := range splittedS {
		splittedSLen = append(splittedSLen, len(s))
	}

	if len(splittedSLen) == 0 {
		splittedSLen = []int{0}
	}

	return slices.Equal(clues, splittedSLen)

}

func bruteForceSolve(lines []line) int {
	acc := 0
	for _, line := range lines {
		accLog := 0
		permutations := permuteSentence(line.sentence)
		for _, permutation := range permutations {
			if isSentenceValid(permutation, line.clues) {
				acc += 1
				accLog += 1
			}
		}
		fmt.Println("Sentence: ", line.sentence, " Clues: ", line.clues, " Valid permutations: ", accLog)
	}
	return acc
}

func expandLine(line line) line {
	sInit := line.sentence
	cInit := line.clues
	for i:=0; i<4; i++ {
		line.sentence += "?" + sInit
		line.clues = append(line.clues, cInit...)
	}
	return line
}


// END HELPERS

func part1(filePath string) int {
	startTime := time.Now()
	result := 0

	lines := parseFile(filePath)
	result = bruteForceSolve(lines)

	fmt.Println("Part 1 results: ", result)
	fmt.Println("Part 1 execution time: ", time.Since(startTime))

	return result
}

func part2(filePath string) int {
	// TODO brute force takes too long just to permutate the first line
	startTime := time.Now()
	result := 0

	lines := parseFile(filePath)
	for i, line := range lines {
		line = expandLine(line)
		lines[i] = line
	}
	result = bruteForceSolve(lines)

	fmt.Println("Part 2 result: ", result)
	fmt.Println("Part 2 execution time: ", time.Since(startTime))

	return result
}

func main() {
	filePath := "./input.txt"
	// txtlines := parseFile(filePath)
	// analyze(txtlines)
	// //Analyzis maxSentenceLength:  20  minSentenceLength:  10  maxCluesLength:  6  minCluesLength:  2  maxCluesSum:  18  minCluesSum:  2
	part1(filePath)
	// part2(filePath)
}
