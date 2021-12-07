package day02

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

const horizontalPositionKeyName = "horizontalPosition"
const depthKeyName = "depth"

func Solution() {
	lines, _ := utils.ReadFile("day02/input.txt")

	fmt.Println("--- Day 02 ---")
	part01(lines)
}

func part01(lines []string) {
	position := map[string]int{horizontalPositionKeyName: 0, depthKeyName: 0}

	for _, line := range lines {
		movement, quantity := splitLine(line)

		if movement == "forward" {
			position[horizontalPositionKeyName] += quantity

		} else if movement == "up" {
			position[depthKeyName] -= quantity

		} else if movement == "down" {
			position[depthKeyName] += quantity
		}
	}

	fmt.Println("What do you get if you multiply your final horizontal position by your final depth?")
	fmt.Printf("Answer: %d\n\n", position[horizontalPositionKeyName]*position[depthKeyName])
}

func splitLine(line string) (movement string, quantity int) {
	split := strings.Split(line, " ")

	if len(split) != 2 {
		return "", 0
	}

	q, _ := strconv.Atoi(split[1])
	return split[0], q
}
