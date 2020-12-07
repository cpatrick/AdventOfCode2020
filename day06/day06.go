package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func makeLetterSet(line string) map[rune]struct{} {
	letters := make(map[rune]struct{})
	for _, char := range line {
		letters[char] = struct{}{}
	}
	return letters
}

func countLetters(line string) int {
	letters := makeLetterSet(line)
	return len(letters)
}

func countSharedLetters(group []string) int {
	curSet := makeLetterSet(group[0])
	for index, curGroup := range group {
		if index == 0 {
			continue
		}
		newSet := make(map[rune]struct{})
		trySet := makeLetterSet(curGroup)
		for key := range trySet {
			_, ok := curSet[key]
			if ok {
				newSet[key] = struct{}{}
			}
		}
		curSet = newSet
	}
	return len(curSet)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var parsedLines []string
	var groups [][]string
	curGroup := []string{}
	curLine := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			groups = append(groups, curGroup)
			parsedLines = append(parsedLines, curLine)
			curGroup = []string{}
			curLine = ""
			continue
		}
		curGroup = append(curGroup, line)
		if curLine == "" {
			curLine = line
		} else {
			curLine = curLine + line
		}
	}
	groups = append(groups, curGroup)          // get last line
	parsedLines = append(parsedLines, curLine) // get last line (again)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalLetters := 0
	for _, line := range parsedLines {
		totalLetters += countLetters(line)
	}
	fmt.Println(totalLetters)
	sharedLetters := 0
	for _, group := range groups {
		sharedLetters += countSharedLetters(group)
	}
	fmt.Println(sharedLetters)
}
