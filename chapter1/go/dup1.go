package main

import (
	"fmt"
	"bufio"
	"os"
)
// Поиск повторяющихся строк
func main() {
	// создаю пустой хэш (отображение) с ключами типа string, значениями типа int
	counts := make(map[string]int)
	
	// переменная input ссылается на функцию, считывающую данные из stdin
	input := bufio.NewScanner(os.Stdin)
	
	// считывает строку без \n, возвращает true или false
	for input.Scan() {
		// по ключу (считанной строке) инкрементируем значение (по умолчанию 0)
		counts[input.Text()]++ 
	}
	
	// цикл по диапазону: получаем ключ и значение
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
