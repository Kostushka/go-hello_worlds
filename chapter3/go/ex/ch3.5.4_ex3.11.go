package main

import (
	"bytes"
	"fmt"
)

func main() {

	str := "12-34+5.4678+9.4"

	fmt.Printf("%s\t%s\n", str, comma(str))
}

func comma(s string) string {
	var buf bytes.Buffer

	count := 3
	// идем по строке с конца
	for i := len(s) - 1; i >= 0; i-- {
		// запись в байтовый срез символа строки
		fmt.Fprintf(&buf, "%c", s[i])

		if i == 0 {
			break
		}

		flag := false
		// учет предыдущего символа как продолжение строкового представления числа
		// (строковое представление числа может включать не один строковый символ)
		switch s[i-1] {
		case '.':
		case '-':
		case '+':
			flag = true
		}
		if flag {
			continue
		}

		// учет точки как продолжение строкового представления числа
		if s[i] == '.' {
			continue
		}

		// три числа пропущено, вставка запятой
		if count == 0 {
			buf.WriteString(",")
			count = 3
			continue
		}
		// декремент счетчика пропущенных чисел
		count--
	}

	return reverse(buf.String())
}

func reverse(s string) string {
	n := len(s)
	// строку -> байтовый срез для изменения
	str := []byte(s)
	for i, j := 0, n-1; i < j; {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}

	return string(str)
}
