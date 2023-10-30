package main

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"io/ioutil"
	"time"
)

/* 
Выполняет параллельную выборку URL и сообщает
о затраченном времени и размере ответа для каждого из них
*/
func main() {

	start := time.Now()
	// создать канал строк
	ch := make(chan string)

	// пройтись по массиву урлов
	for _, url := range os.Args[1:] {
		// для каждого урла запустить подпрограмму
		go fetch(url, ch)
	}

	// получить строки из канала и вывести
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	// получить структуру ответа
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// игнорируем тело ответа
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	// время, которое прошло с момента запуска подпрограммы
	secs := time.Since(start).Seconds()
	// отправить строку с данными в канал
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
