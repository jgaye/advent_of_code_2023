import os

template_lines = """
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// HELPERS

func parseFile(filename string) {
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
}

// END HELPERS

func part1(filePath string) int {
	startTime := time.Now()
	result := 0
	
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
	filePath := "./input1.txt"
	part1(filePath)
	part2(filePath)
}

"""

def main():
	for i in range(9,25):
		if not  os.path.isdir(f"day{i}"):
			os.mkdir(f"day{i}")
		lines = []
		with open(f"day{i}/src.go", "w") as f:
			f.write(template_lines)
		with open(f"day{i}/src_test.go", "w") as f:
                        pass
		with open(f"day{i}/input.txt", "w") as f:
                        pass
		with open(f"day{i}/example.txt", "w") as f:
                        pass  

if __name__ == "__main__":
    main()
