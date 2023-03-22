package day_5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RucksackReorganisation(filename string, part int) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf(err.Error(), "err")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var stackLines []string
	var partOneStacks [][]string
	var partTwoStacks [][]string

	for scanner.Scan() {
		line := scanner.Text()
		if len(partOneStacks) > 0 {
			numOfCrates := 0
			originStack := 0
			targetStack := 0
			_, err = fmt.Sscanf(line, "move %d from %d to %d", &numOfCrates, &originStack, &targetStack)
			if err != nil {
				log.Fatalf("Can't parse line %v: %v", line, err)
			}
			partOneStacks = MoveCratesPartOne(partOneStacks, numOfCrates, originStack-1, targetStack-1)
			partTwoStacks = MoveCratesPartTwo(partTwoStacks, numOfCrates, originStack-1, targetStack-1)
		} else {
			if line == "" {
				partOneStacks = ParseCrates(stackLines[:len(stackLines)-1])
				partTwoStacks = ParseCrates(stackLines[:len(stackLines)-1])
			} else {
				stackLines = append(stackLines, line)
			}
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	resultOne := ""
	for _, stack := range partOneStacks {
		resultOne += stack[0]
	}

	if part == 1 {
		return resultOne
	}

	resultTwo := ""
	for _, stack := range partTwoStacks {
		resultTwo += stack[0]
	}

	return resultTwo
}

func ParseCrates(lines []string) [][]string {
	var stacks [][]string
	for _, line := range lines {
		for index := 0; index < len(line); index += 4 {
			row := index / 4
			fmt.Println(row, "ROW")
			if len(stacks) <= row {
				stacks = append(stacks, []string{})
			}
			if index+1 < len(line) {
				value := line[index+1]
				if value >= 'A' && value <= 'Z' {
					stacks[row] = append(stacks[row], string(value))
				}
			}
		}
	}
	return stacks
}

func MoveCratesPartOne(stacks [][]string, numOfCrates int, origin int, target int) [][]string {
	for i := 0; i < numOfCrates; i++ {
		//Get the top element
		topElement := stacks[origin][0]
		//Remove from origin stack
		stacks[origin] = stacks[origin][1:]
		//Add top element from origin to the target stack
		stacks[target] = append([]string{topElement}, stacks[target]...)
	}
	return stacks
}

// Crates moved stay in the same order
func MoveCratesPartTwo(stacks [][]string, numOfCrates int, origin int, target int) [][]string {
	//Get number of crates from origin stack
	topElements := append([]string{}, stacks[origin][:numOfCrates]...)
	//Remove from origin stack
	stacks[origin] = stacks[origin][numOfCrates:]
	//Add top elements from origin to the target stack
	stacks[target] = append(topElements, stacks[target]...)

	return stacks
}
