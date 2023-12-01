package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)


func main() {
	strs := parseInput("./day1/puzzle2/example.txt")
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

func getSum(str string) int {
	fmt.Println(str)

	var a, b string
	var c int

	for _, char := range str{


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

	fmt.Println(a+b)
	c, _ = strconv.Atoi(string(a + b))
	return c
}