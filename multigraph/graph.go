// Package multigraph provides a multigraph data structure with basic graph operations
package multigraph

import (
	"fmt"
	"sync"
	"time"
)

// Node data structure
type Node struct {
	id int64
}

// Edge is a data structure representing an edge in the graph
type Edge struct {
	id   int64
	from *Node
	to   *Node
	EdgeData
}

// EdgeData contains the necessary diffusion data
type EdgeData struct {
	reviewID        int64
	diffusionTime   time.Time
	diffusionNumber int64
}

// Graph is the graph data structure
type Graph struct {
	nodes map[int64]*Node
	edges map[int64]map[int64]map[int64]*Edge
	// edges map[]
	edgeIDs   []int64
	edgeCount int64
	lock      sync.RWMutex
}

// ID satisfies the Node interface
func (node *Node) ID() int64 {
	return int64(node.id)
}

// NewUndirectedMultiGraph returns an initialized empty Graph
func NewUndirectedMultiGraph() *Graph {
	return &Graph{
		nodes: make(map[int64]*Node),
		edges: make(map[int64]map[int64]map[int64]*Edge),

		edgeIDs:   make([]int64, 0),
		edgeCount: 0,
	}
}

// Node satisfies the Multigraph interface
func (g *Graph) Node(id int64) *Node {
	return g.nodes[id]
}

// AdjacentEdges returns a slice with the edges(edges) adjacent to each of the nodes in nodeIds
func (g *Graph) AdjacentEdges(nodeIds map[int64]time.Time) []*Edge {
	edgesMap := make([]map[int64]map[int64]*Edge, 0)
	for nodeID := range nodeIds {
		edgesMap = append(edgesMap, g.edges[nodeID])
	}

	edges := make([]*Edge, 0)
	existsMap := make(map[*Edge]bool)

	for i := range edgesMap {
		for j := range edgesMap[i] {
			for k := range edgesMap[i][j] {
				edge := edgesMap[i][j][k]
				// if edge.To().ID() < edge.From().ID() {
				// 	edge = edge.ReversedEdge()
				// }
				if _, exists := existsMap[edge]; !exists {
					edges = append(edges, edgesMap[i][j][k])
					existsMap[edge] = true
				}
				edges = append(edges, edgesMap[i][j][k])
			}
		}
	}
	return edges
}

// AdjacentEdgesSimple is used for simple graphs without timestamps, using numbers intead.
func (g *Graph) AdjacentEdgesSimple(nodeIds map[int64]int64) []*Edge {
	edgesMap := make([]map[int64]map[int64]*Edge, 0)

	for nodeID := range nodeIds {
		edgesMap = append(edgesMap, g.edges[nodeID])
	}

	edges := make([]*Edge, 0)
	existsMap := make(map[*Edge]bool)

	for i := range edgesMap {
		for j := range edgesMap[i] {
			for k := range edgesMap[i][j] {
				edge := edgesMap[i][j][k]
				// if edge.To().ID() < edge.From().ID() {
				// 	edge = edge.ReversedEdge()
				// }
				if _, exists := existsMap[edge]; !exists {
					edges = append(edges, edgesMap[i][j][k])
					fmt.Printf("From: %d, To: %d, dn: %d\n", edge.From().ID(), edge.To().ID(), edge.DiffusionNumber())
					existsMap[edge] = true
				}
			}
		}
	}
	return edges
}

// EdgeCount returns the number of edges(edges) in the graph
func (g *Graph) EdgeCount() int64 {
	return g.edgeCount
}

