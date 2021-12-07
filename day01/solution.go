package day01

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
)

func Solution() {
	lines, _ := utils.ReadFile("day01/input.txt")

	fmt.Println("--- Day 01 ---")
	part01(lines)
	part02(lines)
}

func part01(lines []string) {
	count := 0
	i := 0

	for j := 1; j < len(lines); j++ {
		iNum, _ := strconv.Atoi(lines[i])
		jNum, _ := strconv.Atoi(lines[j])

		if iNum < jNum {
			count++
		}

		i++
	}

	fmt.Println("How many measurements are larger than the previous measurement?")
	fmt.Printf("Answer: %d\n\n", count)
}

func part02(lines []string) {

	countMap := map[rune]int{}
	letter := 'A'
	count := 0

	for i := 2; i < len(lines); i++ {
		a, _ := strconv.Atoi(lines[i])
		b, _ := strconv.Atoi(lines[i-1])
		c, _ := strconv.Atoi(lines[i-2])

		countMap[letter] = a + b + c
		letter++
	}

	for i := 'B'; i < letter; i++ {
		prev := countMap[i - 1]
		if prev < countMap[i] {
			count++
		}
	}

	fmt.Println("How many sums are larger than the previous sum?")
	fmt.Printf("Answer: %d\n\n", count)
}
