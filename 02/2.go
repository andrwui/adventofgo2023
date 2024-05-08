package main

import (
	"adventofgo/helper"
	"fmt"
	"strconv"
	"strings"
)

func Part2(lines []string) int {

	sum := 0

	for _, line := range lines {

		max := map[string]int{}

		game := strings.Split(line, ":")[1]

		for _, roll := range strings.Split(game, ";") {

			roll = strings.TrimSpace(roll)

			for _, color := range strings.Split(roll, ", ") {

				qty, err := strconv.Atoi(strings.TrimSpace(strings.Split(color, " ")[0]))
				helper.Check(err)

				curColor := strings.Split(color, " ")[1]

				if qty > max[curColor] {
					max[curColor] = qty
				}

				if max[curColor] == 0 {
					max[curColor] = qty
				}

			}

		}

		sum += max["red"] * max["green"] * max["blue"]

	}

	fmt.Printf("Answer for part two: %d", sum)
	return 1
}
