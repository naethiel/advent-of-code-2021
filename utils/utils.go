package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func mustReadFile() *os.File {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ", os.Args[0], "<input file path>")
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal("reading input file", err)
	}

	return file
}

func ReadInputLines() []string {
	file := mustReadFile()

	defer file.Close()

	reader := bufio.NewScanner(file)
	var lines []string

	for reader.Scan() {
		lines = append(lines, reader.Text())
	}

	return lines
}

func ReadInputInts() []int64 {
	file := mustReadFile()

	defer file.Close()

	reader := bufio.NewScanner(file)
	var lines []int64

	for reader.Scan() {
		nbr, err := strconv.ParseInt(reader.Text(), 10, 64)

		if err != nil {
			log.Fatal("converting line to int", err)
		}

		lines = append(lines, nbr)
	}

	return lines
}
