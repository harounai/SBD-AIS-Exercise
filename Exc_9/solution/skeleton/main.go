package main

import (
	"exc9/mapred"
	"fmt"
	"os"
	"strings"
)

// Main function
func main() {
	// todo read file
	data, err := os.ReadFile("res/meditations.txt")
	if err != nil {
		panic(err)
	}

	text := strings.Split(string(data), "\n") // split

	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(text)
	_ = results // temp so it compiles.

	// todo print your result to stdout
	for word, count := range results {
		fmt.Println(word, count)
	}
}
