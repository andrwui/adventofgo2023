package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Part2(lines []string) int {

	start := time.Now()

	end := time.Now()

	answ := 0

	times := strings.Fields(lines[0])[1:]
	records := strings.Fields(lines[1])[1:]

	bigTime, _ := strconv.Atoi(strings.Join(times, ""))
	bigRecord, _ := strconv.Atoi(strings.Join(records, ""))

	for j := 0; j <= bigTime; j++ {
		// fmt.Printf("%d * %d: %d (%t)\n", j, bigTime-j, j*(bigTime-j), j*(bigTime-j) > bigRecord)
		if j*(bigTime-j) > bigRecord {
			answ++
		}
	}

	fmt.Printf("Answer for part two: %d\n", answ)
	fmt.Printf("Execution time: %d\n", end.Sub(start))
	return 1

}
