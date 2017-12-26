package main

import (
	"util"
	"fmt"
)

func findCycles(memBanks []int) (int, int) {
	seenAllocations := make(map[string]int)
	steps := 0
	firstCycle := 0
	for seenAllocations[util.IntListToString(memBanks)] < 2 {
		if seenAllocations[util.IntListToString(memBanks)] == 1 && firstCycle == 0 {
			firstCycle = steps
		}
		seenAllocations[util.IntListToString(memBanks)] += 1
		argMax, max := util.ArgMax(memBanks)
		memBanks[argMax] = 0
		currentBank := (argMax + 1) % (len(memBanks))
		for i := max; i > 0; i-- {
			memBanks[currentBank] += 1
			currentBank = (currentBank + 1) % (len(memBanks))
		}
		steps++
	}
	return steps, firstCycle
}

func main() {
	memBanks := util.ReadNumbersFromFile("inputs/day6.txt")
	fmt.Println(memBanks)
	secondCycle, firstCycle := findCycles(memBanks)

	fmt.Println("Part 1: ",firstCycle)
	fmt.Println("Part 2: ",secondCycle - firstCycle)

}