// NodeCount returns the number of nodes in the graph
func (g *Graph) NodeCount() int64 {
	return int64(len(g.nodes))
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(n *Node) *Node {
	g.lock.Lock()
	defer g.lock.Unlock()
	if _, exists := g.nodes[n.ID()]; exists {
		// panic(fmt.Sprintf("simple: node ID collision: %d", n.ID()))
		return g.nodes[n.ID()]
	}

	g.nodes[n.ID()] = n
	g.edges[n.ID()] = make(map[int64]map[int64]*Edge)
	return g.nodes[n.ID()]
}

// AddEdge adds an undirected edge to the graph between two nodes
// func (g *Graph) AddEdge(node1 *GraphNode, node2 *GraphNode) {
// 	g.lock.Lock()

// 	if g.edges == nil {
// 		g.edges = make(map[int64][]*GraphNode)
// 	}
// 	g.edges[node1.id] = append(g.edges[node1.id], node2)
// 	g.edges[node2.id] = append(g.edges[node2.id], node1)

// 	g.lock.Unlock()
// }

// From satisfies the Edge interface
func (edge Edge) From() *Node {
	return edge.from
}

// To satisfies the Edge interface
func (edge Edge) To() *Node {
	return edge.to
}

// ID satisfies the Edge interface
func (edge Edge) ID() int64 {
	return edge.id
}

// HasEdgeBetween returns whether an edge exists between nodes x and y.
func (g *Graph) HasEdgeBetween(xid, yid int64) bool {
	_, ok := g.edges[xid][yid]
	return ok
}

// EdgesBetween returns an array of edges between two nodes.
func (g *Graph) EdgesBetween(xid, yid int64) []*Edge {
	var edges []*Edge
	for _, l := range g.edges[xid][yid] {
		if l.From().ID() != xid {
			l = l.ReversedEdge()
		}
		edges = append(edges, l)
	}
	return edges
}

// ReversedEdge returns a new Edge with the F and T fields
// swapped. The UID of the new Edge is the same as the
// UID of the receiver. The Edges within the Edge are
// not altered.
func (edge *Edge) ReversedEdge() *Edge { edge.from, edge.to = edge.to, edge.from; return edge }

func (g *Graph) incrementEdgeCount() {
	g.edgeCount++
}

// NewEdge returns a new Edge from the source to the destination node.
// The returned Edge will have a graph-unique ID.
// The Edge's ID does not become valid in g until the Edge is added to g.
func (g *Graph) NewEdge(from, to *Node, reviewID int64, diffusionTime time.Time, diffusionNumber int64) *Edge {
	defer g.incrementEdgeCount()
	return &Edge{
		from: from,
		to:   to,
		EdgeData: EdgeData{
			reviewID:        reviewID,
			diffusionTime:   diffusionTime,
			diffusionNumber: diffusionNumber,
		},
		id: g.edgeCount,
	}
}

// DiffusionTime is a getter for the DiffusionTime of a edge
func (ld *EdgeData) DiffusionTime() time.Time {
	return ld.diffusionTime
}

// DiffusionNumber is a getter for the DiffusionTime of a edge
func (ld *EdgeData) DiffusionNumber() int64 {
	return ld.diffusionNumber
}

// GetEdgeData is a getter method for the EdgeData of a edge
func (g *Graph) GetEdgeData(uid, vid, edgeID int64) EdgeData {
	return g.edges[uid][vid][edgeID].EdgeData
}

// NewNode initializes a new node object with the provided id.
func (g *Graph) NewNode(id int64) *Node {
	// if len(g.nodes) == 0 {
	// 	return Node(0)
	// }
	return &Node{id: id}
}

// SetEdge adds l, a edge from one node to another. If the nodes do not exist, they are added and set to the nodes of the edge otherwise.
func (g *Graph) SetEdge(l *Edge) {
	g.lock.Lock()
	defer g.lock.Unlock()
	var (
		from = l.From()
		fid  = from.ID()
		to   = l.To()
		tid  = to.ID()
		lid  = l.ID()
	)

	if _, exists := g.nodes[fid]; !exists {
		g.AddNode(from)
	}

	if g.edges[fid][tid] == nil {
		g.edges[fid][tid] = make(map[int64]*Edge)
	}
	if _, exists := g.nodes[tid]; !exists {
		g.AddNode(to)
	}

	if g.edges[tid][fid] == nil {
		g.edges[tid][fid] = make(map[int64]*Edge)
	}

	g.edges[fid][tid][lid] = l
	g.edges[tid][fid][lid] = l
	g.edgeIDs = append(g.edgeIDs, lid)
}
