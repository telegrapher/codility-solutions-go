package cyclic_rotation

//import "fmt"

func RotateIndex(initial int, arrayLenght int, jump int) int {
	//fmt.Printf("Initial %d Jump %d Length %d. ", initial, jump, arrayLenght)
	//fmt.Printf("%d -> %d \n", (initial + jump), ((initial + jump) % arrayLenght))
	return ((initial + jump) % arrayLenght)
}

func Solution(A []int, K int) []int {

	// Start handling possible corner cases.
	// Empty array.
	if len(A) == 0 {
		return A
	}
	var rotPos int
	// If the rotation movements are bigger than the array size, avoid unnecesary movements.
	if K >= len(A) {
		rotPos = K % len(A)
	} else {
		rotPos = K
	}

	B := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		//fmt.Printf("The i=%d and B position is %d\n", i, RotateIndex(i, len(A), rotPos))
		B[RotateIndex(i, len(A), rotPos)] = A[i]
	}
	return B
}
