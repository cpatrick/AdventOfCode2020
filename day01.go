package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var nums []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		for _, j := range nums {
			for _, k := range nums {
				if i+j+k == 2020 {
					fmt.Println(i * j * k)
				}
			}
		}
		nums = append(nums, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
