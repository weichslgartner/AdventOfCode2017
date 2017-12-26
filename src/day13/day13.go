package main

import (
	"util"
	"fmt"
)

//const LENGTH = 91

type Wall struct {
	depth      int
	wallRange  int
	currentPos int
	currentDir int
}

func getWalls(lines []string) []Wall {
	walls := make([]Wall, 0, 10000)
	for _, line := range lines {
		numbers := util.ExtractAllNumbers(line)
		depth := numbers[0]
		if depth >= len(walls) {
			walls = walls[0:depth+1]
		}
		walls[depth] = Wall{depth, numbers[1], 0, 1}
	}
	return walls
}

func getHit(walls []Wall, delay int) bool {
	newWalls := make([]Wall, len(walls))
	copy(newWalls, walls)

	updatePositions(newWalls, delay)
	//fmt.Println(delay, newWalls)
	for i := 0; i < len(walls); i++ {
		currentWall := newWalls[i]
		if currentWall.wallRange > 0 {
			if currentWall.currentPos == 0 {
				return true
			}
		}

		updatePositionsOneStep(newWalls)
		//fmt.Println(walls)
	}
	return false
}

func updatePositions(walls []Wall, delay int) {
	for i, wall := range walls {
		if wall.wallRange > 0 {
			//reduces delay to delay inside the cycle
			iDelay := delay % (wall.wallRange*2 - 2)
			for j := 0; j < iDelay; j++ {
				wall = updateSingleWall(wall)
			}
			walls[i] = wall

		}

	}
}

func updatePositionsOneStep(walls []Wall) {
	for i, wall := range walls {
		if wall.wallRange > 0 {
			walls[i] = updateSingleWall(wall)
		}

	}
}
func updateSingleWall(wall Wall) Wall {
	if wall.currentPos+wall.currentDir > wall.wallRange-1 {
		wall.currentDir = -1
	} else if wall.currentPos+wall.currentDir < 0 {
		wall.currentDir = +1
	}
	wall.currentPos = wall.currentPos + wall.currentDir
	return wall
}

func part1(walls []Wall) {
	severity := 0
	for i := 0; i < len(walls); i++ {
		currentWall := walls[i]
		if currentWall.wallRange > 0 {
			if currentWall.currentPos == 0 {
				severity += currentWall.depth * currentWall.wallRange
			}
		}
		updatePositionsOneStep(walls)
		//fmt.Println(walls)
	}
	fmt.Println("Part 1:", severity)
}

func part2(walls []Wall) {
	delay := 0
	for getHit(walls, delay) {
		delay++
	}
	fmt.Println("Part 2:", delay)
}

func main() {
	lines := util.ReadFileLines("inputs/day13.txt")
	walls := getWalls(lines)
	part1(walls)
	walls = getWalls(lines)
	part2(walls)
	//fmt.Println(walls)
}
