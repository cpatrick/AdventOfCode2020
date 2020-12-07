package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func getRow(line string) int {
	lower := 0
	upper := 127
	for _, c := range line {
		size := (upper - lower) + 1
		if c == 'F' {
			upper = upper - size/2
		} else if c == 'B' {
			lower = lower + size/2
		}
	}
	return lower
}

func getColumn(line string) int {
	lower := 0
	upper := 7
	for _, c := range line {
		size := (upper - lower) + 1
		if c == 'L' {
			upper = upper - size/2
		} else if c == 'R' {
			lower = lower + size/2
		}
	}
	return lower
}

func parseLine(line string) int {
	return getRow(line[:7])*8 + getColumn(line[7:])
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maxID := 0
	var IDs []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curID := parseLine(scanner.Text())
		IDs = append(IDs, curID)
		if curID > maxID {
			maxID = curID
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(maxID)
	sort.Ints(IDs)
	for index := range IDs {
		if index < 8 {
			continue
		}
		missing := IDs[index-1] + 1
		if missing != IDs[index] {
			fmt.Println(missing)
			break
		}
	}
}
