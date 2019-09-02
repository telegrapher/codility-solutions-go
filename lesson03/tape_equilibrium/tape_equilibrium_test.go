package tape_equilibrium

import "testing"
import "fmt"

var tapeEquilibriumTests = []struct {
	array  []int
	result int
}{
	{[]int{3, 1, 2, 4, 3}, 1},
	{[]int{3, -3, -10}, 10},
	{[]int{-3, -3, -10}, 4},
	{[]int{0, 0, 0}, 0},
	{[]int{1000, 1000, 0}, 0},
}

func TestParser(t *testing.T) {
	for _, currentTest := range tapeEquilibriumTests {
		t.Run(fmt.Sprintf("Test %v", currentTest.array), func(t *testing.T) {
			testResult := Solution(currentTest.array)
			if currentTest.result != testResult {
				t.Fatalf("Test returned %d, expected %d\n", testResult, currentTest.result)
			}
		})
	}
}
