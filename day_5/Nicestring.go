package main

import (
	"bufio"
	"fmt"
	"os"
)

func isNice(s string) bool {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	disallowed := map[string]bool{"ab": true, "cd": true, "pq": true, "xy": true}

	vowelCount := 0
	hasDoubleLetter := false

	// Check for disallowed substrings and other conditions
	for i := 0; i < len(s); i++ {
		// Count vowels
		if vowels[rune(s[i])] {
			vowelCount++
		}

		// Check for double letters
		if i > 0 && s[i] == s[i-1] {
			hasDoubleLetter = true
		}

		// Check for disallowed substrings
		if i > 0 {
			pair := s[i-1 : i+1]
			if disallowed[pair] {
				return false
			}
		}
	}

	// Must have at least 3 vowels and a double letter
	return vowelCount >= 3 && hasDoubleLetter
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	niceCount := 0

	// Process each line (string) in the file
	for scanner.Scan() {
		line := scanner.Text()
		if isNice(line) {
			niceCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Number of nice strings:", niceCount)
}
