package main

import (
	"adventofgo/helper"
	"fmt"
	"strconv"
	"unicode"
)

func Part1(lines []string) int {

	sum := 0

	for _, line := range lines {
		var lineAcummulator []string

		for _, char := range line {

			if unicode.IsDigit(char) {
				lineAcummulator = append(lineAcummulator, string(char))

			}
		}

		strResult, err := strconv.Atoi(lineAcummulator[0] + lineAcummulator[len(lineAcummulator)-1])
		helper.Check(err)

		sum += strResult

	}
	fmt.Println("Answer for part one:", sum)
	return sum
}
