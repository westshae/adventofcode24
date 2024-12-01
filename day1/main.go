package main

import (
	_ "container/list"
	"fmt"
	"os"
)

func main() {
	// leftSide := list.New()
	// rightSide := list.New()
	fmt.Println(*loadFile("./input.txt"))
	fmt.Println("Hello, World!")
}

func loadFile(path string) *string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	content := string(bytes)
	return &content
}
