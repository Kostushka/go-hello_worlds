package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
)

func main() {

	// файл с формой должен быть указан в аргументах командной строки при запуске сервера
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "there is no html file with the form in the command line args\n")
		os.Exit(1)
	}
	fmt.Println(os.Args[1])
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// прочитать содержимое файла в срез
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%T\n", f)
	fmt.Fprintf(w, "%s", string(f))
}
