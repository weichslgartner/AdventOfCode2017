package main

import (
	"util"
	"strings"
	"fmt"
	"time"
	"sync"
)

type rotation int

const (
	_0        = 1 + iota
	_90
	_180
	_270
	_0FLIPX
	_90FLIPX
	_180FLIPX
	_270FLIPX
)

var concurrent bool = false

type Pattern [][]bool

var waitGroup sync.WaitGroup

func cutAndMatch(currentPattern Pattern, y int, x int, stepSize int, newPattern *Pattern, patternMapWithRotations map[string]Pattern, innerY int, innerX int) {
	subPattern := cutOutPattern(currentPattern, y, x, stepSize)
	matchPattern(patternMapWithRotations, stepSize, subPattern, newPattern, innerY, innerX)
	if concurrent {
		waitGroup.Done()
	}

}

func determineStepAndInitNewPattern(currentPattern Pattern, ) (int, Pattern) {
	var stepSize int
	var newPattern Pattern
	if len(currentPattern)%2 == 0 {
		stepSize = 2
		newPattern = make([][]bool, (len(currentPattern)/2)*3)
	} else if len(currentPattern)%3 == 0 {
		stepSize = 3
		newPattern = make([][]bool, (len(currentPattern)/3)*4)
	} else {
		fmt.Println("Wrong Patternsize")
		return 0, nil
	}
	for y := 0; y < len(newPattern); y++ {
		newPattern[y] = make([]bool, len(newPattern))
	}
	return stepSize, newPattern
}

func matchPattern(patternMapWithRotations map[string]Pattern, stepSize int, subPattern Pattern, newPattern *Pattern, innerY int, innerX int) {
	for in, out := range patternMapWithRotations {
		inNew := string2Pattern(in)
		if len(inNew) != stepSize {
			continue
		}
		if match(subPattern, inNew) {
			copyIntoPattern(newPattern, out, innerY, innerX, stepSize+1)
			return
		}

	}

}

func transformToPatternMapWithRotations(patternMap map[string]Pattern, patternMapWithRotations map[string]Pattern) {
	for in, out := range patternMap {
		inNew := string2Pattern(in)
		var rot_input Pattern
		rot_input = make([][]bool, len(inNew))
		for i := 0; i < len(inNew); i++ {
			rot_input[i] = make([]bool, len(inNew))
		}
		for rot := _0; rot <= _270FLIPX; rot++ {
			copyIntoPattern(&rot_input, inNew, 0, 0, len(inNew))
			outpattern := rotate(rot_input, rot)
			patternMapWithRotations[pattern2String(outpattern)] = out
		}
	}
}
func copyIntoPattern(new *Pattern, src Pattern, y int, x int, step int) {
	for i := 0; i < step; i++ {
		for j := 0; j < step; j++ {
			(*new)[i+y][j+x] = src[i][j]
		}
	}
}

func rotate(pattern Pattern, rot int) Pattern {
	before := numberOn(pattern)
	switch rot {
	case _0:
	case _90:
		if len(pattern) == 2 {
			pattern[0][0], pattern[0][1], pattern[1][0], pattern[1][1] = pattern[1][0], pattern[0][0], pattern[1][1], pattern[0][1]
		} else {
			pattern[0][0], pattern[0][1], pattern[0][2], pattern[1][0], pattern[1][2], pattern[2][0], pattern[2][1], pattern[2][2] =
				pattern[2][0], pattern[1][0], pattern[0][0], pattern[2][1], pattern[0][1], pattern[2][2], pattern[1][2], pattern[0][2]

		}
	case _180:
		if len(pattern) == 2 {
			pattern[0][0], pattern[0][1], pattern[1][0], pattern[1][1] = pattern[1][1], pattern[1][0], pattern[0][1], pattern[0][0]
		} else {
			pattern[0][0], pattern[0][1], pattern[0][2], pattern[1][0], pattern[1][2], pattern[2][0], pattern[2][1], pattern[2][2] =
				pattern[2][2], pattern[2][1], pattern[2][0], pattern[1][2], pattern[1][0], pattern[0][2], pattern[0][1], pattern[0][0]

		}
	case _270:
		if len(pattern) == 2 {
			pattern[0][0], pattern[0][1], pattern[1][0], pattern[1][1] = pattern[0][1], pattern[1][1], pattern[0][0], pattern[1][0]
		} else {
			pattern[0][0], pattern[0][1], pattern[0][2], pattern[1][0], pattern[1][2], pattern[2][0], pattern[2][1], pattern[2][2] =
				pattern[0][2], pattern[1][2], pattern[2][2], pattern[0][1], pattern[2][1], pattern[0][0], pattern[1][0], pattern[2][0]

		}
	case _0FLIPX:
		if len(pattern) == 2 {
			pattern[0][0], pattern[0][1], pattern[1][0], pattern[1][1] = pattern[0][1], pattern[0][0], pattern[1][1], pattern[1][0]
		} else {
			pattern[0][0], pattern[0][1], pattern[0][2], pattern[1][0], pattern[1][2], pattern[2][0], pattern[2][1], pattern[2][2] =
				pattern[0][2], pattern[0][1], pattern[0][0], pattern[1][2], pattern[1][0], pattern[2][2], pattern[2][1], pattern[2][0]
		}
	case _90FLIPX:
		if len(pattern) == 2 {
			pattern[0][0], pattern[0][1], pattern[1][0], pattern[1][1] = pattern[0][0], pattern[1][0], pattern[0][1], pattern[1][1]
		} else {
			pattern[0][0], pattern[0][1], pattern[0][2], pattern[1][0], pattern[1][2], pattern[2][0], pattern[2][1], pattern[2][2] =
				pattern[0][0], pattern[1][0], pattern[2][0], pattern[0][1], pattern[2][1], pattern[0][2], pattern[1][2], pattern[2][2]

		}
	case _180FLIPX:
		if len(pattern) == 2 {
			pattern[0][0], pattern[0][1], pattern[1][0], pattern[1][1] = pattern[1][0], pattern[1][1], pattern[0][0], pattern[0][1]
		} else {
			pattern[0][0], pattern[0][1], pattern[0][2], pattern[1][0], pattern[1][2], pattern[2][0], pattern[2][1], pattern[2][2] =
				pattern[2][0], pattern[2][1], pattern[2][2], pattern[1][0], pattern[1][2], pattern[0][0], pattern[0][1], pattern[0][2]

		}
	case _270FLIPX:
		if len(pattern) == 2 {
			pattern[0][0], pattern[0][1], pattern[1][0], pattern[1][1] = pattern[1][1], pattern[0][1], pattern[1][0], pattern[0][0]
		} else {
			pattern[0][0], pattern[0][1], pattern[0][2], pattern[1][0], pattern[1][2], pattern[2][0], pattern[2][1], pattern[2][2] =
				pattern[2][2], pattern[1][2], pattern[0][2], pattern[2][1], pattern[0][1], pattern[2][0], pattern[1][0], pattern[0][0]

		}
	}
	after := numberOn(pattern)

	if before != after {
		fmt.Errorf("Rotationerror in %v", rot)
	}
	return pattern
}

