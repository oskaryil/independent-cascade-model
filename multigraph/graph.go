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

// Line is a data structure representing an edge in the graph
type Line struct {
	id   int64
	from *Node
	to   *Node
	LineData
}

// LineData contains the necessary diffusion data
type LineData struct {
	reviewID        int64
	diffusionTime   time.Time
	diffusionNumber int64
}

// Graph is the graph data structure
type Graph struct {
	nodes map[int64]*Node
	lines map[int64]map[int64]map[int64]*Line
	// edges map[]
	lineIDs   []int64
	lineCount int64
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
		lines: make(map[int64]map[int64]map[int64]*Line),

		lineIDs:   make([]int64, 0),
		lineCount: 0,
	}
}

// Node satisfies the Multigraph interface
func (g *Graph) Node(id int64) *Node {
	return g.nodes[id]
}

// AdjacentEdges returns a slice with the lines(edges) adjacent to each of the nodes in nodeIds
func (g *Graph) AdjacentEdges(nodeIds map[int64]time.Time) []*Line {
	linesMap := make([]map[int64]map[int64]*Line, 0)
	for nodeID := range nodeIds {
		linesMap = append(linesMap, g.lines[nodeID])
	}

	lines := make([]*Line, 0)
	existsMap := make(map[*Line]bool)

	for i := range linesMap {
		for j := range linesMap[i] {
			for k := range linesMap[i][j] {
				line := linesMap[i][j][k]
				// if line.To().ID() < line.From().ID() {
				// 	line = line.ReversedLine()
				// }
				if _, exists := existsMap[line]; !exists {
					lines = append(lines, linesMap[i][j][k])
					existsMap[line] = true
				}
				lines = append(lines, linesMap[i][j][k])
			}
		}
	}
	return lines
}

// AdjacentEdgesSimple is used for simple graphs without timestamps, using numbers intead.
func (g *Graph) AdjacentEdgesSimple(nodeIds map[int64]int64) []*Line {
	linesMap := make([]map[int64]map[int64]*Line, 0)

	for nodeID := range nodeIds {
		linesMap = append(linesMap, g.lines[nodeID])
	}

	lines := make([]*Line, 0)
	existsMap := make(map[*Line]bool)

	for i := range linesMap {
		for j := range linesMap[i] {
			for k := range linesMap[i][j] {
				line := linesMap[i][j][k]
				// if line.To().ID() < line.From().ID() {
				// 	line = line.ReversedLine()
				// }
				if _, exists := existsMap[line]; !exists {
					lines = append(lines, linesMap[i][j][k])
					fmt.Printf("From: %d, To: %d, dn: %d\n", line.From().ID(), line.To().ID(), line.DiffusionNumber())
					existsMap[line] = true
				}
			}
		}
	}
	return lines
}

// LineCount returns the number of lines(edges) in the graph
func (g *Graph) LineCount() int64 {
	return g.lineCount
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
	g.lines[n.ID()] = make(map[int64]map[int64]*Line)
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

// From satisfies the Line interface
func (line Line) From() *Node {
	return line.from
}

// To satisfies the Line interface
func (line Line) To() *Node {
	return line.to
}

// ID satisfies the Line interface
func (line Line) ID() int64 {
	return line.id
}

// HasEdgeBetween returns whether an edge exists between nodes x and y.
func (g *Graph) HasEdgeBetween(xid, yid int64) bool {
	_, ok := g.lines[xid][yid]
	return ok
}

// LinesBetween returns an array of lines between two nodes.
func (g *Graph) LinesBetween(xid, yid int64) []*Line {
	var lines []*Line
	for _, l := range g.lines[xid][yid] {
		if l.From().ID() != xid {
			l = l.ReversedLine()
		}
		lines = append(lines, l)
	}
	return lines
}

// ReversedLine returns a new Line with the F and T fields
// swapped. The UID of the new Line is the same as the
// UID of the receiver. The Lines within the Edge are
// not altered.
func (line *Line) ReversedLine() *Line { line.from, line.to = line.to, line.from; return line }

func (g *Graph) incrementLineCount() {
	g.lineCount++
}

// NewLine returns a new Line from the source to the destination node.
// The returned Line will have a graph-unique ID.
// The Line's ID does not become valid in g until the Line is added to g.
func (g *Graph) NewLine(from, to *Node, reviewID int64, diffusionTime time.Time, diffusionNumber int64) *Line {
	defer g.incrementLineCount()
	return &Line{
		from: from,
		to:   to,
		LineData: LineData{
			reviewID:        reviewID,
			diffusionTime:   diffusionTime,
			diffusionNumber: diffusionNumber,
		},
		id: g.lineCount,
	}
}

// DiffusionTime is a getter for the DiffusionTime of a line
func (ld *LineData) DiffusionTime() time.Time {
	return ld.diffusionTime
}

// DiffusionNumber is a getter for the DiffusionTime of a line
func (ld *LineData) DiffusionNumber() int64 {
	return ld.diffusionNumber
}

// GetLineData is a getter method for the LineData of a line
func (g *Graph) GetLineData(uid, vid, lineID int64) LineData {
	return g.lines[uid][vid][lineID].LineData
}

// NewNode initializes a new node object with the provided id.
func (g *Graph) NewNode(id int64) *Node {
	// if len(g.nodes) == 0 {
	// 	return Node(0)
	// }
	return &Node{id: id}
}

// SetLine adds l, a line from one node to another. If the nodes do not exist, they are added and set to the nodes of the line otherwise.
func (g *Graph) SetLine(l *Line) {
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

	if g.lines[fid][tid] == nil {
		g.lines[fid][tid] = make(map[int64]*Line)
	}
	if _, exists := g.nodes[tid]; !exists {
		g.AddNode(to)
	}

	if g.lines[tid][fid] == nil {
		g.lines[tid][fid] = make(map[int64]*Line)
	}

	g.lines[fid][tid][lid] = l
	g.lines[tid][fid][lid] = l
	g.lineIDs = append(g.lineIDs, lid)
}
