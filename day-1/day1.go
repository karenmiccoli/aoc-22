package day1

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CalculateCalories(filename string, part int) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf(err.Error(), "err")
	}

	var results []int
	individualCalories := strings.Split(string(data), "\n\n")

	for i, _ := range individualCalories {
		separatedCalories := convertSliceToInt(strings.Split(individualCalories[i], "\n"))
		result := sumSlice(separatedCalories)
		results = append(results, result)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(results)))

	if part == 1 {
		return results[0]
	}
	return results[0] + results[1] + results[2]
}

//Takes a slice of strings and converts to slice of ints
func convertSliceToInt(stringSlice []string) []int {
	intSlice := make([]int, len(stringSlice))

	for i, stringNum := range stringSlice {
		intSlice[i], _ = strconv.Atoi(stringNum)
	}

	return intSlice
}

//Takes a slice of ints and gets sum of all values
func sumSlice(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
