package lb3

import (
	"reflect"
	"testing"
)

func TestTask8Lab3(t *testing.T) {
	tests := []struct {
		name    string
		stones  []int
		stones1 []int
		stones2 []int
		err     bool
	}{
		{
			name:    "Test 1",
			stones:  []int{10, 5, 5},
			stones1: []int{5, 5},
			stones2: []int{10},
			err:     false,
		},
		{
			name:    "Test 2",
			stones:  []int{10, 5, 5, 3, 1, 4},
			stones1: []int{5, 5, 4},
			stones2: []int{10, 3, 1},
			err:     false,
		},
		{
			name:    "Test 3",
			stones:  []int{25, 5, 1, 4},
			stones1: []int{},
			stones2: []int{},
			err:     true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stones1, stones2, err := Lab3Task8(test.stones)

      if !((len(stones1) == 0 && len(stones1) == 0) || reflect.DeepEqual(stones1, test.stones1)) ||
				!((len(stones2) == 0 && len(stones2) == 0) || reflect.DeepEqual(stones2, test.stones2)) ||
				err != test.err {
				t.Errorf("Lab3Task6 %v failed", test.name)
			}
		})
	}
}
