package main

import (
	"fmt"
	"util"
	"strconv"
)

func exitTheMaze(numbers []int, part1 bool) int {
	steps := 0
	instPointer := 0
	for instPointer < len(numbers) {
		jump := numbers[instPointer]
		if part1 || numbers[instPointer] < 3 {
			numbers[instPointer]++
		} else {
			numbers[instPointer]--
		}
		instPointer += jump
		steps++
	}
	return steps
}

func main() {
	lines := util.ReadFileLines("inputs/day5.txt")
	numbers := make([]int, len(lines))
	for i, element := range lines {
		number, _ := strconv.Atoi(element)
		numbers[i] = number
	}
	steps := 0
	//copy that the two parts don't alter the original input
	numberz := make([]int, len(lines))
	copy(numberz, numbers)
	steps = exitTheMaze(numberz, true)
	fmt.Println("Part 1: ", steps)
	copy(numberz, numbers)
	steps = exitTheMaze(numberz, false)
	fmt.Println("Part 2: ", steps)

}
