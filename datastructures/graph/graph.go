package graph

import "container/list"

type Graph struct {
	nodeLookup map[int]*GraphNode
}

func (g *Graph) AddNode(node *GraphNode) {
	g.nodeLookup[node.id] = node
}

func (g *Graph) GetNode(id int) *GraphNode {
	return g.nodeLookup[id]
}

func (g *Graph) AddEdge(src int, dst int) {
	s := g.GetNode(src)
	d := g.GetNode(dst)
	s.MarkAdjacent(d)
}

func NewGraph() *Graph {
	return &Graph{
		nodeLookup: make(map[int]*GraphNode, 0),
	}
}

type GraphNode struct {
	id       int
	adjacent *list.List
}

func (n *GraphNode) ID() int {
	return n.id
}

func (n *GraphNode) MarkAdjacent(node *GraphNode) {
	n.adjacent.PushBack(node)
}

func (n *GraphNode) GetAdjacent() *list.List {
	return n.adjacent
}

func NewGraphNode(id int) *GraphNode {
	return &GraphNode{
		id:       id,
		adjacent: &list.List{},
	}
}
