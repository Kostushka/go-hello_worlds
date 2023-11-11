package main

import (
	"fmt"
	"sort"
)

func main() {
	counts := map[string]int{}
	counts["hello"] = 2
	counts["bye"] = 3
	counts["afara"] = 10
	counts["regs"] = 5
	delete(counts, "hello")
	fmt.Printf("%d %d\n", counts["hello"], counts["bye"])
	for name, count := range counts {
		fmt.Printf("%s %d\n", name, count)
	}

	fmt.Println("---sort---")
	names := make([]string, 0, len(counts))
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s %d\n", name, counts[name])
	}

	v1, ok := counts["bob"]
	fmt.Println(v1, ok)
	v2, ok := counts["regs"]
	fmt.Println(v2, ok)
}
