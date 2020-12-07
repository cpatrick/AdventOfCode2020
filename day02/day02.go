package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func validatePassword(password string, char string, lower int, upper int) bool {
	count := 0
	for _, c := range password {
		if string(c) == char {
			count++
		}
	}
	return lower <= count && count <= upper
}

func validatePasswordNew(password string, char string, lower int, upper int) bool {
	lowChar := string(password[lower-1])
	upChar := string(password[upper-1])
	if lowChar == char && upChar == char {
		return false
	} else if lowChar == char {
		return true
	} else if upChar == char {
		return true
	}
	return false
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0
	newCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ": ")
		password := split[1]
		exp := strings.Split(split[0], " ")
		char := exp[1]
		bounds := strings.Split(exp[0], "-")
		lower, err := strconv.Atoi(bounds[0])
		if err != nil {
			log.Fatal(err)
		}
		upper, err := strconv.Atoi(bounds[1])
		if err != nil {
			log.Fatal(err)
		}
		if validatePassword(password, char, lower, upper) {
			count++
		}
		if validatePasswordNew(password, char, lower, upper) {
			newCount++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	fmt.Println(newCount)
}
