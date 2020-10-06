package multigraph

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/uid"
	"gonum.org/v1/gonum/graph/iterator"
	"sync"
	"time"
	"fmt"
)

type Timestamp time.Time

type GraphNode struct {
	id int64
	// neighbors []graph.Node
}

type Line struct {
	id int64
	from *GraphNode
	to *GraphNode
}

type EdgeData struct {
	reviewId int64
	diffusionTime Timestamp
}

func (node *GraphNode) String() string {
	return fmt.Sprintf("%v", node.id)
}

type Graph struct {
	Nodes []*GraphNode	
	nodes map[int64]graph.Node
	edges map[int64][]*GraphNode
	// map lineId->diffusionTime
	edgeData map[int64]*EdgeData
	lineIds uid.Set
	nEdges int64
	lock sync.RWMutex
}

// ID satisfies the Node interface
func (node *GraphNode) ID() int64 {
	return node.id
}

// NewGraphNode Returns a new graph node
func NewGraphNode(id int64) *GraphNode {
	return &GraphNode{id: id}
}

// Node satisfies the Multigraph interface
func (g *Graph) Node(id int64) *GraphNode {
	for _, node := range g.Nodes {
		if(node.id == id) {
			return node
		}
	}
	return nil
}

// AddNode implements the NodeAdder interface
func (g *Graph) AddNode(newNode *GraphNode) {
	if(g.Node(newNode.id) != nil) {
		panic("Node already exists in graph")
	}

	g.lock.Lock()

	g.Nodes = append(g.Nodes, newNode)

	g.lock.Unlock()
	
}

// AddEdge adds an undirected edge to the graph between two nodes
func (g *Graph) AddEdge(node1 *GraphNode, node2 *GraphNode) {
	g.lock.Lock()

	if g.edges == nil {
		g.edges = make(map[int64][]*GraphNode)
	}
	g.edges[node1.id] = append(g.edges[node1.id], node2)
	g.edges[node2.id] = append(g.edges[node2.id], node1)
	g.nEdges++;


	g.lock.Unlock()
}

// From satisfies the Line interface
func (line *Line) From() *GraphNode {
	return line.from
} 

// To satisfies the Line interface
func (line *Line) To() *GraphNode {
	return line.to
} 

// ID satisfies the Line interface
func (line *Line) ID() int64 {
	return line.id
}

// ReversedLine satisfies the Line interface
func (line *Line) ReversedLine() *Line {
	return &Line{from: line.to, to: line.from, id: line.id}
}

