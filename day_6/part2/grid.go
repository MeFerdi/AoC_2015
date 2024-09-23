package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const gridSize = 1000

// Initialize a 1000x1000 grid of lights, all with brightness 0.
var lights [gridSize][gridSize]int

// Main function to execute the application
func main() {
	// Specify the path to your input file
	filePath := "input.txt" // Change this to your file path

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Process each instruction from the file
	for scanner.Scan() {
		instruction := strings.TrimSpace(scanner.Text())
		if instruction != "" {
			processInstruction(instruction)
		}
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Calculate the total brightness of all lights
	totalBrightness := calculateTotalBrightness()
	fmt.Println("Total brightness of all lights:", totalBrightness)
}

// Process a single instruction to modify the lights grid
func processInstruction(instruction string) {
	var x1, y1, x2, y2 int

	// Determine if the instruction is "turn on", "turn off", or "toggle"
	if strings.HasPrefix(instruction, "turn on") {
		fmt.Sscanf(instruction, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		// Increase brightness by 1
		applyAction("turn on", x1, y1, x2, y2)
	} else if strings.HasPrefix(instruction, "turn off") {
		fmt.Sscanf(instruction, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		// Decrease brightness by 1
		applyAction("turn off", x1, y1, x2, y2)
	} else if strings.HasPrefix(instruction, "toggle") {
		fmt.Sscanf(instruction, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		// Increase brightness by 2
		applyAction("toggle", x1, y1, x2, y2)
	} else {
		fmt.Printf("Instruction '%s' is malformed\n", instruction)
	}
}

// Apply the action to the specified grid area
func applyAction(action string, x1, y1, x2, y2 int) {
	// Normalize coordinates to ensure x1, y1 is the top left and x2, y2 is the bottom right
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	// Perform the appropriate action on the grid
	for i := y1; i <= y2; i++ {
		for j := x1; j <= x2; j++ {
			switch action {
			case "turn on":
				lights[i][j] += 1
			case "turn off":
				if lights[i][j] > 0 {
					lights[i][j] -= 1
				}
			case "toggle":
				lights[i][j] += 2
			}
		}
	}
}

// Calculate the total brightness of all lights
func calculateTotalBrightness() int {
	totalBrightness := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			totalBrightness += lights[i][j]
		}
	}
	return totalBrightness
}
