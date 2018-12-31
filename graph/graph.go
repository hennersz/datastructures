package graph

//Node represents a node in a graph
type Node struct {
	Name     int
	Adjacent []Edge
}

//Edge represents an edge between 2 nodes
type Edge struct {
	Start  int
	End    int
	Weight int
}

//Graph contains nodes and edges
type Graph struct {
	Nodes []Node
	Edges []Edge
}

//CreateGraph creates a new graph from a 2d array of ints
func CreateGraph(inputData [][]int) Graph {
	var nodes []Node
	var edges []Edge
	for i, node := range inputData {
		var newAdjacentList []Edge
		for j, dist := range node {
			if dist != -1 {
				newEdge := Edge{i, j, dist}
				newAdjacentList = append(newAdjacentList, newEdge)
				edges = append(edges, newEdge)
			}
		}
		newNode := Node{i, newAdjacentList}
		nodes = append(nodes, newNode)
	}

	return Graph{nodes, edges}
}
