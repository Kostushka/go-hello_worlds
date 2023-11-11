package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	hash := make(map[string]int)
	outline(nil, doc, hash)
	for key, value := range hash {
		fmt.Println(key, value)
	}
}

func outline(stack []string, n *html.Node, hash map[string]int) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		hash[n.Data]++
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c, hash)
	}
}
