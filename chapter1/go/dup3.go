package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

/*
Поиск повторяющихся строк
Выводит кол-во появлений и текст каждой строки,
которая появляется в stdin более одного раза
*/
func main() {
	// создаю пустой хэш (отображение) с ключами типа string, значениями типа int
	counts := make(map[string]int)
	
	// прохожусь по аргументам командной строки
	for _, filename := range os.Args[1:] {
		// считываю все содержимое файла, возвращаю байтовый срез
		data, err := ioutil.ReadFile(filename)
		// обработка ошибки
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			continue
		}
		// прохожусь по строкам, на которые был разбит байтовый срез, преобразованный в string
		for _, line := range strings.Split(string(data), "\n") {
			// инкрементирую счетчик по ключу
			counts[line]++
		}
	}
	// прохожусь по хэшу
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
