package day1_test

import (
	"day1"
	"testing"
)

func TestCalculateCalories(t *testing.T) {
	t.Run("Part 1 - Find the elf with the most calories", func(t *testing.T) {
		expectedValue := 70374
		caloriesCalculations := day1.CalculateCalories()
		if caloriesCalculations[0].Value != expectedValue {
			t.Errorf("Expected %d, got: %d", expectedValue, caloriesCalculations[0].Value)
		}
	})
	t.Run("Part 2 - Add the top three elf calorie counts", func(t *testing.T) {
		expectedValue := 204610
		result := day1.FindTopThreeElves()
		if result != expectedValue {
			t.Errorf("Expected %d, got: %d", expectedValue, result)
		}
	})
}
