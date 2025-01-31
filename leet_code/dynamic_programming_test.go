package leet_code

import "testing"

func TestMaxProfit(t *testing.T) {
	maxProfit := MaxProfit([]int{7, 1, 5, 3, 6, 4})
	expectedMaxProfit := 5

	if maxProfit != expectedMaxProfit {
		t.Errorf("Expected max profit to be %d, got %d", expectedMaxProfit, maxProfit)
	}
}
