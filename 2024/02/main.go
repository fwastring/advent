package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)


func main() {
	start := time.Now()
    dat, _ := os.ReadFile("./input/input")

	data := string(dat)
	lines := strings.Split(data, "\n")
	// lines := []string{"63 65 64"}
	counter := 0
	for _, line := range lines {
		levels := strings.Split(line, " ")
		levelsInt := make([]int64, 0)
		for _, level  := range levels {
			levelValue, _ := strconv.ParseInt(level, 10, 0)
			levelsInt = append(levelsInt, levelValue)	
		}
		if isSafe(levelsInt){
			fmt.Print(levelsInt)
			fmt.Print("\n")
			counter++
		} else {
			for i := range levelsInt {
				fmt.Print(i)
				fmt.Print(levelsInt)
				newLevels := remove(levelsInt, i)
				fmt.Print(newLevels)
				if isSafe(newLevels) {
					fmt.Print("\nold", levelsInt)
					fmt.Print("\nnew", newLevels)
					fmt.Print("\n")
					counter++
					break;
				}
			}
		}
	}
	end := time.Now()	

	execution := end.Nanosecond()-start.Nanosecond()
	fmt.Printf("exectution time: %f ms\n", float64(execution)/1000000)
	fmt.Printf("amount: %d\n", counter)
}

func remove(slice []int64, s int) []int64 {
	listCopy := make([]int64, len(slice))
	copy(listCopy, slice)
    return append(listCopy[:s], listCopy[s+1:]...)
}

func isSafe(list []int64) bool {
	return check(list) && (compare(list, sortList(list, true)) ||  compare(list, sortList(list, false)))
}

func compare(list1 []int64, list2 []int64) bool {
	for i := range list1 {
		first := list1[i]
		second := list2[i]
		if first != second {
			return false
		}
	}
	return true
}

func check(list []int64) bool {
	for i := range list {
		if i != 0 {
			if absolute(list[i], list[i-1]) > 3 || absolute(list[i], list[i-1]) < 1 {
				return false
			}
		}
	}
	return true
}

func absolute(value1 int64, value2 int64) int64 {
	return int64(math.Abs(float64(value1 - value2)))
}

func sortList(list []int64, ascending bool) []int64 {
	listCopy := make([]int64, len(list))
	copy(listCopy, list)
	sort.Slice(listCopy, func(i, j int) bool {
		if ascending {
			return list[i] < list[j]
		} else {
			return list[i] > list[j]
		}
	})
	return listCopy
}

