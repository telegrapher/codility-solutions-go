/*
 * Use maps (hash tables)
 * Hash table offers fast lookups and deletes
 * An implementation that would delete elements at the end of the array
 * was too slow
 */

package oddoccurrence

import "fmt"

func findFirstPair(val int, S []int) int {
	for index, element := range S {
		if val == element {
			fmt.Printf("Pair found Return %d\n", index)
			return index
		}
	}
	return -1
}

func Solution(A []int) int {

	var m map[int]int
	m = make(map[int]int)
	// Iterate over every element of the array
	for i := 0; i < len(A); i++ {
		//Check if the key A[i] exists
		_, ok := m[A[i]]
		if ok {
			//If the key exists, delete the key
			delete(m, A[i])
		} else {
			//If the key doesn't exist, create it
			m[A[i]] = 1
		}
	}
	// An odd element will have its key created but not deleted.
	// There should be only one
	if len(m) == 1 {
		for key, _ := range m {
			return key
		}
	}
	return -1
}
