package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/jgaye/adventOfCode2023/tools"
)

// HELPERS

type field struct {
	id int
	rows []string
	columns []string
}

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

func checkRowSymmetry(up []string, down []string) bool {
	// slices.Reverse(up) // SOMEHOW ONE CASE WHERE REVERSE WAS WONRG (with duplicated elements)
	sizeToCheck := tools.Min(len(up), len(down))

	for i:=0; i<sizeToCheck; i++ {
		if up[len(up)-1-i] != down[i] {
			return false
		}
	}

	return true
}

func checkColumnSymmetry(left []string, right []string) bool {
	// slices.Reverse(left)
	sizeToCheck := tools.Min(len(left), len(right))

	for i:=0; i<sizeToCheck; i++ {
		if left[len(left)-1-i] != right[i] {
			return false
		}
	}

	return true
}

// END HELPERS

func part1(filePath string) int {
	startTime := time.Now()
	result := 0

	inputs := parseFile(filePath)
	fields := parseFields(inputs)

	for _, field := range fields {
		fieldResult := 0
		for j:=0; j<len(field.columns)-1; j++ {
			if field.columns[j] == field.columns[j+1] {
				if checkColumnSymmetry(field.columns[0:j+1], field.columns[j+1:len(field.columns)]){
					fieldResult = j+1
					// fmt.Printf("field %d, symmetry on column %d, adding %d\n", field.id, j+1, fieldResult)
					break
				}
			}
		}
		for i:=0; i<len(field.rows)-1; i++ {
			if field.rows[i] == field.rows[i+1] {
				if checkRowSymmetry(field.rows[0:i+1], field.rows[i+1:len(field.rows)]){
					fieldResult = 100*(i+1)
					// fmt.Printf("field %d, symmetry on row %d, adding %d\n", field.id, i+1, fieldResult)
					break
				}
			}
		}
		if fieldResult == 0 {
			fmt.Printf("FOR field %d, NO symmetry\n", field.id)
		}
		result += fieldResult
	}
	// fmt.Println("var: ", fields)
	
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
	// part2(filePath)
}

