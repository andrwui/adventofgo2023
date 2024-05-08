package main

import (
	"adventofgo/helper"
	"fmt"
	"os"
)

func main() {

	path, err := os.Getwd()
	helper.Check(err)

	lines := helper.GetFileLines(path + "/input.txt")

	Part1(lines)
	fmt.Println("")
	Part2(lines)

}
