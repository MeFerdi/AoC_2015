package main

import (
	"bufio"
	"fmt"
	"os"
)

// Check if the string contains a pair of any two letters that appears at least twice without overlapping
func hasRepeatingPair(s string) bool {
	pairs := make(map[string]int)

	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if pos, exists := pairs[pair]; exists && pos < i-1 {
			return true
		}
		pairs[pair] = i
	}

	return false
}

// Check if the string contains at least one letter that repeats with exactly one letter between them
func hasRepeatingLetterWithGap(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

// Check if the string is nice based on the new rules
func isNicer(s string) bool {
	return hasRepeatingPair(s) && hasRepeatingLetterWithGap(s)
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
		if isNicer(line) {
			niceCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(niceCount)
}
