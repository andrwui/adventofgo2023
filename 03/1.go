package main

import (
	"adventofgo/helper"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Part1(lines []string) int {

	start := time.Now()

	answ := 0

	for _, card := range lines {

		numbers := strings.TrimSpace(strings.Split(card, ":")[1])

		winningNumbers := strings.Split(strings.TrimSpace(strings.Split(numbers, "|")[0]), " ")
		playerNumbers := strings.Split(strings.TrimSpace(strings.Split(numbers, "|")[1]), " ")

		cardPoints := 0

		for _, wn := range winningNumbers {
			for _, pn := range playerNumbers {
				if pn != "" && wn != "" {
					intWn, err := strconv.Atoi(wn)
					helper.Check(err)
					intPn, err := strconv.Atoi(pn)
					helper.Check(err)

					if intWn == intPn {
						if cardPoints < 1 {
							cardPoints++
						} else {
							cardPoints *= 2
						}

					}

				}

			}
		}

		answ += cardPoints

	}

	end := time.Now()

	fmt.Printf("Answer for part one: %d\n", answ)
	fmt.Printf("Execution time: %d\n", end.Sub(start))
	return 1

}
