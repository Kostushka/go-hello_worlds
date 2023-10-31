package main

import "fmt"
import "strings"

func main() {
	str := "tmp/dir/file.c"
	fmt.Printf("basename: %s\tstr: %s\n", basename(str), str)
}

func basename(str string) string {
	slash := strings.LastIndex(str, "/")
	str = str[slash+1:]
	if dot := strings.LastIndex(str, "."); dot >= 0 {
		str = str[:dot]
	}

	return str
}
