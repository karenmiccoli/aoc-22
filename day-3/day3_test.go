package day_3

import "testing"

func TestReorganiseRucksack(t *testing.T) {
	tests := []struct {
		name, filename      string
		expectedValue, part int
	}{
		{"Part 1 - Reorganise rucksack finding common char & adding score example from site", "simple-example.txt", 157, 1},
		{"Part 2 - Reorganise rucksack finding common char in the three lines, adding score example from site", "simple-example.txt", 70, 2},
		{"Part 1 - Reorganise rucksack finding common char & adding score", "example.txt", 8139, 1},
		{"Part 2 - Reorganise rucksack finding common char in the three lines, adding score", "example.txt", 2668, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ReorganiseRucksack(test.filename, test.part)

			if result != test.expectedValue {
				t.Errorf("Expected %d, got: %d", test.expectedValue, result)
			}
		})
	}
}
