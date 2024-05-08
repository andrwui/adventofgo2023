package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	clearScreen()

	scanner := bufio.NewScanner(os.Stdin)

	dirs, err := listDirectories(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, dir := range dirs {
		fmt.Printf("%d. %s\n", i+1, dir)
	}

	fmt.Print("Enter directory number to run: ")
	scanner.Scan()
	dirNumber := strings.TrimSpace(scanner.Text())

	fmt.Println("")

	selectedDirIndex := 0
	fmt.Sscan(dirNumber, &selectedDirIndex)

	if selectedDirIndex < 1 || selectedDirIndex > len(dirs) {
		fmt.Println("Invalid directory number.")
		return
	}

	selectedDir := dirs[selectedDirIndex-1]

	fmt.Println("Running directory " + selectedDir + "...")

	fmt.Println("")

	cmd := exec.Command("go", "run", ".")
	cmd.Dir = selectedDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func listDirectories(path string) ([]string, error) {
	var dirs []string

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			if file.Name() != "helper" {
				dirs = append(dirs, file.Name())
			}
		}
	}

	return dirs, nil
}

func clearScreen() {
	clear := exec.Command("cmd", "/c", "cls")
	clear.Stdout = os.Stdout
	clear.Run()
}
