package lb1

import "testing"

func TestTask1(t *testing.T) {
	tests := []struct {
		name           string
		input          []int64
		expectedResult int64
	}{
		{
			name:           "Test 1",
			input:          []int64{1, 2, 4, 8, 16, 32},
			expectedResult: 64,
		},
		{
			name:           "Test 2",
			input:          []int64{1, 3, 6, 10, 15},
			expectedResult: 2,
		},
		{
			name:           "Test 3",
			input:          []int64{2, 3, 5, 7, 11, 13, 17},
			expectedResult: 1,
		},
		{
			name:           "Test 4",
			input:          []int64{2, 4, 6, 8, 10},
			expectedResult: 1,
		},
		{
			name:           "Test 5",
			input:          []int64{1, 5, 10, 20, 40},
			expectedResult: 2,
		},
		{
			name:           "Test 6",
			input:          []int64{2, 3, 7, 14, 29},
			expectedResult: 1,
		},
		{
			name:           "Test 7",
			input:          []int64{1, 2, 3, 5, 8, 13, 21},
			expectedResult: 54,
		},
		{
			name:           "Test 8",
			input:          []int64{1, 2, 3, 6, 10, 20, 40},
			expectedResult: 83,
		},
		{
			name:           "Test 9",
			input:          []int64{2, 4, 8, 16, 32, 64, 128, 256},
			expectedResult: 1,
		},
		{
			name:           "Test 10",
			input:          []int64{1, 3, 5, 9, 13, 25, 49},
			expectedResult: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Task1(test.input)
			if result != test.expectedResult {
				t.Errorf("Test %s failed: expected %d, got %d", test.name, test.expectedResult, result)
			}
		})
	}
}
