package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
)

func main() {

	// проход по аргументам командной строки
	for _, url := range os.Args[1:] {
		// http.Get выполняет http запрос, возвращает результат в структуре resp
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// Поле Body в структуре resp содержит ответ сервера в виде потока, доступного для чтения
		// считываем весь ответ и сохраняем в переменной
		b, err := ioutil.ReadAll(resp.Body)
		// поток закрываем
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
