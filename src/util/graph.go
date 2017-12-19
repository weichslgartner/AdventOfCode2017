package util

type Node struct {
	id          string
	successors  []string
}

func (n Node) IsLeaf() bool {
	return len(n.successors) == 0
}


type Graph struct{
	graph map[string]Node
	numberNodes int
	root Node
}


func (g Graph) init(size int){
	g.graph = make(map[string]Node,size)
}

func (g Graph) addNode(n Node){
	g.graph[n.id] = n
}
