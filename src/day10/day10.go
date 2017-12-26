package main

import (
	"util"
	"fmt"
	"io/ioutil"
)

/*
deals differently with negative numbers than the build in % operator
 */
func mod(number int, length int) int {
	if number < 0 {
		return mod(length+number, length)
	} else {
		return number % length
	}
}

func revertSubList(list []int, start int, stop int, length int) {
	copyList := make([]int, len(list))
	copy(copyList, list)
	for i := 0; i < length; i++ {
		index1 := mod(start+i, len(list))
		index2 := mod(stop-i, len(list))
		list[index1] = copyList[index2]
	}
}

func doTheHash(list []int, inputLengths []int, numberRounds int) {
	start := 0
	skipSize := 0
	for i := 0; i < numberRounds; i++ {
		for _, length := range inputLengths {
			startInvert := (start) % len(list)
			endInvert := (start + length - 1) % len(list)
			revertSubList(list, startInvert, endInvert, length)
			start += length + skipSize
			skipSize++
		}
	}

}

func generateList(length int) []int {
	intList := make([]int, length)
	for i := 0; i < length; i++ {
		intList[i] = i
	}
	return intList
}

func getInputLengths(filename string) []int {
	file, _ := ioutil.ReadFile(filename)
	lengths := append(file, 17, 31, 73, 47, 23)
	intLengths := make([]int, len(lengths))
	for i, element := range lengths {
		intLengths[i] = int(element)
	}
	return intLengths
}

func convertToDenseHash(list []int) []int {
	denseHash := make([]int, len(list)/16)
	for i := 0; i < len(list); i += 16 {
		for j := 0; j < 16; j++ {
			denseHash[i/16] ^= list[i+j]
		}
	}
	return denseHash
}

func convertToHex(list []int) string {
	hexString := ""
	for _, element := range list {
		hexString += fmt.Sprintf("%.02x", element)
	}
	return hexString
}

func part1() {
	inputLength := util.ReadNumbersFromFile("inputs/day10.txt")
	hashList := generateList(256)
	//fmt.Println(inputLength)
	//fmt.Println(hashList)
	doTheHash(hashList, inputLength, 1)
	//fmt.Println(hashList)
	fmt.Println("Result Part 1: ", hashList[0]*hashList[1])
}

func part2() {
	lengths := getInputLengths("inputs/day10.txt")
	hashList := generateList(256)
	doTheHash(hashList, lengths, 64)
	denseHash := convertToDenseHash(hashList)
	hex := convertToHex(denseHash)
	fmt.Println("Hash Part 2: ", hex)
}

func main() {
	part1()
	part2()
}
