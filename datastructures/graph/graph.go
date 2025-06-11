package graph

import "container/list"

// Graph implements a minimal graph structure of nodes connected by edges to adjacent nodes.
type Graph struct {
	nodeLookup map[int]*GraphNode
}

// AddNode saves the given node in the Graph's lookup table.
func (g *Graph) AddNode(node *GraphNode) {
	g.nodeLookup[node.id] = node
}

// GetNode returns the node reference for the given ID.
func (g *Graph) GetNode(id int) *GraphNode {
	return g.nodeLookup[id]
}

// AddEdge connects the nodes associated with the given IDs.
func (g *Graph) AddEdge(src int, dst int) {
	s := g.GetNode(src)
	d := g.GetNode(dst)
	s.MarkAdjacent(d)
}

// NewGraph instantiates a Graph and allocates its lookup table.
func NewGraph() *Graph {
	return &Graph{
		nodeLookup: make(map[int]*GraphNode, 0),
	}
}

// GraphNode represents a node with an ID and a linked list of references to connected nodes.
type GraphNode struct {
	id       int
	adjacent *list.List
}

// ID returns the ID of the node.
func (n *GraphNode) ID() int {
	return n.id
}

// MarkAdjacent adds the given node to the linked list of connected nodes.
func (n *GraphNode) MarkAdjacent(node *GraphNode) {
	n.adjacent.PushBack(node)
}

// GetAdjacent returns a reference to the linked list of connected nodes.
func (n *GraphNode) GetAdjacent() *list.List {
	return n.adjacent
}

// NewGraphNode instantiates a GraphNode and allocates its list of connected nodes.
func NewGraphNode(id int) *GraphNode {
	return &GraphNode{
		id:       id,
		adjacent: &list.List{},
	}
}
