package main

import (
	"util"
	"strings"
	"strconv"
	"fmt"
)

type component struct {
	portA int
	portB int
}

func (c component) not(i int) int {
	if c.portA == i {
		return c.portB
	} else {
		return c.portA
	}
}

const MAXRECURSIONS = 1000000

var numberRecursions = 0
var bestBridge = 0
var longestBridge = 0
var strengthLongestBridge = 0

func removeFromList(comp component, componentList []component) []component {
LOOP:
	for i, c := range componentList {
		if (c.portA == comp.portA && c.portB == comp.portB) || (c.portA == comp.portB && c.portB == comp.portA) {
			if i < len(componentList)-1 {
				componentList = append(componentList[:i], componentList[i+1:]...)
				break LOOP
			} else {
				componentList = componentList[:i]
				break LOOP
			}

		}
	}
	return componentList
}

func parseInput(lines []string) []component {
	componentList := make([]component, 0)
	for _, line := range lines {
		tokens := strings.Split(line, "/")
		a, _ := strconv.Atoi(tokens[0])
		b, _ := strconv.Atoi(tokens[1])
		curComponent := component{a, b}
		componentList = append(componentList, curComponent)
	}
	return componentList
}

func buildChain(outPorts int, currentBride []component, availComponents []component) {
	if numberRecursions > MAXRECURSIONS {
		fmt.Println("To many recursions")
		return
	}
	numberRecursions++

	next := findComponents(availComponents, outPorts)
	if len(next) == 0 {
		strength := calculateStrength(currentBride)
		if strength > bestBridge {
			bestBridge = strength

		}
		length := len(currentBride)
		if (length > longestBridge) || (length == longestBridge && strength > strengthLongestBridge) {
			longestBridge = length
			strengthLongestBridge = strength
		}

		return
	}
	for _, comp := range next {
		componentListTemp := append([]component(nil), availComponents...)
		componentListTemp = removeFromList(comp, componentListTemp)
		currentBrideTemp := append([]component(nil), currentBride...)
		currentBrideTemp = append(currentBrideTemp, comp)
		buildChain(comp.not(outPorts), currentBrideTemp, componentListTemp)
	}

}

func calculateStrength(components []component) int {
	strength := 0
	for _, comp := range components {
		strength += comp.portB + comp.portA
	}
	return strength
}

func findComponents(componentList []component, port int) []component {
	curList := make([]component, 0)
	for _, comp := range componentList {
		if comp.portA == port || comp.portB == port {
			curList = append(curList, comp)
		}
	}
	return curList
}

func main() {
	lines := util.ReadFileLines("inputs/day24.txt")
	componentList := parseInput(lines)
	curList := findComponents(componentList, 0)
	for _, comp := range curList {
		currentBridge := make([]component, 0)
		componentListTemp := append([]component(nil), componentList...)
		componentListTemp = removeFromList(comp, componentListTemp)
		currentBridge = append(currentBridge, comp)
		buildChain(comp.not(0), currentBridge, componentListTemp)
	}
	fmt.Println("Part1: ", bestBridge)
	fmt.Println("Part2: ", strengthLongestBridge)
}
