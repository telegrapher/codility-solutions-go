package frogjmp_test

import "testing"
import "fmt"
import . "../frogjmp"

var frogjmpTests = []struct {
	X, Y, Z, result int
}{
	{10, 85, 30, 3},
	{10, 85, 25, 3},
	{10, 86, 25, 4},
	{0, 76, 25, 4},
	{0, 75, 75, 1},
	{0, 75, 76, 1},
	{0, 0, 100, 0},
}

func TestParser(t *testing.T) {
	for _, currentTest := range frogjmpTests {
		t.Run(fmt.Sprintf("Test %v", currentTest), func(t *testing.T) {
			testResult := Solution(currentTest.X, currentTest.Y, currentTest.Z)
			if currentTest.result != testResult {
				t.Fatalf("test %v returned %d\n", currentTest, testResult)
			}
		})
	}
}
