package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dominikbraun/graph"
)


func main() {
	start := time.Now()
	dat, _ := os.ReadFile("./input")

	data := string(dat)
	g := graph.New(graph.IntHash, graph.Directed())
	lines := strings.Split(data, "\n")
	rules := make([][]int, 0)
	instructions := make([][]int, 0)
	amount := 0

	for _, line := range lines {
		if strings.Contains(line, ",") {
			numbers := strings.Split(line, ",")
			intValues := make([]int, 0)
			for _, number := range numbers {
				intValue, _ := strconv.ParseInt(number, 10, 0)
				intValues = append(intValues, int(intValue))
			}
			instructions = append(instructions, intValues)
		}
		if strings.Contains(line, "|") {
			numbers := strings.Split(line, "|")
			intValues := make([]int, 0)
			for _, number := range numbers {
				intValue, _ := strconv.ParseInt(number, 10, 0)
				intValues = append(intValues, int(intValue))
			}
			rules = append(rules, intValues)
		}
	}

	for _, rule := range rules {
		g.AddVertex(rule[0])
		g.AddVertex(rule[1])
		g.AddEdge(rule[0], rule[1])
	}

	sum := 0
	for _, instruction := range instructions {
		if !checkInstructions(g, instruction) {
			// Reorder until the sequence is valid
			for {
				isValid := true
				for i := 1; i < len(instruction); i++ {
					path, err := graph.ShortestPath(g, instruction[i-1], instruction[i])
					if err != nil || len(path) == 0 || len(path) > 2 {
						// Swap adjacent elements to try reordering
						instruction[i], instruction[i-1] = instruction[i-1], instruction[i]
						isValid = false
					}
				}
				fmt.Println(instruction)
				if isValid {
					break
				}
			}

			// Add the middle value of the reordered sequence
			index := len(instruction) / 2
			sum += instruction[index]
			amount++
		}
	}

	// Calculate and print execution time
	execution := time.Since(start)
	fmt.Printf("Execution time: %f ms\n", float64(execution.Microseconds())/1000)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Amount: %d\n", amount)
	fmt.Printf("Total: %d\n", len(instructions))
}

func checkInstructions(g graph.Graph[int, int], instruction []int) bool {
	for i := range instruction {
		if i > 0 {
			for _, value := range instruction[i:] {
				path, err := graph.ShortestPath(g, instruction[i-1], value)
				if err != nil {
					return false
				}
				if len(path) > 2 {
					return false
				}
			}
		}
	} 
	return true
}
