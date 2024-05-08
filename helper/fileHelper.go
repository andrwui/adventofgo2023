package helper

import (
	"bufio"
	"os"
)

func GetFileLines(path string) []string {

	var results []string

	file, e := os.Open(path)
	Check(e)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	return results

}