func cutOutPattern(pattern Pattern, y int, x int, step int) Pattern {
	subPattern := make([][]bool, step)
	for i := 0; i < step; i++ {
		subPattern[i] = pattern[y+i][x:x+step]
	}
	return subPattern
}

func match(subPattern Pattern, inPattern Pattern) bool {
	if len(subPattern) != len(inPattern) {
		return false
	}
	for y, line := range subPattern {
		for x, _ := range line {
			if subPattern[y][x] != inPattern[y][x] {
				return false
			}
		}
	}
	return true
}

func numberOn(pattern Pattern) int {
	numberOn := 0
	for _, line := range pattern {
		for _, element := range line {
			if element {
				numberOn++
			}
		}
	}
	return numberOn
}

func printPattern(pattern Pattern) {
	for _, line := range pattern {
		for _, element := range line {
			if element {
				fmt.Println("#")
			} else {
				fmt.Println(".")
			}
		}
	}
	fmt.Println()
}

func parsePatternMap(lines []string) map[string]Pattern {
	patternMap := make(map[string]Pattern)
	for _, line := range lines {
		line := strings.Replace(line, " ", "", -1)
		patternz := strings.Split(line, "=>")
		var inputPattern, outputPattern Pattern
		for i, p := range patternz {
			p_lines := strings.Split(p, "/")
			if i == 0 {
				inputPattern = line2Pattern(p_lines)
			} else {
				outputPattern = line2Pattern(p_lines)
			}

		}
		patternMap[pattern2String(inputPattern)] = outputPattern
	}
	return patternMap
}

func string2Pattern(s string) Pattern {
	width := 2
	if len(s)%3 == 0 {
		width = 3
	}
	lines := make([]string, width)
	for i := 0; i < len(s); i += width {
		lines[i/width] = s[i:i+width]
	}
	return line2Pattern(lines)
}

func line2Pattern(p_lines []string) Pattern {
	pattern := make([][]bool, len(p_lines[0]))
	for i, _ := range pattern {
		pattern[i] = make([]bool, len(p_lines[0]))
	}
	for y, p_line := range p_lines {
		chars := strings.Split(p_line, "")
		for x, char := range chars {
			pattern[y][x] = string(char) == "#"
		}
	}
	return pattern
}

func pattern2String(pattern Pattern) string {
	result := ""
	for _, line := range pattern {
		for _, element := range line {
			if element {
				result += "#"
			} else {
				result += "."
			}
		}
	}
	return result
}

func main() {
	lines := util.ReadFileLines("inputs/day21.txt")
	patternMap := parsePatternMap(lines)
	patternMapWithRotations := make(map[string]Pattern)
	iterations := 18
	transformToPatternMapWithRotations(patternMap, patternMapWithRotations)
	currentPattern := line2Pattern([]string{".#.", "..#", "###"})
	for i := 1; i <= iterations; i++ {
		start := time.Now()
		stepSize, newPattern := determineStepAndInitNewPattern(currentPattern)
		innerY := 0

		for y := 0; y < len(currentPattern); y += stepSize {
			innerX := 0
			for x := 0; x < len(currentPattern); x += stepSize {
				if concurrent {
					waitGroup.Add(1)
					go cutAndMatch(currentPattern, y, x, stepSize, &newPattern, patternMapWithRotations, innerY, innerX)
				} else {
					cutAndMatch(currentPattern, y, x, stepSize, &newPattern, patternMapWithRotations, innerY, innerX)
				}

				innerX += stepSize + 1
			}
			innerY += stepSize + 1
		}
		if concurrent {
			waitGroup.Wait()
		}

		currentPattern = newPattern
		elapsed := time.Since(start)
		fmt.Printf("Iteration: %v, pixels on: %v  length grid %v; time %v\n", i, numberOn(currentPattern), len(currentPattern), elapsed)
	}
	fmt.Printf("Iteration: %v, pixels on: %v \n", iterations, numberOn(currentPattern))
}
