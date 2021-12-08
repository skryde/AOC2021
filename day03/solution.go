package day03

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
)

type asd struct {
	count0 int
	count1 int
}

func Solution() {
	lines, _ := utils.ReadFile("day03/input.txt")

	fmt.Println("--- Day 03 ---")
	part01(lines)
	part02(lines)
}

func part01(lines []string) {

	var mostCommonBits []asd

	for _, line := range lines {
		for i, digit := range line {
			if len(mostCommonBits) != len(line) {
				mostCommonBits = append(mostCommonBits, asd{
					count0: 0,
					count1: 0,
				})
			}

			if digit == '0' {
				mostCommonBits[i].count0++

			} else if digit == '1' {
				mostCommonBits[i].count1++
			}
		}
	}

	gammaRateStr := ""
	epsilonRateStr := ""

	for _, commonBit := range mostCommonBits {
		if commonBit.count0 > commonBit.count1 {
			gammaRateStr += "0"
			epsilonRateStr += "1"

		} else {
			gammaRateStr += "1"
			epsilonRateStr += "0"
		}
	}

	gammaRate, _ := strconv.ParseInt(gammaRateStr, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonRateStr, 2, 64)

	fmt.Println("What is the power consumption of the submarine?")
	fmt.Printf("Answer: %d\n\n", gammaRate*epsilonRate)
}

func part02(lines []string) {

}
