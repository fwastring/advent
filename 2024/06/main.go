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

	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	charMap := make([][]string, len(lines))

	for i, line := range lines {
		chars := strings.Split(line, "")
		charMap[i] = chars
	}

	isValid := true
	loops := 0
	counter := 0
	for outerI := range charMap {
		for outerJ := range charMap[outerI] {
			for i, line := range lines {
				chars := strings.Split(line, "")
				charMap[i] = chars
			}
			if !(charMap[outerI][outerJ] == "^") && !(charMap[outerI][outerJ] == "#") {
				charMap[outerI][outerJ] = "#"
				isValid = true
				loops = 0

				for isValid {
					for i := range charMap {
						for j := range charMap[i] {
							valid, changed := checkMap(charMap, i, j)
							isValid = valid
							if changed {
								after := charMap[i][j]
								if after == "X" {
									loops++
								}
							}
							if loops > 100 {
								fmt.Println(outerI, outerJ)
								isValid = false
								counter++
								break
							}
						}
						if isValid == false {
							break
						}
					}
				}
			}
		}
	}

	// Calculate and print execution time
	execution := time.Since(start)
	fmt.Printf("Execution time: %f ms\n", float64(execution.Microseconds())/1000)
	fmt.Printf("Amount: %d\n", counter)
}

func checkMap(charMap [][]string, i, j int) (bool, bool) {

	switch charMap[i][j] {
	case "^":
		if i == 0 {
			return false, false
		} else {
			if charMap[i-1][j] == "#" {
				charMap[i][j] = ">"
			} else {
				charMap[i-1][j] = "^"
				charMap[i][j] = "X"
			}
		}
	case ">":
		if j == len(charMap[i])-1 {
			return false, false
		} else {
			if charMap[i][j+1] == "#" {
				charMap[i][j] = "v"
			} else {
				charMap[i][j+1] = ">"
				charMap[i][j] = "X"
			}
		}
	case "<":
		if j == 0 {
			return false, false
		} else {
			if charMap[i][j-1] == "#" {
				charMap[i][j] = "^"
			} else {
				charMap[i][j-1] = "<"
				charMap[i][j] = "X"
			}
		}
	case "v":
		if i == len(charMap[i])-1 {
			return false, false
		} else {
			if charMap[i+1][j] == "#" {
				charMap[i][j] = "<"
			} else {
				charMap[i+1][j] = "v"
				charMap[i][j] = "X"
			}
		}
	default:
		return true, false
	}
	return true, true
}
