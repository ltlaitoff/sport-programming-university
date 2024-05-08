package lb2

import (
	"fmt"
	"testing"
)

func TestTask7Lab2(t *testing.T) {
	tests := []struct {
		name           string;
		base         int 
		expectedResult string
	}{
		{
			name:   "Test 1",
      base: 10,
      expectedResult: "4037914",
		},
		{
			name:   "Test 2",
      base: 12,
      expectedResult: "522956314",
		},
		{
			name: "Test 3",
      base: 100,
      // Just believe
      expectedResult: "94269001683709979260859834124473539872070722613982672442938359305624678223479506023400294093599136466986609124347432647622826870038220556442336528920420940314",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Lab2Task7(test.base)
  
      fmt.Println(result)
			if (result != test.expectedResult) {
				t.Errorf("Lab2Task5 %v failed", test.name)
			}
		})
	}
}
