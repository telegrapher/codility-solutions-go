package binarygap

import "fmt"

func findBiggestInSlice(slice []int) int {
	biggestFound := 0
	for _, element := range slice {
		if element > biggestFound {
			biggestFound = element
		}
	}
	return biggestFound
}

func biggestGap(binarySlice []int) int {

	// Check ints until it finds a 1 or it ends
	firstOne := 0
	for index, element := range binarySlice {
		if element == 1 {
			firstOne = index
			break
		}
	}

	var binaryGaps []int
	currentGap := 0
	// Continue evaluating until the end.
	for i := firstOne + 1; i < len(binarySlice); i++ {
		if binarySlice[i] == 0 {
			currentGap++
		} else {
			binaryGaps = append(binaryGaps, currentGap)
			currentGap = 0
		}
	}
	return findBiggestInSlice(binaryGaps)
}

func Solution(N int) int {
	//Create a slice of 32 bytes length and capacity
	binaryForm := make([]int, 32, 32)
	for i := 0; i < 31; i++ {
		fmt.Printf("Processing for %d\n", N)
		binaryForm[i] = N % 2
		N = N / 2
	}
	binaryForm[31] = 0
	fmt.Printf("The array is: \n %v\n", binaryForm)

	return biggestGap(binaryForm)
}
