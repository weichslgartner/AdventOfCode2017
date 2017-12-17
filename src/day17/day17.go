package main

import (
	"fmt"
)



func findTarget(buffer []int, targetNumber int) int{
	for i, element := range buffer {
		if element == targetNumber {
			return buffer[(i+1)%len(buffer)]
		}
	}
	fmt.Errorf("Number not found! Panic!11")
	return -1
}

func numberAfterZero(maxNumber int, stepsize int) int {
	insertnumber := 1
	currentPosition := 0
	numberAfterZero  := -1
	for insertnumber < maxNumber+1 {
		currentPosition = (currentPosition + stepsize) % insertnumber
		currentPosition += 1
		if currentPosition == 1{
			numberAfterZero = insertnumber
		}
		insertnumber++
	}
	return numberAfterZero
}

func createBuffer(maxNumber int, stepsize int) []int {
	buffer := make([]int, 1, maxNumber+1)
	insertNumber := 1
	currentPosition := 0
	buffer[currentPosition] = 0
	for insertNumber < maxNumber+1 {
		currentPosition = (currentPosition + stepsize) % len(buffer)
		currentPosition += 1
		if currentPosition == len(buffer) {
			buffer = append(buffer, insertNumber)
		} else {
			rest := make([]int, len(buffer[currentPosition:]))
			copy(rest, buffer[currentPosition:])
			buffer = append(buffer[0:currentPosition], insertNumber)
			for _, element := range rest {
				buffer = append(buffer, element)
			}
		}
		insertNumber++

	}
	return buffer
}

func main() {
	stepsize := 345
	buffer := createBuffer(2017, stepsize)
	part1 := findTarget(buffer, 2017)
	fmt.Println("Part 1:", part1)
	part2 := numberAfterZero(50e6, stepsize)
	fmt.Println("Part 2:", part2)


}