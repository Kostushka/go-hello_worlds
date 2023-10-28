package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"net/http"
)

func main() {

	prefix := "http://"
	
	for _, url := range os.Args[1:] {
		// префикс http должен быть
		if !strings.HasPrefix(url, prefix) {
			// добавить префикс
			url = prefix + url
			fmt.Printf("url: %s\n", url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %s: %v\n", url, err)
		}
		fmt.Printf("%s", b)
	}
}
