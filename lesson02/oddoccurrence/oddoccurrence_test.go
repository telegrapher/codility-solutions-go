package oddoccurrence_test

import (
	"fmt"
	. "github.com/telegrapher/codility-solutions-go/lesson02/oddoccurrence"
	"testing"
)

var oddOccurrenceTests = []struct {
	input  []int
	result int
}{
	{[]int{1, 3, 5, 7, 5, 3, 1}, 7},
	{[]int{1}, 1},
	{[]int{1, 1, 1}, 1},
	{[]int{1, 3, 3}, 1},
	{[]int{1, 3, 3, 3, 3}, 1},
	{[]int{3, 3, 3, 3, 1}, 1},
}

func TestParser(t *testing.T) {

	for _, currentTest := range oddOccurrenceTests {
		t.Run(fmt.Sprintf("Test %d", currentTest.input), func(t *testing.T) {
			result := Solution(currentTest.input)
			if result != currentTest.result {
				t.Fatalf("Solution for %d returned %d, expected %d.\n",
					currentTest.input, result, currentTest.result)
			}
		})
	}
}
