package day_7

import "testing"

func TestFindMarkers(t *testing.T) {
	tests := []struct {
		name, filename      string
		expectedValue, part int
	}{
		{"Part 1 - Find total size under 100000 using simple example", "simple-example.txt", 95437, 1},
		{"Part 1 - Find total size under 100000", "example.txt", 1449447, 1},
		{"Part 2 - Find smallest directory to delete simple example", "simple-example.txt", 24933642, 2},
		{"Part 2 - Find smallest directory to delete", "example.txt", 8679207, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FindFileSizeUnder100000(test.filename, test.part)

			if result != test.expectedValue {
				t.Errorf("Expected %d, got: %d", test.expectedValue, result)
			}
		})
	}
}
