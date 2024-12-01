package main

import (
	"bufio"
	_ "container/list"
	"fmt"
	"os"
)

func main() {
	// leftSide := list.New()
	// rightSide := list.New()
	file, scanner, err := loadFile("input.txt")

	if err != nil {
		return
	}

	defer file.Close()

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func loadFile(path string) (*os.File, *bufio.Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)
	return file, scanner, nil
}
