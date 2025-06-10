package dfs

import "github.com/willroberts/ds-algo-review/datastructures/graph"

// DepthFirstSearch searches a Graph structure's edges recursively.
// Returns 'true' when a path exists between the given nodes.
func DepthFirstSearch(g *graph.Graph, src int, dst int) bool {
	return depthFirstSearch(g.GetNode(src), g.GetNode(dst), make(map[int]struct{}, 0))
}

func depthFirstSearch(src *graph.GraphNode, dst *graph.GraphNode, visited map[int]struct{}) bool {
	if _, ok := visited[src.ID()]; ok {
		return false
	}
	visited[src.ID()] = struct{}{}
	if src == dst {
		return true
	}
	for e := src.GetAdjacent().Front(); e != nil; e = e.Next() {
		if depthFirstSearch(e.Value.(*graph.GraphNode), dst, visited) {
			return true
		}
	}
	return false
}
