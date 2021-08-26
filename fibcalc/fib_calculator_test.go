package fibcalc

import (
	"math/big"
	"testing"
)

func TestComputeFibonacci(t *testing.T) {
	cases := []struct {
		input    *big.Int
		expected *big.Int
	}{
		{
			input:    big.NewInt(0),
			expected: big.NewInt(0),
		},
		{
			input:    big.NewInt(1),
			expected: big.NewInt(1),
		},
		{
			input:    big.NewInt(8),
			expected: big.NewInt(21),
		},
	}

	for _, testCase := range cases {
		actual := NewFibonacciCalculator().GetTerm(testCase.input)
		if actual.Cmp(testCase.expected) != 0 {
			t.Errorf("Error for input (%s): expected %s, actual %s", testCase.input, testCase.expected, actual)
		}
	}
}
