package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	right, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	down, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	treesFound := 0
	curRight := 0
	verbose := true

	scanner := bufio.NewScanner(file)
	// skip first "down" lines
	for i := 0; i < down; i++ {
		scanner.Scan()
		if verbose {
			fmt.Println(scanner.Text())
		}
	}
	for scanner.Scan() {
		curRight += right
		line := scanner.Text()
		if curRight >= len(line) {
			curRight = curRight - len(line)
		}
		if string(line[curRight]) == "#" {
			line = line[:curRight] + "X" + line[curRight+1:]
			treesFound++
		} else {
			line = line[:curRight] + "O" + line[curRight+1:]
		}
		if verbose {
			fmt.Println(line)
		}
		for i := 0; i < down-1; i++ {
			scanner.Scan()
			if verbose {
				fmt.Println(scanner.Text())
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(treesFound)
}
