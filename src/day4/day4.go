package main

import (
	"strings"
	"fmt"
	"util"
	"sort"
)
func countValidPassphrases(lines []string, part1 bool)int{
	number:=0

	for _,line := range lines{
		strings.Replace(line, "\r","",-1)
		words := strings.Fields(line)
		valid := false
		if part1{
			valid =isValidLine(words)
		}else{
			valid =isValidLineAnagram(words)
		}


		if valid{
			number++
		}

	}
	return number
}
func isValidLine(words []string)bool {
	wordMap := make(map[string]bool)
	valid := true
	for _, word := range words {
		_, exists := wordMap[word]
		//not valid
		if exists {
			valid = false
			break
		}else{
			wordMap[word] = true
		}

	}
	return valid
}


func isValidLineAnagram(words []string)bool {
	wordMap := make(map[string]bool)
	valid := true
	for _,word := range words {
		letters := strings.Split(word, "")
		sort.Strings(letters)
		sortedLetters :=strings.Join(letters,"")
		_, exists := wordMap[sortedLetters]
		//not valid
		if exists {
			valid = false
			break
		}else{
			wordMap[sortedLetters] = true
		}

	}
	return valid
}

func main() {
	lines := util.ReadFileLines("inputs/day4.txt")
	number := countValidPassphrases(lines, true)
	fmt.Println("Part 1: ",number)
	number = countValidPassphrases(lines, false)
	fmt.Println("Part 2: ",number)
}
