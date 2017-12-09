package main

import (
	"util"
	"regexp"
	"fmt"
	"strconv"
	"math"
)


type Condition struct {
	operandLeft  string
	operandRight int
	operand      string
}

type Instruction struct {
	register string
	inst     string
	value    int
}

func parseInstructionsAndEvaluate(lines []string) (map[string]int, int) {
	registerMap := make(map[string]int, 0)
	re, _ := regexp.Compile(`(\w+)\s+(\w+)\s(-*[0-9]+) if (\w+)\s+([><=!]+)\s+(-*[0-9]+)`)
	highestValue :=math.MinInt32
	for _, line := range lines {
		condition, instruction := parseLine(re, line, registerMap)
		highestValue = eval(registerMap, instruction, condition, highestValue)

	}

	//fmt.Println(registerMap)
	return registerMap, highestValue
}
func parseLine(re *regexp.Regexp, line string, registerMap map[string]int) (Condition, Instruction) {
	result := re.FindStringSubmatch(line)
	rightOperandCondition, _ := strconv.Atoi(result[6])
	condition := Condition{result[4], rightOperandCondition, result[5]}
	value, _ := strconv.Atoi(result[3])
	instruction := Instruction{result[1], result[2], value}
	_, exists := registerMap[result[1]]
	if !exists {
		registerMap[result[1]] = 0
	}
	return condition, instruction
}


func eval(registers map[string]int, instruction Instruction, condition Condition, highestValue int) int{
	leftSide := registers[condition.operandLeft]
	rightSide := condition.operandRight
	conditionResult := false
	switch condition.operand {
	case "==":
		conditionResult = leftSide == rightSide
	case ">=":
		conditionResult = leftSide >= rightSide
	case ">":
		conditionResult = leftSide > rightSide
	case "<":
		conditionResult = leftSide < rightSide
	case "<=":
		conditionResult = leftSide <= rightSide
	case "!=":
		conditionResult = leftSide != rightSide
	default:
		fmt.Println("Unknown Condition")
	}
	if (conditionResult) {
		currentRegister :=registers[instruction.register]
		switch instruction.inst{
		case "inc":
			currentRegister += instruction.value
		case "dec":
			currentRegister -= instruction.value
		default:
			fmt.Println("Unknown instruction")
		}
		registers[instruction.register] = currentRegister
		if currentRegister > highestValue{
			highestValue = currentRegister
		}
	}
	return highestValue

}

func main() {
	lines := util.ReadFileLines("inputs/day8.txt")
	registers, highestValue := parseInstructionsAndEvaluate(lines)
	maxValue :=math.MinInt32
	for _, value := range registers{
		if value > maxValue{
			maxValue = value
		}
	}
	fmt.Println("Largest register value in the end: ",maxValue)
	fmt.Println("Highest value during evaluation: ",highestValue)
}
