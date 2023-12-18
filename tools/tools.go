package tools

import (
	"strings" 
	"strconv"
)

func SumS(slice []int) int{
	r := 0
	for _, a := range slice{
		r += a
	}
	return r
}

func SAtoi(slice string) []int{
	r := []int{}
	x := strings.Split(slice, " ")
	for i:=0; i<len(x); i++{
		x, _ := strconv.Atoi(x[i])
		r = append(r, x)
	}
	return r
}