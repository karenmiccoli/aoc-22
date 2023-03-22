package day_6

import (
	"log"
	"os"
)

func FindMarkers(filename string, marker int) int {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf(err.Error(), "err")
	}

	input := string(file)
	var result = 0

	for i := 0; i < len(input); i++ {
		charMap := make(map[byte]bool)
		for j := 0; j < marker; j++ {
			charMap[input[i+j]] = true
		}
		if len(charMap) == marker {
			result = i + marker
			break
		}
	}

	return result
}
