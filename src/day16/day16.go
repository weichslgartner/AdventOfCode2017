package main

import (
	"strings"
	"util"
	"os"
	"fmt"
	"bufio"
)

func main() {
	//input := "s1,x3/4,pe/b"
	input := util.ReadFileToString("inputs/day16_optimized.txt")
	commands := strings.Split(input,",")
	programs := "abcdefghijklmnop" //"abcde"//
	optimize := false
	//programs := "abcde"//
	resultMap := make(map[string]int)
	programList := generateProgramList(programs)
	for i:= 0; i <100000;i++{


		if (optimize){
			commands = doTheDance(commands, &programList,false)
			if(i==2){
				f, _ := os.Create("inputs/day16_optimized.txt")
				w := bufio.NewWriter(f)
				w.WriteString(strings.Join(commands, ","))
				w.Flush()

			}
		}else{
			doTheDance(commands, &programList,false)
			//fmt.Println(strings.Join(programList, ""))
		}
		newString := strings.Join(programList, "")
		entry, exists := resultMap[newString]
		if exists {
			fmt.Printf("skip %v - %v \n", entry, i,)
			resultMap[newString] = i
		}
		if(i%1000 == 0){
			fmt.Printf("i: %v, length: %v \n", i, len(resultMap))
		}
	}
	fmt.Println()
	fmt.Println(strings.Join(programList, ""))
}

func isInvariantCycle(commands []string) bool{
	isInvariant := true
	for _,command := range commands{
		if command == "p"{
			isInvariant = false
			break
		}
	}
	return isInvariant
}

func doTheDance(commands []string, programList *[]string,profile bool) []string{
	resultMap := make(map[string]int)
	commandList := make([]string,0)
	newCommands :=  make([]string,0)
	lastIndex := 0
	//fmt.Println(commands)
	for i, command := range commands {
		switch string(command[0]) {
		case "s":
			numbers := util.ExtractAllNumbers(command)
			spin(programList, numbers[0])
			commandList = append(commandList, "s")
		case "x":
			numbers := util.ExtractAllNumbers(command)
			swapByPosition(programList, numbers[0], numbers[1])
			commandList = append(commandList, "x")
		case "p":
			swapByName(programList, string(command[1]), string(command[3]))
			commandList = append(commandList, "p")
		}

		if profile{
			newString :=strings.Join(*programList, "")
			entry, exists := resultMap[newString]
			if exists && isInvariantCycle(commandList[entry:i+1]){

			//	fmt.Printf("skip %v - %v: %v \n", entry, i, commandList[entry:i+1])
				newCommands = append(newCommands,commands[lastIndex:entry]...)
			//	fmt.Println("Cut out : ",commands[entry:i+1])
				lastIndex =i+1
				//commands = append(commands[:entry], commands[entry+i:]...)
			}
			resultMap[newString] = i
		}
	//	fmt.Println(strings.Join(*programList, ""))
	}
	if (len(newCommands) == 0){
		newCommands = commands
	}
	//fmt.Println(newCommands)
	return newCommands
	//fmt.Println(strings.Join(programList, ""))


}

func swapByName(programList *[]string, s string, s2 string){
	positionA, positionB := -1, -1
	for i,str := range *programList{
		if str == s{
			positionA = i
		}
		if str == s2{
			positionB = i
		}
		if positionA != -1 && positionB != -1{
			break
		}
	}
	if positionA == -1 && positionB == -1{
		os.Stderr.WriteString("Invalid letters")
	}

	swapByPosition(programList,positionA,positionB)
}

func generateProgramList(s string) []string {
	result := make([]string,len(s))
	for i,char := range s{
		result[i] = string(char)
	}
	return result
}
func swapByPosition(s *[]string, i int, i2 int) {
	(*s)[i] , (*s)[i2] = (*s)[i2] , (*s)[i]

}

func spin(s *[]string, numbers int){
	spin := numbers % len(*s)
	*s = append((*s)[len(*s)-spin:len(*s)],(*s)[0:len(*s)-spin]...)
}
