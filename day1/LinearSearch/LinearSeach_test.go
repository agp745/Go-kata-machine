package linearsearch

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	foo := []int{2, 3, 7, 94, 230, 145, 745, 22}

	tests := []struct {
		Target         int
		ExpectedResult bool
	}{
		{2, true},
		{3, true},
		{94, true},
		{745, true},
		{10, false},
		{23, false},
		{22, true},
	}

	for i, test := range tests {
		result := linear_search(foo, test.Target)

		if result != tests[i].ExpectedResult {
			t.Fatalf("tests[%d] - expected: %v, got: %v", i, tests[i].ExpectedResult, result)
		}
	}
}
