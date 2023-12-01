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
	strs := parseInput("./input.txt")
	sums := []int{}

	for _, str := range strs {
		sums = append(sums, getSum(str))
	}

	sum := 0
	for _, num := range sums {
		sum += num
	}
	fmt.Println(sum)

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

func strToStrint(str string) string {
	switch str {
	case "one":
		return "o1ne"
	case "two":
		return "t2wo"
	case "three":
		return "t3hree"
	case "four":
		return "f4our"
	case "five":
		return "f5ive"
	case "six":
		return "s6ix"
	case "seven":
		return "s7even"
	case "eight":
		return "e8ight"
	case "nine":
		return "n9ine"
	default:
		return "0"
	}
}

func getSum(str string) int {
	fmt.Println(str)

	var a, b string
	var c int

	var numbersAsStr = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, nas := range numbersAsStr {
		str = strings.Replace(str, nas, strToStrint(nas), -1)
	}

	fmt.Println(str)

	for _, char := range str {
		_, err := strconv.Atoi(string(char))
		if err == nil {
			a = string(char)
			fmt.Println(a)
			break
		}
	}

	for i := len(str) - 1; i >= 0; i-- {
		char := rune(str[i])
		_, err := strconv.Atoi(string(char))
		if err == nil {
			b = string(char)
			fmt.Println(b)
			break
		}
	}

	fmt.Println(a + b)
	c, _ = strconv.Atoi(string(a + b))
	return c
}
