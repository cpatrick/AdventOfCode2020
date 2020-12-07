package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func fullyValidate(line string) bool {
	required := map[string]func(string) bool{
		"byr": func(val string) bool {
			num, err := strconv.Atoi(val)
			if err != nil {
				return false
			}
			return num >= 1920 && num <= 2002
		},
		"iyr": func(val string) bool {
			num, err := strconv.Atoi(val)
			if err != nil {
				return false
			}
			return num >= 2010 && num <= 2020
		},
		"eyr": func(val string) bool {
			num, err := strconv.Atoi(val)
			if err != nil {
				return false
			}
			return num >= 2020 && num <= 2030
		},
		"hgt": func(val string) bool {
			re := regexp.MustCompile(`^([0-9]+)(cm|in)$`)
			matches := re.FindStringSubmatch(val)
			if matches == nil {
				return false
			}
			num, err := strconv.Atoi(matches[1])
			if err != nil {
				return false
			}
			if matches[2] == "cm" {
				return num >= 150 && num <= 193
			} else if matches[2] == "in" {
				return num >= 59 && num <= 76
			} else { // should never happen
				return false
			}
		},
		"hcl": func(val string) bool {
			re := regexp.MustCompile(`^#[a-f0-9]{6}$`)
			return re.MatchString(val)
		},
		"ecl": func(val string) bool {
			validColors := map[string]struct{}{
				"amb": struct{}{},
				"blu": struct{}{},
				"brn": struct{}{},
				"gry": struct{}{},
				"grn": struct{}{},
				"hzl": struct{}{},
				"oth": struct{}{},
			}
			_, ok := validColors[val]
			return ok
		},
		"pid": func(val string) bool {
			re := regexp.MustCompile(`^[0-9]{9}$`)
			return re.MatchString(val)
		},
	}
	for _, field := range strings.Split(line, " ") {
		subfields := strings.Split(field, ":")
		key := subfields[0]
		val := subfields[1]
		if validator, ok := required[key]; ok {
			if !validator(val) {
				return false
			}
		}
	}
	return true
}

func validate(line string) bool {
	required := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	for _, seek := range required {
		foundSeek := false
		for _, field := range strings.Split(line, " ") {
			if field[:3] == seek {
				foundSeek = true
				break
			}
		}
		if !foundSeek {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var parsedLines []string
	curLine := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsedLines = append(parsedLines, curLine)
			curLine = ""
			continue
		}
		if curLine == "" {
			curLine = line
		} else {
			curLine = curLine + " " + line
		}
	}
	parsedLines = append(parsedLines, curLine) // get last line

	valid := 0
	fullyValid := 0
	for _, line := range parsedLines {
		if validate(line) {
			valid++
			if fullyValidate(line) {
				fullyValid++
			}
		}
	}
	fmt.Println(valid)
	fmt.Println(fullyValid)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
