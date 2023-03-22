package day_10

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func SignalStrength(filename string, cyclesToCalc []int) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf(err.Error(), "err")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var results = make(map[int]int)
	var X = 1
	var valueToAdd, cycleNum = 0, 0
	for scanner.Scan() {
		var numCycles = 1

		commandParts := strings.Split(scanner.Text(), " ")
		if commandParts[0] == "addx" {
			numCycles = 2
			valueToAdd = convertToInt(commandParts[1])
		}

		for i := 0; i < numCycles; i++ {
			cycleNum++
			if i > 0 {
				X += valueToAdd
			}
			results[cycleNum] = X
		}
	}

	var count = 0
	for _, cycle := range cyclesToCalc {
		count += cycle * results[cycle-1]
	}

	return count
}

func convertToInt(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Failed to convert to int")
	}
	return number
}
