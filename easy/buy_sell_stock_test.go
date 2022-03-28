package easy

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var prices = []int{7, 1, 5, 3, 6, 4}
	maxp := MaxProfit(prices)
	if maxp != 7 {
		t.Fail()
	}
}

func TestMaxProfitOnePass(t *testing.T) {
	var prices = []int{7, 1, 5, 3, 6, 4}
	maxp := MaxProfitOnePass(prices)
	if maxp != 5 {
		t.Fail()
	}
}
