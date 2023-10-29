package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	// связывает функцию обработчик с /, запускает сервер с портом 8000
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// путь отправить в качестве ответа
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}
