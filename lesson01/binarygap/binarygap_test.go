package binarygap_test

import (
	"fmt"
	"math"
	"testing"

	. "github.com/telegrapher/codility-solutions-go/lesson01/binarygap"
)

var binaryGapTests = []struct {
	input  int
	result int
}{
	{int(math.Pow(2, 31) - 1), 0},  // 111...111
	{int(math.Pow(2, 30) + 1), 29}, // 100...001
	{int(math.Pow(2, 30)), 0},      // 100...000
	{int(math.Pow(2, 0)), 0},       // 000...001
	{5, 1},                         // 101
}

func TestParser(t *testing.T) {

	for _, currentTest := range binaryGapTests {
		t.Run(fmt.Sprintf("Test %d", currentTest.input), func(t *testing.T) {
			result := Solution(currentTest.input)
			if result != currentTest.result {
				t.Fatalf("Solution for %d returned %d, expected %d.\n",
					currentTest.input, result, currentTest.result)
			}
		})
	}
}
