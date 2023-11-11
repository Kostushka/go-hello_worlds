package main

import (
	"fmt"
	"bufio"
	"os"
)

// подсчет слов
func main() {
	wordfreq()
}

func wordfreq() {
	words := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// разбить на слова
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		words[word]++
	}

	for word, count := range words {
		fmt.Printf("%s\t%d\n", word, count)
	} 
	
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
}
