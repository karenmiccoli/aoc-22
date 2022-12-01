package day1

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type kv struct {
	Key   int
	Value int
}

func CalculateCalories() []kv {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatalf(err.Error(), "err")
	}

	//Split on \n\n to get individual blocks of calories
	//Loop through those blocks, convert string to ints
	//Sum the slice and add results to map where key is the index and value is number of calories
	//Sort slice from the highest calorie value (descending order)

	results := make(map[int]int)

	individualCalories := strings.Split(string(data), "\n\n")

	for i, _ := range individualCalories {
		separatedCalories := convertSliceToInt(strings.Split(individualCalories[i], "\n"))
		result := sumSlice(separatedCalories)
		results[i] = result
	}

	return sortByDescendingValue(results)
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

//Takes a map and sorts in descending order returning a slice of structs separated with keys and values
func sortByDescendingValue(results map[int]int) []kv {
	var ss []kv
	for k, v := range results {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}

//FindTopThreeElves - After calculating the elf calories, the top three and summed and returned
func FindTopThreeElves() int {
	elfCalories := CalculateCalories()
	topThreeElves := elfCalories[:3]

	var result = 0
	for _, elf := range topThreeElves {
		result += elf.Value
	}

	return result
}
