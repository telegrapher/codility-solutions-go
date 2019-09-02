package tape_equilibrium

//import "fmt"
/*
* Each value is calculated as:
* Modulo of 2 times the left part - Total sum of array
* Left - Right = Left - (Total - Left) = 2Left - Total
 */
func Solution(array []int) int {
	sumElem := 0
	for _, elem := range array {
		sumElem = sumElem + elem
	}

	//fmt.Printf("Total is %d\n", sumElem)

	minDiff := 2147483647
	accumulated := 0
	for i := 0; i < len(array)-1; i++ {
		accumulated = accumulated + array[i]
		//fmt.Printf("Accumulated is %d\n", accumulated)
		diff := (2 * accumulated) - sumElem
		//fmt.Printf("Diff is %d\n", diff)
		if diff < 0 {
			diff = diff * (-1)
		}
		if diff < minDiff {
			minDiff = diff
			//fmt.Printf("MinDiff set to  %d\n", minDiff)
		}
	}
	return minDiff
}
