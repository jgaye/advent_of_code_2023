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

// setting up optimized execution for part 2
var m = sync.Mutex{}
var wg = sync.WaitGroup{}

// global maps
var seedToSoil []Range
var soilToFertilizer []Range
var fertilizerToWater []Range
var waterToLight []Range
var lightToTemperature []Range
var temperatureToHumidity []Range
var humidityToLocation []Range

// global result
var minLocs []int

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

type Range struct {
	source      int
	destination int
	r           int
}

func (el Range) find(source int) (int, bool) {
	if el.source <= source && source <= el.source+el.r {
		return el.destination + (source - el.source), true
	} else {
		return source, false
	}
}

func parseMaps(txtlines []string) {
	var currentMap *[]Range
	seedToSoil = []Range{}
	soilToFertilizer = []Range{}
	fertilizerToWater = []Range{}
	waterToLight = []Range{}
	lightToTemperature = []Range{}
	temperatureToHumidity = []Range{}
	humidityToLocation = []Range{}

	for _, line := range txtlines {
		if line == "seed-to-soil map:" {
			currentMap = &seedToSoil
		} else if line == "soil-to-fertilizer map:" {
			currentMap = &soilToFertilizer
		} else if line == "fertilizer-to-water map:" {
			currentMap = &fertilizerToWater
		} else if line == "water-to-light map:" {
			currentMap = &waterToLight
		} else if line == "light-to-temperature map:" {
			currentMap = &lightToTemperature
		} else if line == "temperature-to-humidity map:" {
			currentMap = &temperatureToHumidity
		} else if line == "humidity-to-location map:" {
			currentMap = &humidityToLocation
		} else if line == "" {
			continue
		} else if strings.HasPrefix(line, "seeds") {
			continue
		} else {
			lsplit := strings.Split(line, " ")
			source, _ := strconv.Atoi(lsplit[1])
			destination, _ := strconv.Atoi(lsplit[0])
			r, _ := strconv.Atoi(lsplit[2])
			*currentMap = append(*currentMap, Range{source, destination, r})
		}
	}
}

// TODO: make this mare readable with a generic loop and before/after vars
func seedToLocation(seed int) int {
	fmt.Println("seed: ", seed)

	soil := seed
	for _, el := range seedToSoil {
		res, found := el.find(seed)
		if found {
			soil = res
			break
		}
	}
	fmt.Println("soil: ", soil)

	fertilizer := soil
	for _, el := range soilToFertilizer {
		res, found := el.find(soil)
		if found {
			fertilizer = res
			break
		}
	}
	fmt.Println("fertilizer: ", fertilizer)

	water := fertilizer
	for _, el := range fertilizerToWater {
		res, found := el.find(fertilizer)
		if found {
			water = res
			break
		}
	}
	fmt.Println("water: ", water)

	light := water
	for _, el := range waterToLight {
		res, found := el.find(water)
		if found {
			light = res
			break
		}
	}
	fmt.Println("light: ", light)

	temperature := light
	for _, el := range lightToTemperature {
		res, found := el.find(light)
		if found {
			temperature = res
			break
		}
	}
	fmt.Println("temperature: ", temperature)

	humidity := temperature
	for _, el := range temperatureToHumidity {
		res, found := el.find(temperature)
		if found {
			humidity = res
			break
		}
	}
	fmt.Println("humidity: ", humidity)

	location := humidity
	for _, el := range humidityToLocation {
		res, found := el.find(humidity)
		if found {
			location = res
			break
		}
	}
	fmt.Println("location: ", location)

	return location
}

func parseSeeds(txtlines []string) []int {
	a, _ := strings.CutPrefix(txtlines[0], "seeds: ")
	b := strings.Split(a, " ")
	seeds := make([]int, len(b))
	for i, seed := range b {
		seedInt, _ := strconv.Atoi(seed)
		seeds[i] = seedInt
	}
	return seeds
}

func min(s []int) int {
	min := s[0]
	for _, el := range s {
		if el < min {
			min = el
		}
	}
	return min
}

func getMinLocOfSeedRange(seedStart int, seedRange int) {
	minLoc := 9999999999
	for seed := seedStart; seed < seedStart+seedRange; seed++ {
		location := seedToLocation(seed)
		if location < minLoc {
			minLoc = location
		}
	}

	m.Lock()
	minLocs = append(minLocs, minLoc)
	m.Unlock()

	wg.Done()
}

// END HELPERS

func part1(fileLines []string) {
	seeds := parseSeeds(fileLines)
	acc := []int{}
	for _, seed := range seeds {
		location := seedToLocation(seed)
		acc = append(acc, location)
		fmt.Println("seed", seed, "location", location)
	}

	fmt.Println("min location", min(acc))
}

func part2(fileLines []string) {
	a, _ := strings.CutPrefix(fileLines[0], "seeds: ")
	b := strings.Split(a, " ")

	minLocs = make([]int, 0)

	for i := 0; i < len(b); i = i + 2 {
		seedStart, _ := strconv.Atoi(b[i])
		seedRange, _ := strconv.Atoi(b[i+1])

		wg.Add(1)
		go getMinLocOfSeedRange(seedStart, seedRange)
	}
	wg.Wait()
	fmt.Println("min locations", minLocs)
	fmt.Println("final result", min(minLocs))
}

func main() {
	startTime := time.Now()

	// fileLines := readFile("./example1.txt")
	fileLines := readFile("./input1.txt")
	parseMaps(fileLines)

	part1(fileLines)
	part1Time := time.Now()
	fmt.Println("Part 1 execution time: ", time.Since(startTime))
	part2(fileLines)
	fmt.Println("Part 2 execution time: ", time.Since(part1Time))
}
