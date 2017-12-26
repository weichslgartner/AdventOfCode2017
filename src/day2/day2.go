package main

import (
	"util"
	"strings"
	"strconv"
	"math"
	"fmt"
)

func calcChecksum(lines []string) (int, int){
	checksumPart1 :=0
	checksumPart2 :=0
	for _,line := range lines{
		words := strings.Fields(line)
		checksumPart1 +=checksumLine(words)
		checksumPart2 +=checksumLine2(words)
	}
	return checksumPart1,checksumPart2
}

func checksumLine(words []string) int{
	min := math.MaxInt32
	max := math.MinInt32
	for _, word := range words {
		number, _ := strconv.Atoi(word)
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}

	}
	return max - min
}

func checksumLine2(words []string) int{
	checksum :=0
	numbers :=make([]int, len(words))
	for index, word := range words {
		numbers[index],_ = strconv.Atoi(word)
		for i:=0; i < index; i++{
			numerator := numbers[i]
			denominator := numbers[index]
			if denominator > numerator{
				numerator, denominator = denominator, numerator
			}
			if numerator % denominator == 0{
				checksum += numerator / denominator
				return checksum
			}

		}


	}
	//error case no even dividable numbers were found
	return checksum
}



func main() {
	fileStr := util.ReadFileToString("inputs/day2.txt")
	lines := strings.Split(fileStr,"\n")
	part1, part2 := calcChecksum(lines)
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}
