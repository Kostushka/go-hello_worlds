package main

import (
	"fmt"
	"bufio"
	"os"
)
/*
Поиск повторяющихся строк
Выводит кол-во появлений и текст каждой строки,
которая появляется в stdin/именованном файле более одного раза
*/
func main() {
	// создаю пустой хэш (отображение) с ключами типа string, значениями типа int
	counts := make(map[string]int)
	
	// переменная с массивом аргументов командной строки
	files := os.Args[1:]
	
	if len(files) == 0 {
		// аргументов нет, читаем из stdin
		countLines(os.Stdin, counts)
	} else {
		// аргументы есть, читаем из файлов
		for _, arg := range files {
			// получаем открытый файл и err
			file, err := os.Open(arg)
			// обработка err
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
				continue
			}
			countLines(file, counts)
			// закрываем файл
			file.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			if line == "" {
				fmt.Printf("%d\t\"\"\n", n)
			} else {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		counts[input.Text()]++
	}
}
