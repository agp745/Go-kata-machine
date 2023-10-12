package twocrystalballs

import (
	"testing"
)

func TestTwoCrystallBalls(t *testing.T) {

	tests := []struct {
		Arr            []bool
		ExpextedResult int
	}{
		{[]bool{false, false, true, true, true}, 2},
		{[]bool{false, false, false, true, true}, 3},
		{[]bool{false, false, false, false, true}, 4},
		{[]bool{false, false, false, false, false}, -1},
		{[]bool{true, true, true, true, true}, 0},
		{[]bool{false, true, true, true, true}, 1},
	}

	for i, test := range tests {
		result := two_crystal_balls(test.Arr)

		if result != test.ExpextedResult {
			t.Fatalf("tests[%d] - expected: %v, got: %v", i, test.ExpextedResult, result)
		}
	}
}
