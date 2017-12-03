package main

import (
	"fmt"
	"util"
)

type coord struct {
	x int
	y int
}

func isInMap(set map[coord]int, x_ int, y_ int) bool{
	point := coord{x:x_,y:y_}
	_,exists :=set[point]
	return exists
}

func getManhattanDistance(point coord) int{
	return util.Abs(point.x) + util.Abs(point.y)
}

func getNeighborSum(currentCoord coord,coord2id map[coord]int) int{
	sum :=0
	x:= currentCoord.x
	y:= currentCoord.y
	for i := x-1 ; i <= x+1;i++{
		for j := y-1 ; j <= y+1;j++{
			value,_ :=coord2id[coord{i,j}]
			sum+=value
		}
	}

	return sum
}

func incrementPointValue(currentValue int, currentCoord coord,coord2id map[coord]int,part2 bool) int {
	if !part2{
		return currentValue+1
	}else{
		return getNeighborSum(currentCoord,coord2id)
	}
}
func fillSpiral(coord2id map[coord]int, id2coord map[int]coord, target int, part2 bool) int {
	//init first two points
	i := 1
	currentCoord := coord{x: 0, y: 0}
	coord2id[currentCoord] = i
	id2coord[i] =currentCoord
	currentCoord = coord{x: 1, y: 0}
	i=incrementPointValue(i,currentCoord,coord2id,part2)
	coord2id[currentCoord] = i
	id2coord[i] = currentCoord
	//with the two initial points the spiral can be created
	for ; i <= target; {
		//turn up
		if isInMap(coord2id, currentCoord.x-1, currentCoord.y) && !isInMap(coord2id, currentCoord.x, currentCoord.y+1) {
			currentCoord = coord{x: currentCoord.x, y: currentCoord.y + 1}
			//turn left
		} else if isInMap(coord2id, currentCoord.x, currentCoord.y-1) && !isInMap(coord2id, currentCoord.x-1, currentCoord.y) {
			currentCoord = coord{x: currentCoord.x - 1, y: currentCoord.y}
			//turn down
		} else if isInMap(coord2id, currentCoord.x+1, currentCoord.y) && !isInMap(coord2id, currentCoord.x, currentCoord.y-1) {
			currentCoord = coord{x: currentCoord.x, y: currentCoord.y - 1}
			//turn right
		} else if isInMap(coord2id, currentCoord.x, currentCoord.y+1) && !isInMap(coord2id, currentCoord.x+1, currentCoord.y) {
			currentCoord = coord{x: currentCoord.x + 1, y: currentCoord.y}
		}
		i = incrementPointValue(i,currentCoord,coord2id,part2)
		coord2id[currentCoord] = i
		id2coord[i] = currentCoord
	}
	return i
}


func main() {
	target := 347991
	coord2id := make(map[coord]int)
	id2coord := make(map[int]coord)
	fillSpiral(coord2id, id2coord, target,false)
	fmt.Println("Part1 :",getManhattanDistance(id2coord[target]))

	coord2id = make(map[coord]int)
	id2coord = make(map[int]coord)

	valueGreaterTarget := fillSpiral(coord2id, id2coord, target,true)
	fmt.Println("Part2 :",valueGreaterTarget)
}