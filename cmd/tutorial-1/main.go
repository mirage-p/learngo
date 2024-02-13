package main

import (
	"fmt"
	"time"
)

func main() {
	var n int = 1000000
	var notAllocatedSlice = []int{}
	var allocatedSlice = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(notAllocatedSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(allocatedSlice, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var start = time.Now()

	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(start)
}
