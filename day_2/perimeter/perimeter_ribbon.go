package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func CalculateRibbon(lines []string) int {
	totalRibbon := 0
	for _, line := range lines {
		dimensions := strings.Split(line, "x")
		if len(dimensions) != 3 {
			fmt.Printf("Error: Invalid dimension format in line: %s\n", line)
			os.Exit(1)
		}

		l, err1 := strconv.Atoi(dimensions[0])
		w, err2 := strconv.Atoi(dimensions[1])
		h, err3 := strconv.Atoi(dimensions[2])
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Printf("Error: Invalid number in line: %s\n", line)
			os.Exit(1)
		}

		// Calculating the ribbon to wrap the present (smallest perimeter)
		perimeter1 := 2 * (l + w)
		perimeter2 := 2 * (w + h)
		perimeter3 := 2 * (h + l)
		ribbonToWrap := min(perimeter1, perimeter2, perimeter3)

		//the ribbon for the bow (volume of the present)
		ribbonForBow := l * w * h

		totalRibbon += ribbonToWrap + ribbonForBow
	}
	return totalRibbon
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}

func main() {
	presentDimensions := loadTextFile("input.txt")
	totalRibbon := CalculateRibbon(presentDimensions)
	fmt.Println(totalRibbon)
}
