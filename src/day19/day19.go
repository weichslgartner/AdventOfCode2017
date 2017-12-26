package main

import (
	"util"
	"fmt"
)

type coord struct {
	x int
	y int
}



func doTheWalk(currentPos coord, ymax int, xmax int, array [][]string) (string, int) {
	xDir := 0
	yDir := 1
	solution := ""
	visited := make(map[coord]bool)
	numberSteps := 0
	for currentPos.y < ymax && currentPos.y >= 0 && currentPos.x < xmax && currentPos.x >= 0 {
		currentChar := array[currentPos.y][currentPos.x]
		if currentChar == " " {
			break
		}
		visited[currentPos] = true
		switch currentChar {
		case "|":
		case "-":
		case "+":
			var next coord
		Loop:
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					next = coord{currentPos.x + x, currentPos.y + y}
					if next.y >= ymax || next.y < 0 || next.x >= xmax || next.x < 0 {
						continue
					}
					//only cardinal neighbors; skip the diagonal ones
					if util.Abs(x)+util.Abs(y) == 2 {
						continue
					}
					if visited[next] == false && array[next.y][next.x] != " " && array[next.y][next.x] != "" {
						xDir = x
						yDir = y
						break Loop
					}
				}
			}
		default:
			solution += currentChar
			//fmt.Println(solution)
		}
		currentPos.x += xDir
		currentPos.y += yDir
		numberSteps++
		//fmt.Println(currentPos , currentChar)
	}
	return solution, numberSteps
}

func createGrid(lines []string) ([][]string, coord, int, int) {
	grid := make([][]string, len(lines))
	currentPos := coord{0, 0}
	yMax := len(lines)
	xMax := len(lines)
	//quadratic
	for y, line := range lines {
		grid[y] = make([]string, len(lines))
		for x, char := range line {
			grid[y][x] = string(char)
			if y == 0 && string(char) == "|" {
				currentPos.x = x
			}
		}
	}
	return grid, currentPos, yMax, xMax
}


func main() {
	lines := util.ReadFileLines("inputs/day19.txt")
	array, currentPos, ymax, xmax := createGrid(lines)
	solution, numberSteps := doTheWalk(currentPos, ymax, xmax, array)
	fmt.Println(solution, numberSteps)
}