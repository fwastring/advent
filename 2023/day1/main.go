package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    dat, err := os.ReadFile("./input/example")
	if err != nil {

	}
	data := string(dat)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		letters := make([]string, 0)
		for _, letter := range line {
			if isDigit(letter) {
				letters = append(letters, string(letter))
			}
		}
		fmt.Println(letters)
	}
}

func isDigit(letter rune) bool {
	return (letter > 48 && letter < 60)
}
