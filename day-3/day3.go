package day_3

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var priorityMap = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func ReorganiseRucksack(filename string, part int) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf(err.Error(), "err")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	var score = 0
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		compartmentOne, compartmentTwo := scanner.Text()[:(len(scanner.Text())/2)], scanner.Text()[(len(scanner.Text())/2):]

		for i := 0; i < len(compartmentOne); i++ {
			if strings.Contains(compartmentTwo, string(compartmentOne[i])) {
				score += priorityMap[string(compartmentOne[i])]
				break
			}
		}
	}

	if part == 1 {
		return score
	}

	score = 0
	for i := 0; i < len(lines); i += 3 {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			if strings.Contains(lines[i+1], string(line[j])) && strings.Contains(lines[i+2], string(line[j])) {
				score += priorityMap[string(line[j])]
				break
			}
		}
	}
	return score
}
