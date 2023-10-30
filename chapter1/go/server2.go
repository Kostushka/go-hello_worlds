package main

import (
	"fmt"
	"net/http"
	"sync"
	"log"
)

var mu sync.Mutex
var count int

// веб-сервер, возвращает компонент пути из URL и считает кол-во запросов
func main() {
	// два обработчика двух урлов
	http.HandleFunc("/", handler) // каждый обработчик запускается в отдельной go-подпрограмме
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// доступ к общей переменной должен быть обеспечен только в одном потоке, остальные ждут
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	// доступ к общей переменной должен быть обеспечен только в одном потоке, остальные ждут
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
