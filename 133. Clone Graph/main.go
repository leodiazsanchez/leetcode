package main

func main() {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		if clonedNode, exists := visited[node]; exists {
			return clonedNode
		}

		clone := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}

		visited[node] = clone

		for _, neighbor := range node.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(neighbor))
		}

		return clone
	}

	return dfs(node)
}
