package string_split

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := map[string]int {
		//"babaa":2,
		"ababa":4,
		"aba":0,
		//"bbbbb":6,
	}

	for k, v := range tests {
		testName := fmt.Sprintf("%v is expected to return %v", k, v)
		t.Run(testName, func(t *testing.T) {
			if s := Solution(k); s != v {
				t.Errorf("got %v, want %v", s, v)
			}
		})
	}
}