package main

import (
	"compass.com/go-homework/comment_analyzer"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <directory>")
		return
	}

	dir := os.Args[1]

	statsMap, err := comment_analyzer.ProcessDirectory(dir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	comment_analyzer.PrintStats(statsMap)
}
