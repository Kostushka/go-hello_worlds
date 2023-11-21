package main

import "fmt"
import "bytes"

type IntSet struct {
	words []uint64
}

func main() {
	arr := IntSet{}
	arr.Add(25)
	arr.Add(265)
	arr.Add(2)
	// возвращает адрес структуры
	r := arr.Copy()
	fmt.Println(&arr)
	fmt.Println("кол-во элементов:", arr.Len())
	arr.Remove(2)
	arr.Clear()
	fmt.Println(&arr)
	fmt.Println("кол-во элементов:", arr.Len())
	fmt.Println("copy:", r)
	fmt.Println("кол-во элементов:", r.Len())
	r.AddAll(3, 4, 5, 7)
	fmt.Println(r)
	fmt.Println("кол-во элементов:", r.Len())
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1 << bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if (i < len(s.words)) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
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
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// кол-во элементов
func (s *IntSet) Len() int {
	res := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1 << uint(j)) != 0 {
				res++
			}
		}
	}
	return res
}

// удаляет элемент
func (s *IntSet) Remove(x int) {
	word, bit := x/64, x%64
	if s.words[word]&(1 << bit) != 0 {
		fmt.Printf("%b\n", s.words[word])
		// &^ - заменяет в ответе 1 во втором операнде на 0, иначе берет бит из первого
		s.words[word] &= 0xffffffffffffffff &^ (1 << bit)
		fmt.Printf("%b\n", s.words[word])
	}
}

// удаляет все элементы
func (s *IntSet) Clear() {
	if len(s.words) > 0 {
		s.words = []uint64{}
	}
}

// копия множества
func (s *IntSet) Copy() *IntSet {
	res := IntSet{}
	res.words = make([]uint64, len(s.words))
	copy(res.words, s.words)
	// res.words = append(res.words, s.words...)
	return &res
}

// добавить несколько значений
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		word, bit := val/64, val%64
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}
