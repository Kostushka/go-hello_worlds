package main

import (
	"fmt"
	"crypto/sha256"
)
// кол-во битов, различных в двух 256-битовых дайджестах
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x %x %t %T\n", c1, c2, c1 == c2, c1)
	// for _, v := range c1 {
		// fmt.Printf("%d\t%[1]x\t%08[1]b\n", v)
	// }
	var count int64
	for i := 0; i < len(c1); i++ {
		// fmt.Printf("%d\t%[1]x\t%08[1]b\n", c1[i])
		// fmt.Printf("%d\t%[1]x\t%08[1]b\n", c2[i])
		// fmt.Println()
		for j := 8; j > 0; j-- {
			fmt.Printf("%b %b %d %d %t\n", c1[i], c2[i], c1[i]&1, c2[i]&1, c1[i]&1 == c2[i]&1)
			if !(c1[i]&1 == c2[i]&1) {
				count++
			}
			c1[i] >>= 1
			c2[i] >>= 1
		}
		fmt.Println(count)
	}
	fmt.Println(count)
}
