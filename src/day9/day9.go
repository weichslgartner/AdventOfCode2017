package main

import (
	"util"
	"fmt"
)

func getScore(input string) (int, int) {
	score := 0
	isGarbage := false
	groupsOpen := 0
	canceledCharacters := 0
	currentCharPos := 0
	for currentCharPos < len(input) {
		currentChar := string(input[currentCharPos])
		switch currentChar {
		case "!":
			currentCharPos ++ //ignores the next character
		case "<":
			if isGarbage {
				canceledCharacters++
			}
			isGarbage = true
		case ">":
			isGarbage = false
		case "{":
			if !isGarbage {
				groupsOpen++
			} else {
				canceledCharacters++
			}
		case "}":
			if !isGarbage {
				score += groupsOpen
				groupsOpen--
			} else {
				canceledCharacters++
			}
		default:
			if isGarbage {
				canceledCharacters++
			}

		}
		currentCharPos++
	}
	return score, canceledCharacters
}

func main() {
	input := util.ReadFileToString("inputs/day9.txt")
	score, canceledCharacters := getScore(input)
	fmt.Println("Score: ", score)
	fmt.Println("Canceled characters: ", canceledCharacters)
}
