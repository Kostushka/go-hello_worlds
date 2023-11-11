package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Printf("%d\n", a)
	// надо присвоить, ибо возвращается новый срез
	a = append(a, 4)
	fmt.Printf("%d\n", a)
	a = append(a, 5, 6)
	fmt.Printf("%d\n", a)
	a = append(a, a...)
	fmt.Printf("%d\n", a)
	a = MyAppend(a, 10)
	fmt.Printf("%d\n", a)
}

func MyAppend(x []int, y int) []int {
	// срез
	var z []int
	fmt.Printf("%d %d\n", len(x) + 1, cap(x))
	// длина текущего среза + 1 новый элемент
	zlen := len(x) + 1
	// длина расширенного среза должна быть меньше или равна емкости
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		// емкость увеличить в два раза
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2*len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	fmt.Printf("%d %d\n", len(x) + 1, cap(z))
	// добавить элемент в новый срез
	z[len(x)] = y
	return z
}
