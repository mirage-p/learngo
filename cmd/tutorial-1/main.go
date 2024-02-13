package main

import (
	"fmt"
	"time"
)

func main() {
	var n int = 1000000
	var notAllocatedSlice = []int{}
	var allocatedSlice = make([]int, 0, n)

	fmt.Printf("\nTotal time without preallocation: %v\n", timeLoop(notAllocatedSlice, n))
	fmt.Printf("\nTotal time with preallocation: %v\n", timeLoop(allocatedSlice, n))

	var arr1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location for arrary 1 is: %p", &arr1)
	var result [5]float64 = square(&arr1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value is: %v", arr1)
}

func square(float *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the second array is: %p", float)
	for i := range float {
		float[i] = float[i] * float[i]
	}
	return *float
}

func timeLoop(slice []int, n int) time.Duration {
	var start = time.Now()

	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(start)
}
