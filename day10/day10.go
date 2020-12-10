package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func trib(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}
	return trib(n-1) + trib(n-2) + trib(n-3)
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sequence := []int{
		0,
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		sequence = append(sequence, num)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(sequence)
	sequence = append(sequence, sequence[len(sequence)-1]+3)

	oneDiffCount := 0
	threeDiffCount := 0
	parts := []int{
		1,
	}
	for i := 1; i < len(sequence); i++ {
		diff := sequence[i] - sequence[i-1]
		if diff == 1 {
			parts[threeDiffCount]++
			oneDiffCount++
		} else if diff == 3 {
			threeDiffCount++
			parts = append(parts, 1)
		} else {
			fmt.Printf("Unexpected diff of %d\n", diff)
		}
	}
	fmt.Println(oneDiffCount * threeDiffCount)
	permutations := 1
	for _, n := range parts {
		if n != 0 {
			permutations *= trib(n)
		}
	}
	fmt.Println(permutations)
}
