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

type Galaxy struct {
	y int
	x int
}

func (g Galaxy) Distance(other Galaxy) int {
	return tools.IAbs(g.y-other.y) + tools.IAbs(g.x-other.x)
}

func galaxyEquals(galaxy1 Galaxy, galaxy2 Galaxy) bool {
	return galaxy1.y == galaxy2.y && galaxy1.x == galaxy2.x
}

func parseGalaxies(txtlines []string) []Galaxy {
	galaxies := []Galaxy{}
	for y, eachline := range txtlines {
		asChars := strings.Split(eachline, "")
		for x, eachchar := range asChars {
			if eachchar == "#" {
				galaxies = append(galaxies, Galaxy{y, x})
			}
		}
	}
	return galaxies
}

func parseFile(filename string) (int, int, []Galaxy) {
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

	return len(txtlines), len(txtlines[0]), parseGalaxies(txtlines)
}

func xComparison(galaxy1 Galaxy, galaxy2 Galaxy) int {
	if galaxy1.x == galaxy2.x {
		return 0
	}
	if galaxy1.x > galaxy2.x {
		return 1
	}
	return -1
}

func yComparison(galaxy1 Galaxy, galaxy2 Galaxy) int {
	if galaxy1.y == galaxy2.y {
		return 0
	}
	if galaxy1.y > galaxy2.y {
		return 1
	}
	return -1
}

func expandY(maxY int, galaxies []Galaxy, expansionRate int) []Galaxy {
	slices.SortFunc[[]Galaxy](galaxies, yComparison)
	seenGalaxies := []Galaxy{}

	for i := maxY; i > 0; i-- {
		galaxyOnY := false

		for j := len(galaxies); j > 0; j-- {
			if galaxies[j-1].y == i {
				seenGalaxies = append(seenGalaxies, galaxies[j-1])
				galaxies = slices.Delete(galaxies, j-1, j)
				galaxyOnY = true
			}
		}
		if !galaxyOnY {
			for k, seenGalaxy := range seenGalaxies {
				seenGalaxies[k] = Galaxy{seenGalaxy.y + expansionRate, seenGalaxy.x}
			}
		}
	}

	return append(seenGalaxies, galaxies...)
}

func expandX(maxX int, galaxies []Galaxy, expansionRate int) []Galaxy {
	slices.SortFunc[[]Galaxy](galaxies, xComparison)
	seenGalaxies := []Galaxy{}

	for i := maxX; i > 0; i-- {
		galaxyOnX := false

		for j := len(galaxies); j > 0; j-- {
			if galaxies[j-1].x == i {
				seenGalaxies = append(seenGalaxies, galaxies[j-1])
				galaxies = slices.Delete(galaxies, j-1, j)
				galaxyOnX = true
			}
		}
		if !galaxyOnX {
			for k, seenGalaxy := range seenGalaxies {
				seenGalaxies[k] = Galaxy{seenGalaxy.y, seenGalaxy.x + expansionRate}
			}
		}
	}
	return append(seenGalaxies, galaxies...)
}

func expandUniverse(maxY int, maxX int, galaxies []Galaxy, expansionRate int) []Galaxy {
	galaxies = expandY(maxY, galaxies, expansionRate)
	galaxies = expandX(maxX, galaxies, expansionRate)
	return galaxies
}

// END HELPERS

func part1(filePath string) int {
	startTime := time.Now()
	result := 0
	maxY, maxX, galaxies := parseFile(filePath)

	expansionRate := 1
	expandedUniverse := expandUniverse(maxY, maxX, galaxies, expansionRate)
	fmt.Println("expandedUniverse: ", expandedUniverse)

	for i := 0; i < len(expandedUniverse); i++ {
		for j := i + 1; j < len(expandedUniverse); j++ {
			result += expandedUniverse[i].Distance(expandedUniverse[j])
		}
	}

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
