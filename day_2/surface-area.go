package main

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	loadText("input.txt")
}

func loadText(filename string) []string {
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
func CalculateWraper(str []string) int {
	totalWrappingPaper := 0
	for _, present := range str {
		l, w, h := 0, 0, 0
		dim := 0
		for _, c := range present {
			if c >= '0' && c <= '9' {
				if dim == 0 {
					l = l*10 + int(c-'0')
				} else if dim == 1 {
					w = w*10 + int(c-'0')
				} else {
					h = h*10 + int(c-'0')
				}
			} else if c == 'x' {
				dim++
			}
		}

		surfaceArea := 2*l*w + 2*w*h + 2*h*l
		side1 := l * w
		side2 := w * h
		side3 := h * l
		minSide := side1
		if side2 < minSide {
			minSide = side2
		}
		if side3 < minSide {
			minSide = side3
		}

		totalWrappingPaper += surfaceArea + minSide
	}
	return totalWrappingPaper
}
// func main() {
// 	presentDimensions := loadText("input.txt")

// 	totalWrappingPaper := CalculateWraper(presentDimensions)

// 	fmt.Println(totalWrappingPaper)
// }
