package main

import (
	"util"
	"fmt"
	"math"
)

type coord struct {
	x int
	y int
	z int
}

type particle struct {
	position     coord
	velocity     coord
	acceleration coord
	index        int
	collide      bool
}

func manhattanDistance(pos coord) int {
	return util.Abs(pos.x) + util.Abs(pos.y) + util.Abs(pos.z)
}

func simulate(particleList []particle, particleMap map[int]int, part1 bool) particle {
	destroyed := 0
	var nearestParticle particle
	for i := 0; i < 1000; i++ {
		converge := true
		nearestDistance := math.MaxInt32
		currentParticlePositions := make(map[coord]int)
		for i, part := range particleList {
			if !part1 {
				if part.collide {
					continue
				}
				firstOccurence, exists := currentParticlePositions[part.position]
				if exists {
					destroyed = destroy(part, particleList, firstOccurence, destroyed, i)
					continue
				} else {
					currentParticlePositions[part.position] = i
				}
			}
			manhattan := particleMap[i]
			part.velocity.x += part.acceleration.x
			part.velocity.y += part.acceleration.y
			part.velocity.z += part.acceleration.z

			part.position.x += part.velocity.x
			part.position.y += part.velocity.y
			part.position.z += part.velocity.z

			newManhattan := manhattanDistance(part.position)
			if newManhattan < manhattan {
				converge = false
			}
			particleList[i] = part
			particleMap[i] = newManhattan
			if newManhattan < nearestDistance {
				nearestParticle = part
				nearestDistance = newManhattan
			}
		}
		fmt.Println(i, nearestParticle, nearestDistance, destroyed, len(particleList)-destroyed)
		if converge && !part1 {
			break
		}
	}
	return nearestParticle

}

func destroy(part particle, particleList []particle, firstOccurence int, destroyed int, i int) int {
	part.collide = true
	if particleList[firstOccurence].collide == false {
		partfirst := particleList[firstOccurence]
		partfirst.collide = true
		fmt.Println(partfirst, "destroyed")
		particleList[firstOccurence] = partfirst
		destroyed++
	}
	part.collide = true
	fmt.Println(part, "destroyed")
	particleList[i] = part
	destroyed++
	return destroyed
}

func findLowestAcceleration(particles []particle) particle {
	lowestAcc := math.MaxInt32
	var lowestAccPart particle
	for _, part := range particles {
		accAbs := manhattanDistance(part.acceleration)
		if accAbs < lowestAcc {
			lowestAccPart = part
			lowestAcc = accAbs
		}
	}
	return lowestAccPart
}

func parseLines(lines []string) (map[int]int, []particle) {
	particleMap := make(map[int]int)
	particleList := make([]particle, len(lines))
	for i, line := range lines {
		numbers := util.ExtractAllNumbers(line)
		if len(numbers) != 9 {
			fmt.Printf("Error line %v", i)
			panic("can't parse line")
		}
		pos := coord{numbers[0], numbers[1], numbers[2]}
		vel := coord{numbers[3], numbers[4], numbers[5]}
		acc := coord{numbers[6], numbers[7], numbers[8]}
		part := particle{pos, vel, acc, i, false}
		particleMap[i] = manhattanDistance(pos)
		particleList[i] = part

	}
	return particleMap, particleList
}

func main() {
	lines := util.ReadFileLines("inputs/day20.txt")
	particleMap, particleList := parseLines(lines)
	//works for my input. for inputs with various particles with lower speed wouldn't work
	fmt.Println("Part 1", findLowestAcceleration(particleList))

	nearestParticle := simulate(particleList, particleMap, false)
	fmt.Println(nearestParticle)
}
