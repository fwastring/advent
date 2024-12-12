package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)


func main() {
	start := time.Now()
    dat, _ := os.ReadFile("./input/input")

	data := string(dat)
	lines := strings.Split(data, "\n")

	re := regexp.MustCompile(`(?:mul\(\d*,\d*\)|do\(\)|don\'t\(\))`)	
	multiplications := make([]string, 0)

	// Find all multiplications and conditionals
	for _, line := range lines {
		matches := re.FindAll([]byte(line), -1)
		for _, match  := range matches {
			multiplications = append(multiplications, string(match))
		}
	}

	// Find the indexes to be removed w.r.t conditionals
	indexes := make([]int, 0)
	for i, mult := range multiplications {
		if mult == "don't()" {
			relativeEnd := slices.IndexFunc(multiplications[i:], func(s string) bool { return s == "do()"})
			if relativeEnd == -1 {
				indexes = append(indexes, makeRange(int(i), len(multiplications)-1)...)
			} else {
				end := relativeEnd + i
				fmt.Printf("start: %d, end: %d \n", i, end)
				indexes = append(indexes, makeRange(int(i), int(end)-1)...)
			}
		}
	}

	// Sum up the multiplications
	numbers := regexp.MustCompile(`\d{1,}`)
	var sum int64
	sum = 0
	for i, mult := range multiplications {
		if !slices.Contains(indexes, i) {
			if mult != "don't()" && mult != "do()" {
				matches := numbers.FindAll([]byte(mult), -1) 
				first, _ := strconv.ParseInt(string(matches[0]), 10, 0)
				second, _ := strconv.ParseInt(string(matches[1]), 10, 0)
				sum = sum + (first * second)
			}
		}
	}
	end := time.Now()	

	execution := end.Nanosecond()-start.Nanosecond()
	fmt.Printf("exectution time: %f ms\n", float64(execution)/1000000)
	fmt.Printf("amount: %d\n", sum)
}

func trim(slice []string, start int, end int) []string {
	listCopy := make([]string, len(slice))
	copy(listCopy, slice)
    return append(listCopy[:start], listCopy[end+1:]...)
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

