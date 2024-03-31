package lb1

import "testing"

func TestTask9(t *testing.T) {
	tests := []struct {
		name           string
		x              []int
		y              []int
		expectedResult bool
	}{
		{
			name:           "Check with deletionsj",
			x:              []int{1, 2, 3},
			y:              []int{3, 2, 1, 4, 5},
			expectedResult: true,
		},
		{
			name:           "Check with same values, but changed order",
			x:              []int{1, 2, 3, 4, 5},
			y:              []int{3, 2, 1, 4, 5},
			expectedResult: true,
		},
		{
			name:           "Check with 1 not valid number",
			x:              []int{1, 2, 3},
			y:              []int{3, 2, 4, 5},
			expectedResult: false,
		},
		{
			name:           "Check with different numbers",
			x:              []int{1, 2, 3},
			y:              []int{4, 5, 6},
			expectedResult: false,
		},
		{
			name:           "Big big values",
			x:              bigSequenceNotRandomGenerate(1_000_000),
			y:              bigSequenceNotRandomGenerate(1_000_000),
			expectedResult: true,
		},
		{
			name:           "Big big values where y < x",
			x:              bigSequenceNotRandomGenerate(100_000_000),
			y:              bigSequenceNotRandomGenerate(99_000_000),
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualResult := Task9(test.x, test.y)

			if actualResult != test.expectedResult {
				t.Errorf("Task9 = %v, expected %v", actualResult, test.expectedResult)
			}
		})
	}
}

func bigSequenceNotRandomGenerate(length int) []int {
	sequence := make([]int, length)
	for i := 0; i < length; i++ {
		sequence[i] = i
	}

	return sequence
}
