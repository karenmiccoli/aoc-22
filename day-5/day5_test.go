package day_5

import (
	"testing"
)

func TestReorganiseStacks(t *testing.T) {
	tests := []struct {
		name, filename, expectedValue string
		part                          int
	}{
		{"Part 1 - Reorganise stacks finding first in each stack example from site", "simple-example.txt", "CMZ", 1},
		{"Part 2 - Reorganise stacks in different order example from site", "simple-example.txt", "MCD", 2},
		{"Part 1 - Reorganise stacks finding first in each stack", "example.txt", "TBVFVDZPN", 1},
		{"Part 2 - Reorganise stacks in different order", "example.txt", "VLCWHTDSZ", 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RucksackReorganisation(test.filename, test.part)

			if result != test.expectedValue {
				t.Errorf("Expected %s, got: %s", test.expectedValue, result)
			}
		})
	}
}
