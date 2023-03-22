package day_6

import "testing"

func TestFindMarkers(t *testing.T) {
	tests := []struct {
		name, filename        string
		expectedValue, marker int
	}{
		{"Part 1 - mjqjpqmgbljsphdztnvjfqwrcgsmlb with marker as 4 simple example", "simple-example-1.txt", 7, 4},
		{"Part 1 - bvwbjplbgvbhsrlpgdmjqwftvncz with marker as 4 simple example", "simple-example-2.txt", 5, 4},
		{"Part 1 - nppdvjthqldpwncqszvftbrmjlhg with marker as 4 simple example", "simple-example-3.txt", 6, 4},
		{"Part 1 - nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg with marker as 4 simple example", "simple-example-4.txt", 10, 4},
		{"Part 1 - zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw with marker as 4 simple example", "simple-example-5.txt", 11, 4},
		{"Part 1 - Puzzle input with marker as 4", "example.txt", 1356, 4},
		{"Part 2 - mjqjpqmgbljsphdztnvjfqwrcgsmlb with marker as 14 simple example", "simple-example-1.txt", 19, 14},
		{"Part 2 - bvwbjplbgvbhsrlpgdmjqwftvncz with marker as 14 simple example", "simple-example-2.txt", 23, 14},
		{"Part 2 - nppdvjthqldpwncqszvftbrmjlhg with marker as 14 simple example", "simple-example-3.txt", 23, 14},
		{"Part 2 - nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg with marker as 14 simple example", "simple-example-4.txt", 29, 14},
		{"Part 2 - zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw with marker as 14 simple example", "simple-example-5.txt", 26, 14},
		{"Part 2 - Puzzle input with marker as 14", "example.txt", 2564, 14},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FindMarkers(test.filename, test.marker)

			if result != test.expectedValue {
				t.Errorf("Expected %d, got: %d", test.expectedValue, result)
			}
		})
	}
}
