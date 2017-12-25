package main

import (
	"util"
	"strings"
	"strconv"
	"regexp"
	"fmt"
	"sync"
	"math"
	"sort"
)

const ENTRY int =-1
const EXIT int = math.MaxInt32

var waitGroup sync.WaitGroup

type Instruction struct {
	name string
	op1  string
	op2  string
}

type BasicBlock struct {
	preConditions map[string]int
	postConditions map[string]int
	start int
	end int
	instructions []Instruction
	predecessor []int
	successor []int
}


func parseProgram(lines []string) []Instruction {
	program := make([]Instruction, len(lines))
	for i, line := range lines {
		tokens := strings.Split(line, " ")
		instruction := Instruction{tokens[0], tokens[1], ""}
		if len(tokens) > 2 {
			instruction.op2 = tokens[2]
		}
		program[i] = instruction
	}
	return program
}

func interpret(program []Instruction, registers map[string]int) {
	instructionPointer := 0
	numberMul := 0
	instructionCounter := 0

	for instructionPointer < len(program) {
		currentInstruction := program[instructionPointer]
		instructionCounter++
		//fmt.Print(currentInstruction.name, " ", currentInstruction.op1 , " ", currentInstruction.op2, " ")
		switch currentInstruction.name {
		case "set":
			registers[currentInstruction.op1] = getValue(currentInstruction.op2, registers)
		case "sub":
			registers[currentInstruction.op1] -= getValue(currentInstruction.op2, registers)
		case "mul":
			registers[currentInstruction.op1] *= getValue(currentInstruction.op2, registers)
			numberMul ++
		case "jnz":
			if getValue(currentInstruction.op1, registers) != 0 {
				fmt.Print(instructionPointer , " ")
				printRegisters(registers)
				instructionPointer += getValue(currentInstruction.op2, registers)

				continue
			}
		default:
			fmt.Printf("Unknown Command: %v", currentInstruction.name)
		}
		//fmt.Println(registers[currentInstruction.op1])

		instructionPointer++

	}
	fmt.Println("Number mul", numberMul)
}

func printRegisters(registers map[string]int){

	for _,char:= range "abcdefgh"{
		fmt.Print(string(char) , ": " ,registers[string(char)] , " ")
	}
	fmt.Println()
}

func getValue(name string, registers map[string]int) int {
	var value int
	if isNumber(name) {
		value, _ = strconv.Atoi(name)
	} else {
		value = registers[name]
	}
	return value
}

func isNumber(s string) bool {
	strNumbers := regexp.MustCompile("[0-9]+").FindAllString(s, -1)
	return len(strNumbers) == 1
}



