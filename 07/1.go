package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Part1(lines []string) int {

	start := time.Now()

	answ := 1

	times := strings.Fields(lines[0])[1:]
	records := strings.Fields(lines[1])[1:]

	for i := 0; i < len(times); i++ {
		// fmt.Printf("Time: %s, Record: %s\n", times[i], records[i])
		time, _ := strconv.Atoi(times[i])
		record, _ := strconv.Atoi(records[i])
		wins := 0

		for j := 0; j <= time; j++ {
			// fmt.Printf("%d * %d: %d (%t)\n", j, time-j, j*(time-j), j*(time-j) > record)
			if j*(time-j) > record {
				wins++
			}

		}
		// fmt.Println("Wins:", wins)

		answ *= wins

	}

	end := time.Now()

	fmt.Printf("Answer for part one: %d\n", answ)
	fmt.Printf("Execution time: %d\n", end.Sub(start))
	return 1

}
