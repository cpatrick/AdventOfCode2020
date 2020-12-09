package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func verifyNumber(stack []int, num int) bool {
	for i := 0; i < len(stack); i++ {
		for j := 1; j < len(stack); j++ {
			if stack[i] == stack[j] {
				continue
			}
			if stack[i]+stack[j] == num {
				return true
			}
		}
	}
	return false
}

func sum(buf []int) int {
	res := 0
	for _, num := range buf {
		res += num
	}
	return res
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var stack []int
	var badNumber int
	var sequence []int

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

	for _, num := range sequence {
		if len(stack) < 25 {
			stack = append(stack, num)
			continue
		}
		if !verifyNumber(stack, num) {
			badNumber = num
			break
		}
		stack = append(stack[1:], num)
	}
	fmt.Println(badNumber)

	var buf []int
	for _, num := range sequence {
		buf = append(buf, num)
		if sum(buf) == badNumber {
			sort.Ints(buf)
			fmt.Println(buf[0] + buf[len(buf)-1])
			break
		} else if (sum(buf)) < badNumber {
			continue
		} else { // sum(buf) > badNumber
			for sum(buf) > badNumber {
				buf = buf[1:]
			}
			if sum(buf) == badNumber {
				sort.Ints(buf)
				fmt.Println(buf[0] + buf[len(buf)-1])
				break
			}
		}
	}
}
