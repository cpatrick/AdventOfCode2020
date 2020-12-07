package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bag struct {
	name     string
	contains map[string]int
}

func findContainingBags(bags map[string]bag, curBag bag, target string) int {
	count := 0
	for color := range curBag.contains {
		subBag := bags[color]
		if color == target {
			count++
		} else {
			count += findContainingBags(bags, subBag, target)
		}
	}
	return count
}

func countBags(bags map[string]bag, curBag bag) int {
	count := 0
	for color, num := range curBag.contains {
		subBag := bags[color]
		count += num + num*countBags(bags, subBag)
	}
	return count
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bags := make(map[string]bag)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " bags contain ")
		color := tokens[0]
		subTokens := strings.Split(tokens[1], ", ")
		curBag := bag{color, make(map[string]int)}
		for _, st := range subTokens {
			if st == "no other bags." {
				break
			}
			words := strings.Split(st, " ")
			containedColor := words[1] + " " + words[2]
			num, err := strconv.Atoi(words[0])
			if err != nil {
				log.Fatal(err)
			}
			curBag.contains[containedColor] = num
		}
		bags[color] = curBag
	}

	pathsToBag := 0
	target := "shiny gold"
	for _, curBag := range bags {
		if curBag.name == target {
			continue
		}
		totalPaths := findContainingBags(bags, curBag, target)
		if totalPaths > 0 {
			pathsToBag++
		}
	}
	fmt.Println(pathsToBag)
	fmt.Println(countBags(bags, bags[target]))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
