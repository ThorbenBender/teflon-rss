package main

import "testing"

func TestFunc(x, y int) int {
	return x + y
}

func TestScrapper(t *testing.T) {
	cases := []struct {
		Input    [2]int
		Expected int
	}{
		{
			Input:    [2]int{10, 2},
			Expected: 12,
		},
	}
	for _, cas := range cases {
		actual := TestFunc(cas.Input[0], cas.Input[1])
		if actual != cas.Expected {
			t.Errorf("Failed. %d is not equal %d", actual, cas.Expected)
		}
	}
}
