package main

import (
	"github.com/naethiel/advent-of-code-2021/utils"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadInputLines()

	var positionH int = 0
	var positionV int = 0
	var aim int = 0

	for _, line := range lines {
		splitted := strings.Split(line, " ")

		move := splitted[0]

		count, err := strconv.Atoi(splitted[1])
		if err != nil {
			log.Fatal("invalid move count", err)
		}

		switch move {
		case "forward":
			positionH += count
			positionV += aim * count
		case "up":
			aim -= count
		case "down":
			aim += count
		}

	}

	log.Printf("x: %d - depth: %d - total: %d", positionH, positionV, positionH*positionV)
}
