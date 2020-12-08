package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op    string
	val   int
	count int
}

func parseLine(line string) *instruction {
	parsed := strings.Split(line, " ")
	op := parsed[0]
	var rawVal string
	if parsed[1][:1] == "+" {
		rawVal = parsed[1][1:]
	} else {
		rawVal = parsed[1]
	}
	val, err := strconv.Atoi(rawVal)
	if err != nil {
		log.Fatal(err)
	}
	return &instruction{op, val, 0}
}

func runProgram(program []*instruction) (int, int) {
	pos := 0
	acc := 0
	for {
		if pos >= len(program) {
			break
		}
		cur := program[pos]
		if cur.count == 1 {
			break
		}
		cur.count++
		if cur.op == "acc" {
			acc += cur.val
			pos++
		} else if cur.op == "jmp" {
			pos += cur.val
		} else if cur.op == "nop" {
			pos++
		}
	}
	for _, cur := range program {
		cur.count = 0
	}
	return pos, acc
}

func fixProgram(program []*instruction) int {
	// Flip jmp to nop
	for i := 0; i < len(program); i++ {
		if program[i].op == "jmp" {
			program[i].op = "nop"
			pos, acc := runProgram(program)
			if pos == len(program) {
				return acc
			}
			program[i].op = "jmp"
		}
	}
	// Flip nop to jmp
	for i := 0; i < len(program); i++ {
		if program[i].op == "nop" {
			program[i].op = "jmp"
			pos, acc := runProgram(program)
			if pos == len(program) {
				return acc
			}
			program[i].op = "nop"
		}
	}
	return 0
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var program []*instruction

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		program = append(program, parseLine(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	_, acc := runProgram(program)
	fmt.Println(acc)

	fmt.Println(fixProgram(program))
}
