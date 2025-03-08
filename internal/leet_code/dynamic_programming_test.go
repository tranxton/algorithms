package leet_code

import (
	"testing"
	"time"
)

func TestMaxProfit(t *testing.T) {
	maxProfit := MaxProfit([]int{7, 1, 5, 3, 6, 4})
	expectedMaxProfit := 5

	if maxProfit != expectedMaxProfit {
		t.Errorf("Expected max profit to be %d, got %d", expectedMaxProfit, maxProfit)
	}
}

func TestFibonacci(t *testing.T) {
	var expectedResult uint64 = 17323038258947941269
	expectedDurationLimit := 25 * time.Microsecond

	timeStart := time.Now()
	result := Fibonacci(200)
	duration := time.Since(timeStart)

	if duration > expectedDurationLimit {
		t.Errorf("Duration took too long: %v, expected less than %v", duration, expectedDurationLimit)
	}

	if result != expectedResult {
		t.Errorf("Expected result to be %d, got %d", expectedResult, result)
	}
}

func TestFindLongestCommonSubstring(t *testing.T) {
	target := "common"
	words := []string{"test", "comrade", "comma", "comment", "commode", "precommon"}
	expectedResult := "precommon"
	result := FindLongestCommonSubstring(target, words)

	if result != expectedResult {
		t.Errorf("Expected result to be %s, got %s", expectedResult, result)
	}
}
