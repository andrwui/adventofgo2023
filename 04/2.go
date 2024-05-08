package main

import (
	"adventofgo/helper"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Part2(lines []string) int {

	start := time.Now()

	qtties := make([]int, len(lines))
	for i := range qtties {
		qtties[i] = 1
	}

	for i := 0; i < len(lines); i++ {

		numbers := strings.TrimSpace(strings.Split(lines[i], ":")[1])

		winningNumbers := strings.Split(strings.TrimSpace(strings.Split(numbers, "|")[0]), " ")
		playerNumbers := strings.Split(strings.TrimSpace(strings.Split(numbers, "|")[1]), " ")

		matches := 0

		for _, wn := range winningNumbers {
			for _, pn := range playerNumbers {

				if pn != "" && wn != "" {

					intWn, err := strconv.Atoi(wn)
					helper.Check(err)
					intPn, err := strconv.Atoi(pn)
					helper.Check(err)

					if intWn == intPn {
						matches++
					}

				}
			}
		}

		for e := 0; e < matches; e++ {
			for j := 0; j < qtties[i]; j++ {
				qtties[i+(e+1)]++
			}
		}

	}
	end := time.Now()

	fmt.Printf("Answer for part two: %d\n", sum(qtties))
	fmt.Printf("Execution time: %d\n", end.Sub(start))
	return 1

}

func sum(nums []int) int {
	r := 0
	for _, num := range nums {
		r += num
	}
	return r
}
