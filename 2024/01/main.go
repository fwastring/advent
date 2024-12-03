package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Id int64
}

func main() {
	start := time.Now()
    dat, err := os.ReadFile("./input/input")
	if err != nil {

	}
	data := string(dat)
	lines := strings.Split(data, "\n")
	list1 := make([]int64, 5)
	list2 := make([]int64, 5)
	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		if len(numbers) > 1 {
			first, _ := strconv.ParseInt(numbers[0], 10, 0)
			second, _ := strconv.ParseInt(numbers[1], 10, 0)
			list1 = append(list1, first)
			list2 = append(list2, second)
		}
	}
	freqMap := make(map[Entry]int64)
	for _, id := range list2 {
		freqMap[Entry{Id: id}]++
	}
	counter := int64(0)
	for _, id := range list1 {
		counter = counter + id * freqMap[Entry{Id: id}]	
	}

	end := time.Now()	

	execution := end.Nanosecond()-start.Nanosecond()
	fmt.Print(counter)
	fmt.Printf("exectution time: %d", execution)
}
