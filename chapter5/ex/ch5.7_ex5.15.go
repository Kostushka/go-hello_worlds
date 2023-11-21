package main

import "fmt"
import "os"

func main() {
	fmt.Println(max(5, 4, 2, 1, 9))
	fmt.Println(min(5, 4, 2, 1, 9))
	fmt.Println(max())
	fmt.Println(min())
}

func max(vals ...int) int {
	if len(vals) == 0 {
		fmt.Fprintf(os.Stderr, "нужен хотя бы один аргумент, текущее кол-во аргументов: %d\n", len(vals))
		os.Exit(0)
	}
	max := 0
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	if len(vals) == 0 {
		fmt.Fprintf(os.Stderr, "нужен хотя бы один аргумент, текущее кол-во аргументов: %d\n", len(vals))
		os.Exit(0)
	}
	min := vals[0]
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min
}