func main() {
	lines := util.ReadFileLines("inputs/day23.txt")
	program := parseProgram(lines)
	registers := make(map[string]int)
	for _,char:= range "abcdefgh"{
		registers[string(char)] =0
	}

	//interpret(program,registers)

	findBlocks(program)


	registers["a"] = 1
	fmt.Println(registers["h"])
	//convertToC(program,registers)

}
func convertToC(program []Instruction, registers map[string]int) {
	instructionPointer := 0
	for key,value:= range registers{
		fmt.Printf("int %v = %v;\n",key,value)
	}
	for instructionPointer < len(program) {
		currentInstruction := program[instructionPointer]

		//fmt.Print(currentInstruction.name, " ", currentInstruction.op1 , " ", currentInstruction.op2, " ")
		label :=  strconv.Itoa(instructionPointer)
		switch currentInstruction.name {
		case "set":
			fmt.Printf("label%v: %v = %v;\n",label,currentInstruction.op1, currentInstruction.op2)
		case "sub":
			fmt.Printf("label%v: %v -= %v;\n",label,currentInstruction.op1, currentInstruction.op2)
		case "mul":
			fmt.Printf("label%v: %v *= %v;\n",label,currentInstruction.op1, currentInstruction.op2)
		case "jnz":
			jumpTarget := instructionPointer+ getValue(currentInstruction.op2,nil)
			if jumpTarget>=len(program){
				jumpTarget= math.MaxInt32
			}
			fmt.Printf("if (%v !=0){\n \tgoto label%v;\n}\n",currentInstruction.op1,jumpTarget)


		default:
			fmt.Printf("Unknown Command: %v", currentInstruction.name)
		}
		//fmt.Println(registers[currentInstruction.op1])

		instructionPointer++

	}
	fmt.Printf("label%v: printf(\"ENDE\\%i\",h);\n",math.MaxInt32)
}
func findBlocks(instructions []Instruction) {
	jumpTargets := make([]int,0)
	blockMap := make(map[int]BasicBlock,0)
	jumpTargets = append(jumpTargets,0)
//	blockInstructions := make([]Instruction,0)
//	currentBlockStart := 0
	for point,instruction := range instructions{
		if instruction.name == "jnz"{
			successor := point+getValue(instruction.op2,nil)
			jumpTargets = append(jumpTargets,successor)
			//jumpTargets = append(jumpTargets,point)
		//	if !isNumber(instruction.op1){
				jumpTargets = append(jumpTargets,point+1)
		//	}

		}
	}
	sort.Ints(jumpTargets)

	var currentBlock BasicBlock

	for i,target:=range jumpTargets{

			currentBlock = BasicBlock{}
			currentBlock.start = target
			if i < len(jumpTargets) -1 {
				currentBlock.end = jumpTargets[i+1]-1


			}else {
				currentBlock.end = len(instructions)
			}
			if currentBlock.end < len(instructions){


				successors := make([]int, 0)
				if instructions[currentBlock.end ].name == "jnz" {
					successor := currentBlock.end  + getValue(instructions[currentBlock.end ].op2, nil)
					/*
					if successor >= len(instructions) {
						successor = EXIT
					}
					*/
					successors = append(successors, successor)
					if !isNumber(instructions[currentBlock.end ].op1){
						successors = append(successors, currentBlock.end +1)
					}


				}
				if len(successors) == 0{
					successors = append(successors, currentBlock.end +1)
				}
				currentBlock.successor = successors
			}
			blockMap[currentBlock.start] = currentBlock

	}


/*
	for point,instruction := range instructions{
		blockInstructions = append(blockInstructions, instruction)
		if jumpTargets[point]{
			blockInstructtionsCopy := make([]Instruction,len(blockInstructions))
			if instruction.name == "jnz"{
				successor := point+getValue(instruction.op2,nil)
				if successor >=len(instructions){
					successor =EXIT
				}
				successors := make([]int,0)
				successors = append(successors, successor)
				if !isNumber(instruction.op1){
					successors = append(successors, point+1)
					jumpBlock := BasicBlock{}
					jumpBlock.start =point+1
					preC := make(map[string]int)
					preC[instruction.op1] = 0
					jumpBlock.preConditions = preC
					blockMap[point+1] = jumpBlock

				}
				block, exists :=  blockMap[currentBlockStart]
				if !exists{
					block = BasicBlock{nil,nil,currentBlockStart,point,blockInstructtionsCopy,nil,successors}
				}else{
					block.start, block.end, block.instructions, block.successor =currentBlockStart,point,blockInstructtionsCopy,successors
				}
				blockMap[currentBlockStart] = block
				blockInstructions = make([]Instruction,0)
				currentBlockStart = point+1
			}else{
				successors := make([]int,0)
				successors = append(successors, currentBlockStart-1)
				block, exists :=  blockMap[currentBlockStart]
				if !exists{
					block = BasicBlock{nil,nil,currentBlockStart,point-1,blockInstructtionsCopy,nil,successors}
				}else{
					block.start, block.end, block.instructions =currentBlockStart,point-1,blockInstructtionsCopy
				}
				blockMap[currentBlockStart] = block
				blockInstructions = make([]Instruction,0)
				currentBlockStart = point
			}







		}
	}
*/
/*
	for point,instruction := range instructions{
		fmt.Println(instruction)
		if jumpTargets[point] {
			fmt.Println("==============")
		}

	}
*/

	findPredecessors(blockMap)
	exitBlock := 0
	exitNode(blockMap, exitBlock)
	for _, block := range blockMap {
		fmt.Println(block)
	}

}
func exitNode(blockMap map[int]BasicBlock, exitBlock int) {
	for key, block := range blockMap {
		for _, successor := range block.successor {
			if successor == EXIT {
				exitBlock = key
			}

		}
	}
}
func findPredecessors(blockMap map[int]BasicBlock) {
	for key, block := range blockMap {
		for _, successor := range block.successor {
			if successor != EXIT{
				suc, exists := blockMap[successor]
				if exists{
					suc.predecessor = append(suc.predecessor, key)
					blockMap[successor] = suc
				}else{
					fmt.Printf("%v not in match \n", successor)
				}

			}

		}
	}
}
