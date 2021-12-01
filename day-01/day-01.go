package main

import (
	"fmt"
	"github.com/naethiel/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputInts()

	var count = 0
	var prev int64

	for index, val := range input {
		// check that we have enough elements left in the list
		// to compute a sum of 3
		if index > len(input)-3 {
			continue
		}

		// compute sum of 3
		sum := val + input[index+1] + input[index+2]

		// on first iteration, skip since we have no previous value
		if index == 0 {
			prev = sum
			continue
		}

		// if new sum is greater than previous one, increment
		if sum > prev {
			count++
		}

		// store current sum as previous
		prev = sum
	}

	fmt.Printf("result: %d", count)
}
