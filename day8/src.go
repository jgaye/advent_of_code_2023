package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Node struct {
	Left  string
	Right string
}

type Tree struct {
	Nodes map[string]Node
}

// HELPERS

func parseFile(filename string) (Tree, []string) {
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

	instructions := strings.Split(txtlines[0], "")
	tree := Tree{make(map[string]Node)}
	for _, line := range txtlines[2:] {
		node := Node{line[7:10], line[12:15]}
		tree.Nodes[line[0:3]] = node
	}

	return tree, instructions
}

func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers []int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i], integers[i+1:])
	}

	return result
}

// END HELPERS

func part1(filePath string) int {
	startTime := time.Now()
	tree, instructions := parseFile(filePath)
	nbSteps := 0
	currentNode := "AAA"
	fmt.Printf("START : new currentNode %s, nbSteps %d \n", currentNode, nbSteps)

	for currentNode != "ZZZ" {
		for _, instruction := range instructions {
			nbSteps += 1
			if instruction == "L" {
				currentNode = tree.Nodes[currentNode].Left
			} else {
				currentNode = tree.Nodes[currentNode].Right
			}
			fmt.Printf("new currentNode %s, nbSteps %d \n", currentNode, nbSteps)
		}
	}

	fmt.Printf("Part 1 results %d: \n", nbSteps)
	fmt.Println("Part 1 execution time: ", time.Since(startTime))

	return nbSteps
}

func part2(filePath string) int {
	startTime := time.Now()
	tree, instructions := parseFile(filePath)
	nbSteps := 0
	nbLoops := 0

	startNodes := []string{}
	for key := range tree.Nodes {
		if key[2] == 'A' {
			startNodes = append(startNodes, key)
		}
	}
	currentNodes := make([]string, len(startNodes))
	copy(currentNodes, startNodes)

	nbGhosts := len(startNodes)
	ghostSteps := make([]int, nbGhosts)

	done := false
	for !done  {

		for _, instruction := range instructions {
			nbSteps += 1

			if instruction == "L" {
				for i, currentNode := range currentNodes {
					currentNodes[i] = tree.Nodes[currentNode].Left
				}
			} else {
				for i, currentNode := range currentNodes {
					currentNodes[i] = tree.Nodes[currentNode].Right
				}
			}

			// fmt.Printf("new currentNodes %s, nbSteps %d \n", currentNodes, nbSteps)

			// end for all ghost, but we won't reach that
			for _, v := range currentNodes {
				if v[2] == 'Z' {
					done = true
				} else {
					done = false
					break
				}
			}

			// check if we have a loop for a ghost
			for i, currentNode := range currentNodes {
				if ghostSteps[i] == 0 && currentNode[2] == 'Z' {
					fmt.Printf("Ghost %d completed in %d steps \n", i, nbSteps)
					ghostSteps[i] = nbSteps
				}
			}

			// end when all loops are founds
			for _, v := range ghostSteps {
				if v == 0 {
					done = false
					break
				} else {
					done = true
				}
			}
			
		}
		nbLoops += 1

	}

	// Now we calculate the least common multiple of the ghost loops
	// I just took that from the internet
	result := LCM(ghostSteps[0], ghostSteps[1], ghostSteps[2:])

	fmt.Println("Part 2 result: ", result)
	fmt.Println("Part 2 execution time: ", time.Since(startTime))

	return nbSteps
}

func main() {
	filePath := "./input1.txt"
	part1(filePath)
	part2(filePath)
}
