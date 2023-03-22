package day1_test

import (
	"day1"
	"testing"
)

func TestCalculateCalories(t *testing.T) {
	tests := []struct {
		name, filename      string
		expectedValue, part int
	}{
		{"Part 1 - Find the elf with the most calories using example from site", "simple-example.txt", 24000, 1},
		{"Part 2 - Add the top three elf calorie counts using example from site", "simple-example.txt", 45000, 2},
		{"Part 1 - Find the elf with the most calories", "example.txt", 70374, 1},
		{"Part 2 - Add the top three elf calorie counts", "example.txt", 204610, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := day1.CalculateCalories(test.filename, test.part)

			if result != test.expectedValue {
				t.Errorf("Expected %d, got: %d", test.expectedValue, result)
			}
		})
	}
}
