package main

import (
	"fmt"

	"strconv"
)



func sumIdenticalDigits(fileStr string) int {
	sum := 0
	oldChar  := fileStr[len(fileStr)-1]
	for _, char := range fileStr {
		if oldChar == byte(char) {

			number, _ := strconv.Atoi(string(char))
			sum += number
		}

		oldChar = byte(char)
	}
	return sum
}

func halfWay(fileStr string) int {
	sum := 0
	length := len(fileStr)
	for index, char := range fileStr {
		if fileStr[(index+length/2)%length] == byte(char) {

			number, _ := strconv.Atoi(string(char))
			sum += number
		}


	}
	return sum
}

func main() {
	fileStr := readFileToString("inputs/day1.txt")

	sum := sumIdenticalDigits(fileStr)
	fmt.Printf("Part1 :%v\n",sum)

	sum = halfWay(fileStr)
	fmt.Printf("Part2 :%v\n",sum)
}