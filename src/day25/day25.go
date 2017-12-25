package main

import (
	"util"
	"regexp"
	"fmt"
	"strconv"
)

const RIGHT = 1
const LEFT = -1

type instruction struct{
	write int
	dir int
	nextState string
}


func generateKey(state string, value int) string{
	return state+strconv.Itoa(value)
}



func countOnes(machine map[int]int)int{
	numberOnes :=0
	for _, element := range machine{
		numberOnes +=element
	}
	return numberOnes
}




func parseInstructions (lines []string) (string, int, map[string]instruction) {
	var startState string
	var checkSumSteps int
	var currentState string
	var currentValue int
	var currentInstruction instruction
	instructionMap := make(map[string]instruction)
	for _, line := range lines {

		re := regexp.MustCompile("Begin in state ([A-Z]).")
		match := re.FindStringSubmatch(line)
		if len(match) > 0 {
			startState = match[1]
		}
		re = regexp.MustCompile("Perform a diagnostic checksum after ([0-9]+) steps.")
		match = re.FindStringSubmatch(line)
		if len(match) > 0 {
			checkSumSteps, _ = strconv.Atoi(match[1])
		}
		re = regexp.MustCompile("In state ([A-Z]):")
		match = re.FindStringSubmatch(line)
		if len(match) > 0 {
			currentState = match[1]
		}

		re = regexp.MustCompile("If the current value is ([0-1]):")
		match = re.FindStringSubmatch(line)
		if len(match) > 0 {
			currentValue, _ = strconv.Atoi(match[1])
		}
		re = regexp.MustCompile("- Write the value ([0-1]).")
		match = re.FindStringSubmatch(line)
		if len(match) > 0 {
			value, _ := strconv.Atoi(match[1])
			currentInstruction = instruction{value, 0, ""}
		}

		re = regexp.MustCompile("- Move one slot to the ([a-z]*).")
		match = re.FindStringSubmatch(line)
		if len(match) > 0 {
			switch match[1] {
			case "left":
				currentInstruction.dir = LEFT
			case "right":
				currentInstruction.dir = RIGHT
			default:
				fmt.Errorf("unknown direction")
			}

		}

		re = regexp.MustCompile("- Continue with state ([A-Z]).")
		match = re.FindStringSubmatch(line)
		if len(match) > 0 {
			currentInstruction.nextState = match[1]
			instructionMap[generateKey(currentState, currentValue)] = currentInstruction

		}

	}
	return startState, checkSumSteps, instructionMap
}


func main() {
	lines := util.ReadFileLines("inputs/day25.txt")
	startState, checkSumSteps, instructionMap := parseInstructions (lines)
	fmt.Printf("Start state %v, checksum after %v\n",startState, checkSumSteps)
	//fmt.Println(instructionMap)
	currentPosition := 0
	currentState := startState
	machineMap := make(map[int]int)
	machineMap[0] = 0
	for i := 0; i < checkSumSteps; i++ {
		currentValue := machineMap[currentPosition]
		inst := instructionMap[generateKey(currentState, currentValue)]
		currentValue = inst.write
		machineMap[currentPosition] = currentValue
		currentState = inst.nextState
		currentPosition += inst.dir
	}
	ones := countOnes(machineMap)
	fmt.Printf("Number ones after %v steps: %v",checkSumSteps,ones)
}