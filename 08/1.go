package main

import (
	"fmt"
	"strings"
	"time"
)

func Part1(lines []string) int {

	start := time.Now()

	answ := 0

	instructions := lines[0]
	nodes := lines[2:]
	nextNode := "AAA"

	nodeMap := make(map[string][]string)
	for _, node := range nodes {

		currentNode := strings.TrimSpace(strings.Split(node, "=")[0])
		nodesString := strings.Split(node, "=")[1]
		nodesString = strings.ReplaceAll(nodesString, "(", "")
		nodesString = strings.ReplaceAll(nodesString, ")", "")
		possibleNodes := strings.Fields(strings.Join(strings.Split(nodesString, ","), " "))

		nodeMap[currentNode] = possibleNodes

	}

InfiniteLoop:
	for {
		for _, direction := range instructions {

			if nextNode == "ZZZ" {
				break InfiniteLoop
			}
			if direction == 'L' {
				nextNode = nodeMap[nextNode][0]
				answ++
			} else {
				nextNode = nodeMap[nextNode][1]
				answ++
			}

		}

	}

	end := time.Now()

	fmt.Printf("Answer for part one: %d\n", answ)
	fmt.Printf("Execution time: %d\n", end.Sub(start))
	return 1

}
