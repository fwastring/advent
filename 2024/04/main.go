package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	dat, _ := os.ReadFile("./test")

	data := string(dat)
	lines := strings.Split(data, "\n")
	charMap := make([][]string, len(lines))

	// Create matrix
	for i, line := range lines {
		charMap[i] = strings.Split(line, "")
	}

	// Word search
	counter := 0
	for row := range charMap {
		for col, char := range charMap[row] {
			if char == "A" {
				if checkCross(charMap, row, col) {
					counter++
				}
			}
		}
	}

	// Calculate and print execution time
	execution := time.Since(start)
	fmt.Printf("Execution time: %f ms\n", float64(execution.Microseconds())/1000)
	fmt.Printf("Amount: %d\n", counter)
}

func checkCross(charMap [][]string, row, col int) bool {
	directions := []int{-1, 1}
	mCounter := 0
	sCounter := 0
	if row == 0 || col == 0 || row == len(charMap)-2 || col == len(charMap)-1 {
		return false
	}
	for _, i := range directions {
		for _, j := range directions {
			x := row+i
			y := col+j
			if charMap[x][y] == "M" {
				mCounter++
			}
			if charMap[x][y] == "S" {
				sCounter++
			}
		}
	}
	if mCounter == 2 && sCounter == 2 {
		if charMap[row-1][col-1] == charMap[row+1][col+1] ||  charMap[row-1][col+1] == charMap[row+1][col-1] {
			return false
		}
		return true
	}
	return false
}
