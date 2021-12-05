package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/naethiel/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputLines()
	var gamma, epsilon string

	// Loop on input columns
	for i := 0; i < len(input[0]); i++ {
		zeros := 0
		ones := 0

		// loop through all lines
		for j := 0; j < len(input); j++ {
			// count ones and zeros
			if input[j][i] == '0' {
				zeros++
			} else {
				ones++
			}
		}

		// if zeros are larger gamma is 0 epsilon is 1.
		if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			// if ones are larger, gamma is 1 epsilon is 0.
			gamma += "1"
			epsilon += "0"
		}
	}

	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal("failing to convert gamma to decimal", err)
	}

	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatal("failing to convert epsilon to decimal", err)
	}

	fmt.Println("result: ", g*e)
}
