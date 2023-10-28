package main

import (
	"fmt"
	"bufio"
	"os"
)

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
			// новый хэш для подсчета повторяющихся строк одного файла
			hash := make(map[string]int)
			// получаем открытый файл и err
			file, err := os.Open(arg)
			// обработка err
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
				continue
			}
			countLines(file, hash)
			// закрываем файл
			file.Close()

			for _, n := range hash {
				if n > 1 {
					fmt.Printf("%s\n", arg)
					break
				}
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		counts[input.Text()]++
	}
}
