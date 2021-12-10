package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/naethiel/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputLines()

	var oxygen, co2 string

	oxygen = filter(input, func(value byte, ones int, zeroes int) bool {
		if ones >= zeroes {
			return value == '1'
		} else {
			return value == '0'
		}
	}, 0)

	co2 = filter(input, func(value byte, ones int, zeroes int) bool {
		if zeroes <= ones {
			return value == '0'
		} else {
			return value == '1'
		}
	}, 0)

	o, err := strconv.ParseInt(oxygen, 2, 64)
	if err != nil {
		log.Fatal("failing to convert oxygen to decimal", err)
	}

	c, err := strconv.ParseInt(co2, 2, 64)
	if err != nil {
		log.Fatal("failing to convert co2 to decimal", err)
	}

	fmt.Println("result: ", c*o)
}

func filter(list []string, pred func(val byte, ones int, zeroes int) bool, index int) string {
	var newList []string

	if len(list) == 1 {
		return list[0]
	}
	var zeroes, ones int = 0, 0

	// let's count first
	for _, line := range list {
		if line[index] == '1' {
			ones++
		} else {
			zeroes++
		}
	}

	// then filter
	for _, line := range list {
		if pred(line[index], ones, zeroes) {
			newList = append(newList, line)
		}
	}

	return filter(newList, pred, index+1)
}
