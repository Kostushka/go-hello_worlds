package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	dedup()
}

func dedup() {
	strings := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !strings[line] {
			strings[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
