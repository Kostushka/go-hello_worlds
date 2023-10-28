package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// возвращает http статус ответа
		fmt.Printf("%s\n", resp.Status)
	}
}
