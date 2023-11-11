package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	// "bufio"
	"os"
)

func main() {
	words, img, err := CountWordsAndImages("https://ya.ru")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	fmt.Println(words, img)
}

func CountWordsAndImages(url string) (words, img int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return 
	}
	words, img = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, img int) {
	// input := bufio.NewScanner(n)
	// input.Split(bufio.ScanWords)
	// for input.Scan() {
		words++
		if n.Type == html.ElementNode && n.Data == "img" {
			img++
		}
		fmt.Println(n.Type)
	// }
	return
}
