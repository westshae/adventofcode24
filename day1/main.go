package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	leftSide, rightSide, err := convertScannerToTwoLists()
	if err != nil {
		return
	}
	fmt.Println(leftSide.Len())
	fmt.Println(rightSide.Len())

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

func convertScannerToTwoLists() (list.List, list.List, error) {
	leftSide := list.New()
	rightSide := list.New()
	file, scanner, err := loadFile("input.txt")

	if err != nil {
		return *leftSide, *rightSide, err
	}

	defer file.Close()

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		leftSide.PushBack(words[0])
		rightSide.PushBack(words[1])
	}
	return *leftSide, *rightSide, nil
}
