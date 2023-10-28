package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, error := io.Copy(os.Stdout, resp.Body)
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy %s in stdout: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
