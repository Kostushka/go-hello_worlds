package main

import (
	"bytes"
	"fmt"
)

func main() {

	str := "1234567"

	fmt.Printf("%s\t%s\n", str, comma(str))
}

func comma(s string) string {
	var buf bytes.Buffer

	count := 3
	for i := len(s) - 1; i >= 0; i-- {
		if count == 0 {
			buf.WriteString(",")
			count = 3
		}

		fmt.Fprintf(&buf, "%c", s[i])

		count--
	}
	
	return reverse(buf.String())
}

func reverse(s string) string {
	n := len(s)
	// строку -> байтовый срез для изменения
	str := []byte(s)
	for i, j := 0, n - 1; i < j; {
		str[i], str[j] = str[j], str[i]
		i++
		j--		
	}

	return string(str)
}
