package main

import (
	"util"
	"fmt"
)



func findCycles(numbers []int) (int, int) {
	seenAllocations := make(map[string]int)
	cycleLength := 0
	firstCycle := 0
	for isInMap(seenAllocations, util.IntListToString(numbers)) == false || seenAllocations[util.IntListToString(numbers)] < 2 {
		if isInMap(seenAllocations, util.IntListToString(numbers)) && firstCycle == 0 {
			firstCycle = cycleLength
		}

		seenAllocations[util.IntListToString(numbers)] += 1
		argMax, max := util.ArgMax(numbers)
		numbers[argMax] = 0
		currentBank := (argMax + 1) % (len(numbers))
		for i := max; i > 0; i-- {
			numbers[currentBank] += 1
			currentBank = (currentBank + 1) % (len(numbers))
		}

		cycleLength++
	}
	return cycleLength, firstCycle
}
func isInMap(seenAllocations map[string]int, signature string) bool {
	_, exists := seenAllocations[signature]
	return exists
}


func main() {
	numbers := util.ReadNumbersFromFile("inputs/day6.txt")
	fmt.Println(numbers)
	secondCycle, firstCycle := findCycles(numbers)

	fmt.Println(firstCycle)
	fmt.Println(secondCycle - firstCycle)

}