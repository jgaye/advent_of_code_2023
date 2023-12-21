package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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

	return txtlines
}

func parseFields(inputs []string) []field {
	fields := make([]field, 0)
	pageIndex := 0
	
	for i := 0; i < len(inputs); i++ {
		if inputs[i] == "" {
			previousLine := inputs[i-1]
			newField := field{id: pageIndex, rows: inputs[pageIndex:i]}
			columns := make([]string, len(previousLine))
			for j := 0; j < len(previousLine); j++ {
				column := ""
				for k := 0; k < len(newField.rows); k++ {
					column += string(newField.rows[k][j])
				}
				columns[j] = column
			}
			newField.columns = columns
			pageIndex = i + 1
			fields = append(fields, newField)
		}
	}

	return fields

}

type field struct {
	id int
	rows []string
	columns []string
}

func (f field) print() {
	fmt.Println("Field", f.id)
	for _, row := range f.rows {
		fmt.Println(row)
	}
}

func (field field) tiltNorth() {
	for i, column := range field.columns {

		columnA := strings.Split(column, "")

		// double loop because moving the boulders north once every loop
		for j := 1; j < len(columnA); j++ {
			for i := 1; i < len(columnA); i++ {
				if columnA[i] == "O" && columnA[i-1] == "." {
					columnA[i-1] = "O"
					columnA[i] = "."
				}
			}
		}

		field.columns[i] = strings.Join(columnA, "")
	}

	for l := 0; l < len(field.rows); l++ {
		row := ""
		for k := 0; k < len(field.columns); k++ {
			row += string(field.columns[k][l])
		}
		field.rows[l] = row
	}
}

func (field field) tiltSouth() {
	for i, column := range field.columns {

		columnA := strings.Split(column, "")
		for j := 1; j < len(columnA); j++ {
			for i := 1; i < len(columnA); i++ {
				if columnA[i] == "O" && columnA[i-1] == "." {
					columnA[i-1] = "O"
					columnA[i] = "."
				}
			}
		}

		field.columns[i] = strings.Join(columnA, "")
	}

	for l := 0; l < len(field.rows); l++ {
		row := ""
		for k := 0; k < len(field.columns); k++ {
			row += string(field.columns[k][l])
		}
		field.rows[l] = row
	}
}

func (field field) countWeight() int{
	acc := 0
	for i:=len(field.rows); i >0; i-- {
		row := field.rows[len(field.rows)-i]
		acc+= strings.Count(row, "O")*i
	}
	return acc
}

// END HELPERS

func part1(filePath string) int {
	startTime := time.Now()
	result := 0

	inputs := parseFile(filePath)
	fields := parseFields(inputs)
	for _, field := range fields {
		field.print()
		field.tiltNorth()
		field.print()
		result += field.countWeight()
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

