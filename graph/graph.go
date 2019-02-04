package graph

//Node represents a node in a graph
type Node struct {
	Name     int
	Adjacent []*Edge
}

//Edge represents an edge between 2 nodes
type Edge struct {
	Start  *Node
	End    *Node
	Weight int
}

//Graph contains nodes and edges
type Graph struct {
	Nodes map[int]*Node
	Edges []*Edge
}

//CreateGraph creates a new graph from a 2d, nxn array of ints
//n nodes will be created, numbers from 0 to n-1.
//Each element i,j of the input is the distance from node i to node j.
//-1 is used to mark that there is no edge between 2 nodes.
func CreateGraph(inputData [][]int) Graph {
	nodes := make(map[int]*Node)

	for i := range inputData {
		newNode := Node{i, nil}
		nodes[i] = &newNode
	}
	var edges []*Edge
	for i, node := range inputData {
		for j, dist := range node {
			if dist != -1 {
				start := nodes[i]
				end := nodes[j]
				newEdge := Edge{start, end, dist}
				start.Adjacent = append(start.Adjacent, &newEdge)
				edges = append(edges, &newEdge)
			}
		}
	}

	return Graph{nodes, edges}
}
