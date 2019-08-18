package perm_missing_elem

import "testing"
import "fmt"

var permMissingElemTests = []struct {
	permutation []int
	missing     int
}{
	{[]int{2, 3, 1, 5}, 4},
	{[]int{2, 3, 1, 4}, 5},
	{[]int{2}, 1},
}

func TestParser(t *testing.T) {
	for _, currentTest := range permMissingElemTests {
		t.Run(fmt.Sprintf("Test %v", currentTest.permutation), func(t *testing.T) {
			testResult := Solution(currentTest.permutation)
			if currentTest.missing != testResult {
				t.Fatalf("Test returned %d, expected %d\n", testResult, currentTest.missing)
			}
		})
	}
}
