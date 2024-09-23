package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const gridSize = 1000

// Initialize a 1000x1000 grid of lights, all turned off (false).
var lights [gridSize][gridSize]bool

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

	// Count the number of lights that are lit
	numLightsLit := countLitLights()
	fmt.Println("Number of lights that are lit:", numLightsLit)
}

// Process a single instruction to modify the lights grid
func processInstruction(instruction string) {
	var x1, y1, x2, y2 int

	// Determine if the instruction is "turn on", "turn off", or "toggle"
	if strings.HasPrefix(instruction, "turn on") {
		fmt.Sscanf(instruction, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		// Turn on the lights
		applyAction("turn on", x1, y1, x2, y2)
	} else if strings.HasPrefix(instruction, "turn off") {
		fmt.Sscanf(instruction, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		// Turn off the lights
		applyAction("turn off", x1, y1, x2, y2)
	} else if strings.HasPrefix(instruction, "toggle") {
		fmt.Sscanf(instruction, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
		// Toggle the lights
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
				lights[i][j] = true
			case "turn off":
				lights[i][j] = false
			case "toggle":
				lights[i][j] = !lights[i][j]
			}
		}
	}
}

// Count the number of lights that are on
func countLitLights() int {
	count := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if lights[i][j] {
				count++
			}
		}
	}
	return count
}
