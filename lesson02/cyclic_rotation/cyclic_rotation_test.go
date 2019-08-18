package cyclic_rotation_test

import (
	"fmt"
	. "github.com/telegrapher/codility-solutions-go/lesson02/cyclic_rotation"
	//. "../cyclic_rotation"
	"testing"
)

var rotateIndexTests = []struct {
	initial     int
	arrayLength int
	jump        int
	solution    int
}{
	{0, 7, 1, 1},
	{1, 7, 1, 2},
}

var cyclicRotationTests = []struct {
	inputArray    []int
	rotation      int
	solutionArray []int
}{
	{
		[]int{1, 3, 5, 7, 5, 3, 1},
		1,
		[]int{1, 1, 3, 5, 7, 5, 3},
	},
	{
		[]int{1, 3, 5, 7, 5, 3, 1},
		0,
		[]int{1, 3, 5, 7, 5, 3, 1},
	},
	{
		[]int{1, 3, 5, 7, 5, 3, 1},
		7,
		[]int{1, 3, 5, 7, 5, 3, 1},
	},
	{
		[]int{1},
		7,
		[]int{1},
	},
	{
		nil,
		6,
		nil,
	},
	{
		[]int{1, 3, 5, 7, 5, 3, 1},
		6,
		[]int{3, 5, 7, 5, 3, 1, 1},
	},
}

func slicesEqual(first []int, second []int) bool {
	if len(first) != len(second) {
		return false
	}
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func TestParserRotateIndex(t *testing.T) {
	for _, currentTest := range rotateIndexTests {
		t.Run(
			fmt.Sprintf("Test RotateIndex %d %d %d",
				currentTest.initial, currentTest.arrayLength, currentTest.jump),
			func(t *testing.T) {
				result := RotateIndex(
					currentTest.initial,
					currentTest.arrayLength,
					currentTest.jump)
				if result != currentTest.solution {
					t.Fatalf("Results for RotateIndex(%d,%d,%d) returned %d, expected %d ",
						currentTest.initial,
						currentTest.arrayLength,
						currentTest.jump,
						result,
						currentTest.solution)
				}
			})
	}
}

func TestParserSolution(t *testing.T) {

	for _, currentTest := range cyclicRotationTests {
		t.Run(
			fmt.Sprintf("Test %v Rotated %d", currentTest.inputArray, currentTest.rotation),
			func(t *testing.T) {
				result := Solution(currentTest.inputArray, currentTest.rotation)
				if !slicesEqual(result, currentTest.solutionArray) {
					t.Fatalf("Solution for rotating:%v, %d positions\n"+
						"Returned: %v\n"+
						"Expected: %v\n",
						currentTest.inputArray, currentTest.rotation,
						result,
						currentTest.solutionArray)
				}
			})
	}
}
