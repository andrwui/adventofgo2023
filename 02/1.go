package main

import (
	"adventofgo/helper"
	"fmt"
	"strconv"
	"strings"
)

func Part1(lines []string) int {

	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0

	for _, line := range lines {

		id, err := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
		helper.Check(err)

		game := strings.Split(line, ":")[1]

		isPossible := true

		for _, roll := range strings.Split(game, ";") {

			if !isPossible {
				break
			}

			roll = strings.TrimSpace(roll)

			for _, color := range strings.Split(roll, ", ") {

				qty, err := strconv.Atoi(strings.TrimSpace(strings.Split(color, " ")[0]))
				helper.Check(err)

				curColor := strings.Split(color, " ")[1]
				limit := limits[curColor]

				if qty > limit {
					isPossible = false
				}
			}

		}

		if isPossible {
			sum += id
		}
	}

	fmt.Printf("Answer for part one: %d \n", sum)
	return 1
}
