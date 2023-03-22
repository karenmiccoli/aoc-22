package day_4

import "testing"

func TestCleanupCalculator(t *testing.T) {
	tests := []struct {
		name, filename      string
		expectedValue, part int
	}{
		{"Part 1 - Number of completely overlapping pairs example from site", "simple-example.txt", 2, 1},
		{"Part 2 - Number of partially overlapping pairs example from site", "simple-example.txt", 4, 2},
		{"Part 1 - Number of completely overlapping pairs", "example.txt", 500, 1},
		{"Part 2 - Number of partially overlapping pairs", "example.txt", 815, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CleanupCalculator(test.filename, test.part)

			if result != test.expectedValue {
				t.Errorf("Expected %d, got: %d", test.expectedValue, result)
			}
		})
	}
}
