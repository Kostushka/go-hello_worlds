package main

import "fmt"
import "bytes"

type IntSet struct {
	words []uint64
}

func main() {
	arr := IntSet{}
	arr.Add(25)
	arr.Add(2)
	arr.Add(265)
	fmt.Println(arr.Has(264))
	arr2 := IntSet{}
	arr2.Add(4)
	arr.UnionWith(&arr2)
	fmt.Printf("arr: %s\narr2: %s\n", arr.String(), arr2.String())
	// fmt вызывает метод arr String, но метод доступен по указателю
	fmt.Println(&arr)
	// значение arr не имеет метод String
	fmt.Println(arr)
	fmt.Println(arr.Has(4))
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	fmt.Printf("x: %d x/64: %d x%%64: %d\n", x, word, bit)
	fmt.Printf("%d < %d\n", word, len(s.words))
	fmt.Printf("%b & %b\n", s.words[word], 1 << bit)
	fmt.Printf("%b\n", s.words[word]&(1 << bit))
	return word < len(s.words) && s.words[word]&(1 << bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
		fmt.Println("ok")
	}
	fmt.Printf("s.words: %v\n", s.words)
	fmt.Printf("len: %d\n", len(s.words))
	fmt.Printf("index: %d value: %b int: %b\n", word, s.words[word], 1 << bit)
	s.words[word] |= 1 << bit
	fmt.Printf("%d %b\n", s.words[word], s.words[word])
	fmt.Println(s.words)
}

func (s *IntSet) UnionWith(t *IntSet) {
	fmt.Printf("arr: %b arr2: %b\n", s.words, t.words)
	for i, tword := range t.words {
		if (i < len(s.words)) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
	fmt.Printf("arr: %b arr2: %b\n", s.words, t.words)
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1 << uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Printf("i: %d, j: %d 64*i+j: %d\n", i, j, 64*i+j)
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
