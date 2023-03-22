package day_4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func CleanupCalculator(filename string, part int) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf(err.Error(), "err")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var result = 0
	for scanner.Scan() {
		pairs := convertToIntBoundaries(scanner.Text())

		minA, minB, maxA, maxB := pairs[0][0], pairs[1][0], pairs[0][1], pairs[1][1]
		if part == 1 {
			if (minB >= minA && maxB <= maxA) || (minA >= minB && maxA <= maxB) {
				result++
			}
		} else {
			minPair, maxPair := pairs[0], pairs[1]
			if minB < maxA {
				minPair, maxPair = pairs[1], pairs[0]
			}
			if minPair[1] >= maxPair[0] {
				result += 1
			}
		}
	}

	return result
}

func convertToIntBoundaries(lines string) [][]int {
	pairs := strings.Split(lines, ",")
	result := make([][]int, 0)
	for _, pair := range pairs {
		p := strings.Split(pair, "-")
		pa, _ := strconv.Atoi(p[0])
		pb, _ := strconv.Atoi(p[1])
		result = append(result, []int{pa, pb})
	}

	return result
}
