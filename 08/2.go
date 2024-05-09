package main

import (
	"fmt"
	"time"
)

func Part2(lines []string) int {
	start := time.Now()

	answ := 0

	end := time.Now()

	fmt.Printf("Answer for part one: %d\n", answ)
	fmt.Printf("Execution time: %d\n", end.Sub(start))
	return 1

}
