package main

import (
	"util"
	"fmt"
)

func findCycles(memBanks []int) (int, int) {
	seenAllocations := make(map[string]int)
	steps := 0
	firstCycle := 0
	for isInMap(seenAllocations, util.IntListToString(memBanks)) == false || seenAllocations[util.IntListToString(memBanks)] < 2 {
		if isInMap(seenAllocations, util.IntListToString(memBanks)) && firstCycle == 0 {
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
func isInMap(seenAllocations map[string]int, signature string) bool {
	_, exists := seenAllocations[signature]
	return exists
}

func main() {
	memBanks := util.ReadNumbersFromFile("inputs/day6.txt")
	fmt.Println(memBanks)
	secondCycle, firstCycle := findCycles(memBanks)

	fmt.Println(firstCycle)
	fmt.Println(secondCycle - firstCycle)

}
