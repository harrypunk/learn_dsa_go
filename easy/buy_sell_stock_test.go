package easy

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var prices = []int{7, 1, 5, 3, 6, 4}
	maxp := MaxProfit(prices)
	t.Log(maxp)
}
