package main

import (
	"strconv"
	"util"
	"fmt"

)

const INFECTED = "I"
const CLEAN = "C"
const WEAKENED = "W"
const FLAGGED = "F"



type cell struct{
	x int
	y int
	infected string
}

func coord2String(x int, y int) string{
	return strconv.Itoa(x) + "," +strconv.Itoa(y)
}

func cell2String(c cell) string{
	return coord2String(c.x, c.y )
}





func part1(grid map[string]cell, currentX int, currentY int, dirX int, dirY int, maxBursts int) int {
	numberInfected :=0
	for bursts := 0; bursts < maxBursts; bursts++ {
		element, exists := grid[coord2String(currentX, currentY)]
		if !exists {
			element = cell{currentX, currentY, CLEAN}
		}
		if element.infected == INFECTED{
			dirX, dirY = turnRight(dirX, dirY)
			element.infected = CLEAN
		} else {
			dirX, dirY = turnLeft(dirX, dirY)
			element.infected = INFECTED
			numberInfected++
		}
		grid[coord2String(currentX, currentY)] = element
		currentY += dirY
		currentX += dirX

	}
	return numberInfected
}


func part2(grid map[string]cell, currentX int, currentY int, dirX int, dirY int,  maxBursts int) int {
	numberInfected :=0
	for bursts := 0; bursts < maxBursts; bursts++ {
		element, exists := grid[coord2String(currentX, currentY)]
		if !exists {
			element = cell{currentX, currentY, CLEAN}
		}

		switch element.infected{
		case INFECTED:
			dirX, dirY = turnRight(dirX, dirY)
			element.infected = FLAGGED
		case WEAKENED:
			element.infected = INFECTED
			numberInfected++
		case FLAGGED:
			dirX, dirY = reverse(dirX, dirY)
			element.infected = CLEAN
		case CLEAN:
			dirX, dirY = turnLeft(dirX, dirY)
			element.infected = WEAKENED
		default:
			fmt.Errorf("unknown cell state")
		}

		grid[coord2String(currentX, currentY)] = element
		currentY += dirY
		currentX += dirX

	}
	return numberInfected
}


func reverse(x int, y int) (int, int) {
	return -1*x,-1*y
}

func turnRight(x int, y int) (int, int) {
	if x == 0 && y == 1 {
		return -1,0
	}else if x == 0 && y == -1 {
		return 1,0
	}else if x == -1 && y == 0 {
		return 0,-1
	}else if x == 1 && y == 0 {
		return 0,1
	}
	fmt.Errorf("not defined direction")
	return 0,0
}


func turnLeft(x int, y int) (int, int) {
	if x == 0 && y == 1 {
		return 1,0
	}else if x == 0 && y == -1 {
		return -1,0
	}else if x == -1 && y == 0 {
		return 0,1
	}else if x == 1 && y == 0 {
		return 0,-1
	}
	fmt.Errorf("not defined direction")
	return 0,0
}


func parseGrid(lines []string) (map[string]cell,int,int) {
	maxY := 0
	maxX := 0
	grid := make(map[string]cell)
	for y, line := range lines {
		if y > maxY {
			maxY = y
		}
		for x, char := range line {
			infected := CLEAN
			if string(char) == "#" {
				infected = INFECTED
			}
			currentCell := cell{x, y, infected}
			grid[cell2String(currentCell)] = currentCell
			if x > maxX {
				maxX = x
			}

		}
	}
	return grid ,maxY, maxX
}


func main() {
	lines := util.ReadFileLines("inputs/day22.txt")
	minY := 0
	minX := 0

	grid ,maxY, maxX:=  parseGrid(lines)

	currentX := (maxX - minX)/2
	currentY := (maxY - minY)/2
	dirX := 0
	dirY := -1
	maxBursts := 10000

	numberInfected := part1(grid, currentX, currentY, dirX, dirY, maxBursts)
	fmt.Println("Part 1: ", numberInfected)


	maxBursts = 10000000
	numberInfected = part2(grid, currentX, currentY, dirX, dirY, maxBursts)
	fmt.Println("Part 2: ", numberInfected)
}