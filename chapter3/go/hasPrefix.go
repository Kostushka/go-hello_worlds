package main

import "fmt"

func main() {
	fmt.Printf("%t\n", MyHasPrefix("Hello", "He"))
	fmt.Printf("%t\n", MyHasPrefix("Hello", "he"))
	fmt.Printf("%t\n", MyHasSuffix("Hello", "lo"))
	fmt.Printf("%t\n", MyHasSuffix("Hello", "ll"))
	fmt.Printf("%t\n", MyContains("Hello", "ll"))
	fmt.Printf("%t\n", MyContains("Hello", "la"))
}

func MyHasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func MyHasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix
}

func MyContains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		// if (MyHasPrefix(s[i:], substr)) {
			// return true
		// }
		temp := s[i:]
		if (len(temp) >= len(substr) && temp[:len(substr)] == substr) {
			return true
		}
	}
	return false
}
