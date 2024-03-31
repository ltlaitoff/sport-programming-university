package lb1

import (
	"math/big"
	"testing"
)

func TestTask8(t *testing.T) {
	tests := []struct {
		name           string
		number         *big.Int
		power          *big.Int
		expectedResult int64
	}{
		{
			// On run with default = 0.72s
			// On run with Binary Exponentiation = 0.01s
			name:   "Test 1",
			number: big.NewInt(100),
			power:  big.NewInt(100_000),
		},
		{
			// On run with default = 10.40s
			// On run with Binary Exponentiation = 0.03s
			name:   "Test 2",
			number: big.NewInt(3),
			power:  big.NewInt(1_000_000),
		},
		{
			// On run with default = bigger than 462.099s
			// On run with Binary Exponentiation = 0.44s
			name:   "Test 2 with big value",
			number: big.NewInt(2),
			power:  big.NewInt(10_000_000),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Task8(test.number, test.power)

			expectedNumber := big.NewInt(0).Exp(test.number, test.power, big.NewInt(0))

			if result.Cmp(expectedNumber) != 0 {
				t.Errorf("Task8 %v failed", test.name)
			}
		})
	}
}
