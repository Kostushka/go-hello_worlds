package main

import "fmt"

func main() {
	fmt.Printf("%t\n", anagram("hello", "lhelo"))
	fmt.Printf("%t\n", anagram("hello", "lheo"))
	fmt.Printf("%t\n", anagram("комар", "корма"))
	fmt.Printf("%t\n", anagram("Привет", "тевирП"))
}

func anagram(s, t string) bool {
	// длины строк должны совпадать
	if len(s) != len(t) {
		return false
	}
	// байтовый срез второй строки
	buf := []byte(t)
	// идем по каждому символу первой строки
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(buf); j++ {
			// во второй строке должен быть символ аналогичный текущему символу первой строки
			if s[i] == buf[j] {
				buf[j] = 0
				break
			}
			if j == len(buf)-1 {
				return false
			}
		}
	}
	return true
}
