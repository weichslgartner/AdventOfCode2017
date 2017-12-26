package main

import (
	"util"
	"regexp"
	"fmt"
	"strconv"
)

type Node struct {
	id            string
	depth         int
	weight        int
	subtreeWeight int
	predecessor   string
	successors    []string
}

func (n *Node) IsLeaf() bool {
	return len(n.successors) == 0
}

var deepestUnbalancedNodes []Node

func parseGraph(lines []string) map[string]Node {
	graph := make(map[string]Node, len(lines))
	regNodes := regexp.MustCompile("[a-z]+")
	reWeight := regexp.MustCompile("[0-9]+")
	for _, line := range lines {
		ids := regNodes.FindAllString(line, -1)
		currentID := ids[0]
		weight := reWeight.FindAllString(line, -1)
		intWeight, _ := strconv.Atoi(weight[0])
		currentNode := Node{currentID, 0, intWeight, 0, "", nil}
		if len(ids) > 1 {
			currentNode.successors = ids[1:]

		}
		graph[currentID] = currentNode
		//fmt.Println(currentNode)
	}
	addPredecessors(graph)
	return graph
}

func addPredecessors(graph map[string]Node) {
	for _, element := range graph {
		for _, successor := range element.successors {
			succ, exists := graph[successor]
			if exists {
				succ.predecessor = element.id
				graph[succ.id] = succ
			}
		}
	}
}

func findRoot(graph map[string]Node, start Node) Node {
	currentNode := graph[start.id]
	for currentNode.predecessor != "" {
		currentNode = graph[currentNode.predecessor]
	}
	return currentNode
}

func setSubtreeWeights(graph map[string]Node, node Node) {
	if node.IsLeaf() {
		node.subtreeWeight = node.weight
		graph[node.id] = node
		return
	}
	for _, succ := range node.successors {
		if graph[succ].subtreeWeight == 0 {
			setSubtreeWeights(graph, graph[succ])
		}
		node.subtreeWeight += graph[succ].subtreeWeight

	}
	node.subtreeWeight += node.weight
	graph[node.id] = node
}

func setDepth(graph map[string]Node, root Node, depth int) {
	root.depth = depth
	graph[root.id] = root
	for _, succ := range root.successors {
		setDepth(graph, graph[succ], depth+1)
	}
}

func findUnbalanceSubtree(graph map[string]Node, node Node) {
	balanced := isUnbalanced(node, graph)
	if !balanced {
		//fmt.Println("Unbalanced Node: ", node)
		next := findUnbalanceSuccessor(graph, node)
		addToDeepestUnbalancedSet(next)

		findUnbalanceSubtree(graph, next)
		//fmt.Println(" Subtree: ", node)

	}

}

func addToDeepestUnbalancedSet(next Node) {
	if len(deepestUnbalancedNodes) == 0 {
		deepestUnbalancedNodes = append(deepestUnbalancedNodes, next)
	} else if next.depth > deepestUnbalancedNodes[0].depth {
		deepestUnbalancedNodes = deepestUnbalancedNodes[:0]
		deepestUnbalancedNodes = append(deepestUnbalancedNodes, next)
	} else if next.depth == deepestUnbalancedNodes[0].depth {
		deepestUnbalancedNodes = append(deepestUnbalancedNodes, next)
	}

}

func findUnbalanceSuccessor(graph map[string]Node, node Node) Node {
	countSubTreeWeightsMap := make(map[int]int)
	var unbalancedNode Node
	unbalancedValue := 0
	for _, succ := range node.successors {
		countSubTreeWeightsMap[graph[succ].subtreeWeight] += 1
	}

	for key, value := range countSubTreeWeightsMap {
		if value == 1 {
			unbalancedValue = key
			break
		}
	}

	for _, succ := range node.successors {
		if graph[succ].subtreeWeight == unbalancedValue {
			unbalancedNode = graph[succ]
		}
	}
	return unbalancedNode
}

func isUnbalanced(node Node, graph map[string]Node) bool {
	balanced := true
	currentWeight := 0
	prevWeight := 0
	for i, succ := range node.successors {
		succNode := graph[succ]
		currentWeight = succNode.subtreeWeight
		if i > 0 && currentWeight != prevWeight {
			balanced = false
		}
		prevWeight = currentWeight

	}
	return balanced
}

func repair(graph map[string]Node, node Node) Node {
	pred := graph[node.id].predecessor
	difference := 0
	//find subtreeWeight of neighbor nodes
	for _, succ := range graph[pred].successors {
		if graph[succ].subtreeWeight != node.subtreeWeight {
			difference = graph[succ].subtreeWeight - node.subtreeWeight
		}
	}
	node.weight += difference
	node.subtreeWeight += difference
	return node
}

func main() {
	lines := util.ReadFileLines("inputs/day7.txt")
	deepestUnbalancedNodes = make([]Node, 0)
	graph := parseGraph(lines)

	var root Node
	for _, value := range graph {
		root = findRoot(graph, value)
		break
	}

	setDepth(graph, root, 0)
	fmt.Println("Root node's id is: ", root.id)
	setSubtreeWeights(graph, root)
	findUnbalanceSubtree(graph, root)
	addToDeepestUnbalancedSet(root)
	for _, node := range deepestUnbalancedNodes {
		repairedNode := repair(graph, node)
		fmt.Printf("Repaired Node %v has repaired weight %v.\n", repairedNode.id, repairedNode.weight)
		//graph[repairedNode.id] = repairedNode
	}

}
