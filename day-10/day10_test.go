package day_10

import "testing"

func TestSignalStrength(t *testing.T) {
	tests := []struct {
		name, filename string
		expectedValue  int
		cycles         []int
	}{
		{"Part 1 - Find signal strength using simple example", "simple-example.txt", 13140, []int{20, 60, 100, 140, 180, 220}},
		{"Part 1 - Find signal strength", "example.txt", 14520, []int{20, 60, 100, 140, 180, 220}},
		//{"Part 2 - Find smallest directory to delete simple example", "simple-example.txt", 24933642, 2},
		//{"Part 2 - Find smallest directory to delete", "example.txt", 8679207, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SignalStrength(test.filename, test.cycles)

			if result != test.expectedValue {
				t.Errorf("Expected %d, got: %d", test.expectedValue, result)
			}
		})
	}
}
