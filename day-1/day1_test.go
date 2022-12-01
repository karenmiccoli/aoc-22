package day1_test

import (
	"day1"
	"testing"
)

func TestCalculateCalories(t *testing.T) {
	t.Run("Part 1 - Find the elf with the most calories using example from site", func(t *testing.T) {
		expectedValue := 24000
		caloriesCalculations := day1.CalculateCalories(true)
		if caloriesCalculations[0].Value != expectedValue {
			t.Errorf("Expected %d, got: %d", expectedValue, caloriesCalculations[0].Value)
		}
	})

	t.Run("Part 2 - Add the top three elf calorie counts using example from site", func(t *testing.T) {
		expectedValue := 45000
		result := day1.FindTopThreeElves(true)
		if result != expectedValue {
			t.Errorf("Expected %d, got: %d", expectedValue, result)
		}
	})

	t.Run("Part 1 - Find the elf with the most calories", func(t *testing.T) {
		expectedValue := 70374
		caloriesCalculations := day1.CalculateCalories(false)
		if caloriesCalculations[0].Value != expectedValue {
			t.Errorf("Expected %d, got: %d", expectedValue, caloriesCalculations[0].Value)
		}
	})

	t.Run("Part 2 - Add the top three elf calorie counts", func(t *testing.T) {
		expectedValue := 204610
		result := day1.FindTopThreeElves(false)
		if result != expectedValue {
			t.Errorf("Expected %d, got: %d", expectedValue, result)
		}
	})
}
