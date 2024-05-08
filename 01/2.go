package main

import (
	"adventofgo/helper"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Part2(lines []string) int {

	numbers := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	sum := 0

	for _, line := range lines {
		var lineAcummulator []string

		for i, char := range line {

			if unicode.IsDigit(char) {
				lineAcummulator = append(lineAcummulator, string(char))
			}

			for key, value := range numbers {
				if strings.HasPrefix(line[i:], key) {
					lineAcummulator = append(lineAcummulator, value)
				}
			}

		}

		strResult, err := strconv.Atoi(lineAcummulator[0] + lineAcummulator[len(lineAcummulator)-1])
		helper.Check(err)

		sum += strResult

	}
	fmt.Println("Answer for part two:", sum)
	return sum

}
