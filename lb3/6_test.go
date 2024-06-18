package lb3

import (
	"reflect"
	"testing"
)

func TestTask6Lab2(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		c    []int
	}{
		{
			name: "Test 1",
			a:    []int{10, 5, 1},
			b:    []int{11, 9, 7, 5, 3, 2},
			c:    []int{11, 10, 9, 7, 5, 5, 3, 2, 1},
		},
		{
			name: "Test 2",
			a:    []int{25, 8, 1},
			b:    []int{23, 11, 9, 9, 3, 2},
			c:    []int{25, 23, 11, 9, 9, 8, 3, 2, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Lab3Task6(test.a, test.b)

			if !reflect.DeepEqual(result, test.c) {
				t.Errorf("Lab3Task6 %v failed", test.name)
			}
		})
	}
}
