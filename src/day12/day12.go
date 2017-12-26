package main

import (
	"regexp"
	"util"
	"fmt"
)

type Node struct {
	id         string
	successors []string
}

func (n *Node) IsLeaf() bool {
	return len(n.successors) == 0
}

var visitedNodes map[string]bool

func parseGraph(lines []string) map[string]Node {
	graph := make(map[string]Node, len(lines))
	regNodes := regexp.MustCompile("[0-9]+")
	for _, line := range lines {
		ids := regNodes.FindAllString(line, -1)
		currentID := ids[0]
		currentNode := Node{currentID, nil}
		if len(ids) > 1 {
			currentNode.successors = ids[1:]
		}
		graph[currentID] = currentNode
		//fmt.Println(currentNode)
	}
	return graph
}

func findGroup(graph map[string]Node, currentNode Node) {
	if visitedNodes[currentNode.id] == true {
		return
	}
	visitedNodes[currentNode.id] = true
	if currentNode.IsLeaf() {
		return
	}
	for _, succ := range currentNode.successors {
		findGroup(graph, graph[succ])
	}
}

func part1(graph map[string]Node) {
	findGroup(graph, graph["0"])
	fmt.Println("Part 1 (size of group 0):", len(visitedNodes))
}

func part2(graph map[string]Node) {
	numberGroups := 0
	visitedNodes = make(map[string]bool)
	for key, node := range graph {
		if visitedNodes[key] == true {
			continue
		}
		findGroup(graph, node)
		numberGroups++
	}
	fmt.Println("Part 2 (number of groups):", numberGroups)
}

func main() {
	lines := util.ReadFileLines("inputs/day12.txt")
	visitedNodes = make(map[string]bool)
	graph := parseGraph(lines)
	part1(graph)
	part2(graph)
}
