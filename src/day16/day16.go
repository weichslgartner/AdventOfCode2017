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
	input := util.ReadFileToString("inputs/day16.txt")
	commands := strings.Split(input,",")
	programs := "abcdefghijklmnop" //"abcde"//
	optimize := false
	//programs := "abcde"//
	resultMap := make(map[string]int)
	//programList := generateProgramList(programs)
	i:=0
	iterations := 1000
	cycledetection := true
	for ; i <iterations;{

		//tries to find patterns in the instructions which do not alter the input
		//currently broken;
		//use
		if (optimize){
			commands = doTheDance(commands, &programs,false)
			if(i==2){
				f, _ := os.Create("inputs/day16_optimized.txt")
				w := bufio.NewWriter(f)
				w.WriteString(strings.Join(commands, ","))
				w.Flush()
			}
		}else{
			doTheDance(commands, &programs,false)
			//fmt.Println(strings.Join(programList, ""))
		}
	//	newString := strings.Join(programList, "")
		entry, exists := resultMap[programs]
		if exists  && cycledetection{
			fmt.Printf("skip %v - %v \n", entry, i,)
			resultMap[programs] = i
			//forward
			resultMap = make(map[string]int)
			i=iterations- (iterations %(i-entry))+1
			fmt.Println(i)
		}else{
			resultMap[programs] = i
			i+=1
		}

		if(i%1000 == 0){
			fmt.Printf("i: %v, length: %v \n", i, len(resultMap))
		}
	}
	fmt.Println()
	fmt.Println(programs)
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

func doTheDance(commands []string, programList *string,profile bool) []string{
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
			//newString :=strings.Join(*programList, "")
			entry, exists := resultMap[*programList]
			if exists && isInvariantCycle(commandList[entry:i+1]){

			//	fmt.Printf("skip %v - %v: %v \n", entry, i, commandList[entry:i+1])
				newCommands = append(newCommands,commands[lastIndex:entry]...)
			//	fmt.Println("Cut out : ",commands[entry:i+1])
				lastIndex =i+1
				//commands = append(commands[:entry], commands[entry+i:]...)
			}
			resultMap[*programList] = i
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

func swapByName(programList *string, s string, s2 string){
	positionA, positionB := -1, -1
	for i,str := range *programList{
		if string(str) == s{
			positionA = i
		}
		if string(str) == s2{
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
func swapByPosition(s *string, i int, i2 int) {
	a := (*s)[i]
	b := (*s)[i2]
	replaceLetter(s,i,string(b))
	replaceLetter(s,i2,string(a))

}

func replaceLetter(s *string, position int, letter string){
	*s = (*s)[:position] + letter + (*s)[position+1:]
}

func spin(s *string, numbers int){
	spin := numbers % len(*s)
	*s = (*s)[len(*s)-spin:len(*s)] + (*s)[0:len(*s)-spin]
}
