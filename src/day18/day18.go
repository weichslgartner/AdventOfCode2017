package main

import (
	"util"
	"strings"
	"strconv"
	"regexp"
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

type Instruction struct {
	name string
	op1  string
	op2  string
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

func interpret(program []Instruction, registers map[string]int, part1 bool, id int, queue0 chan int, queue1 chan int) {
	instructionPointer := 0
	lastSound := 0
	numberSend := 0
	if !part1 {
		registers["p"] = id
		fmt.Printf("Program %v starting\n", id)
	}
	for instructionPointer < len(program) {
		currentInstruction := program[instructionPointer]
		switch currentInstruction.name {
		case "set":
			registers[currentInstruction.op1] = getValue(currentInstruction.op2, registers)
			instructionPointer++
		case "add":
			registers[currentInstruction.op1] += getValue(currentInstruction.op2, registers)
			instructionPointer++
		case "mul":
			registers[currentInstruction.op1] *= getValue(currentInstruction.op2, registers)
			instructionPointer++
		case "mod":
			registers[currentInstruction.op1] %= getValue(currentInstruction.op2, registers)
			instructionPointer++
		case "snd":
			//fmt.Println("Played sound with freq: %v ",registers[currentInstruction.op1])
			if part1 {
				lastSound = registers[currentInstruction.op1]
			} else {
				if id == 0 {
					queue0 <- getValue(currentInstruction.op1, registers)

				} else {
					queue1 <- getValue(currentInstruction.op1, registers)
				}
				//fmt.Printf("Program %v send %v\n", id, getValue(currentInstruction.op1,registers))
				numberSend++
			}

			instructionPointer++
		case "rcv":
			if part1 && registers[currentInstruction.op1] != 0 {
				fmt.Printf("Recover sound with freq: %v \n", lastSound)
				return
			} else {
				var receive int
				var timeout bool
				if id == 0 {
					receive, timeout = getFromChannelWithTimeOut(queue1)
				} else if id == 1 {
					receive, timeout = getFromChannelWithTimeOut(queue0)
				} else {
					fmt.Println("Unknown ID")
				}
				if timeout {
					fmt.Printf("Program %v send %v messages\n", id, numberSend)
					waitGroup.Done()
					return
				}
				registers[currentInstruction.op1] = receive
				//fmt.Printf("Program %v receives %v\n", id,receive)
			}
			instructionPointer++
		case "jgz":
			if getValue(currentInstruction.op1, registers) > 0 {
				instructionPointer += getValue(currentInstruction.op2, registers)
			} else {
				instructionPointer++
			}
		default:
			fmt.Printf("Unknown Command: %v", currentInstruction.name)
		}

	}
	if !part1 {
		fmt.Printf("Program %v send %v messages\n", id, numberSend)
	}
	waitGroup.Done()
}

func getFromChannelWithTimeOut(channel chan int) (int, bool) {
	var receive int
	select {
	case receive = <-channel:
		return receive, false
	case <-time.After(time.Second):
		fmt.Println("Timeout")
		return receive, true
	}

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
	lines := util.ReadFileLines("inputs/day18.txt")
	program := parseProgram(lines)
	registers := make(map[string]int)
	interpret(program,registers,true,0,nil, nil)
	queue0 := make(chan int, 100)
	queue1 := make(chan int, 100)
	registers0 := make(map[string]int)
	registers1 := make(map[string]int)
	//execute the two programs in parallel and wait for the deadlock
	waitGroup.Add(1)
	go interpret(program, registers0, false, 0, queue0, queue1)
	waitGroup.Add(1)
	go interpret(program, registers1, false, 1, queue0, queue1)
	waitGroup.Wait()
}
