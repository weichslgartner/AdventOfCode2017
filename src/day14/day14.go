package main

import (
	"util"
	"fmt"
	"strconv"
)

const FREE = 0
const OCCUPIED = -1
const WIDTH = 128
const HEIGHT = 128

var diskArray [][]int

func toOccupiedFree(array []int) []int {
	result := make([]int, len(array))
	for i, cell := range array {
		if cell == 1 {
			result[i] = OCCUPIED
		} else {
			result[i] = FREE
		}

	}
	return result
}

func printGrid() {
	for _, line := range diskArray {
		for _, char := range line {

			if char == OCCUPIED {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func printArray() {
	for _, line := range diskArray {
		for _, char := range line {

			fmt.Printf("|%04d", char)
		}
		fmt.Println()
	}
}

func findRegions() int {
	numberRegions := 1
	for i, line := range diskArray {

		for j, char := range line {
			if char == OCCUPIED {
				searchRegion(i, j, numberRegions)
				numberRegions++
			}
		}
	}
	//fmt.Println(numberRegions-1)
	return numberRegions - 1
}

func searchRegion(y int, x int, regionNumber int) {
	diskArray[y][x] = regionNumber
	if y > 0 && diskArray [y-1][x] == OCCUPIED {
		searchRegion(y-1, x, regionNumber)
	}
	if y < HEIGHT-1 && diskArray [y+1][x] == OCCUPIED {
		searchRegion(y+1, x, regionNumber)
	}
	if x > 0 && diskArray [y][x-1] == OCCUPIED {
		searchRegion(y, x-1, regionNumber)
	}
	if x < WIDTH-1 && diskArray [y][x+1] == OCCUPIED {
		searchRegion(y, x+1, regionNumber)
	}

	return
}

func main() {
	numberOnes := 0
	input := "ljoxqyyw" //"flqrgnkx"//
	diskArray = make([][]int, HEIGHT)

	for i := 0; i < 128; i++ {
		hashInput := input + "-" + strconv.Itoa(i)
		hash := util.KnotHash(hashInput)
		binHash := util.StringToBin(hash)
		diskArray[i] = toOccupiedFree(util.StringToIntArray(binHash))
		numberOnes += util.HammingWeight(binHash)
	}
	//printGrid()
	//printArray()
	fmt.Println("Part 1:", numberOnes)
	fmt.Println("Part 2:", findRegions())

}
