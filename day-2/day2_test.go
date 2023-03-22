package day_2

import (
	"testing"
)

func TestRockPaperScissors(t *testing.T) {
	tests := []struct {
		name, filename      string
		expectedValue, part int
	}{
		{"Part 1 - Calculate score from shape choice and score when X,Y,Z represent Rock,Paper,Scissors example from site", "simple-example.txt", 15, 1},
		{"Part 2 - Calculate score from shape choice and score when X,Y,Z represent Lose,Draw,Win example from site", "simple-example.txt", 12, 2},
		{"Part 1 - Calculate score from shape choice and score when X,Y,Z represent Lose,Draw,Win", "example.txt", 14069, 1},
		{"Part 2 - Calculate score from shape choice and score when X,Y,Z represent Lose,Draw,Win", "example.txt", 12411, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RockPaperScissors(test.filename, test.part)

			if result != test.expectedValue {
				t.Errorf("Expected %d, got: %d", test.expectedValue, result)
			}
		})
	}
}
