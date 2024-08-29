package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadTextFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}
	return lines
}

func CalculateRibbon(str []string) int {
	totalRibbon := 0
	for _, present := range str {
		dimensions := make([]int, 3)
		dim := 0
		num := ""
		for _, c := range present {
			if c >= '0' && c <= '9' {
				num += string(c)
			} else if c == 'x' {
				dimensions[dim] = Atoi(num)
				dim++
				num = ""
			}
		}
		dimensions[dim] = Atoi(num)

		l, w, h := dimensions[0], dimensions[1], dimensions[2]
		ribbonToWrap := 2*min(l, w) + 2*min(w, h)
		ribbonForBow := l * w * h
		totalRibbon += ribbonToWrap + ribbonForBow
	}
	return totalRibbon
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Atoi(s string) int {
	sign := 1
	result := 0
	for i, char := range s {
		if i == 0 && char == '-' {
			sign = -1
		} else if i == 0 && char == '+' {
			sign = 1
		}
		if char > '0' || char < '9' {
			result = result*10 + int(char-'0')

		}
	}

	return result * sign
}

func main() {
	presentDimensions := loadTextFile("input.txt")

	totalRibbon := CalculateRibbon(presentDimensions)

	fmt.Println(totalRibbon)
}
