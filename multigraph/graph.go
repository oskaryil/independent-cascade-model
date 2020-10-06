package multigraph

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/iterator"
	"sync"
	"time"
	"fmt"
)

type Timestamp time.Time

type Node int64

type Line struct {
	id int64
	from graph.Node
	to graph.Node
	LineData
}


// LineData contains the necessary diffusion data
type LineData struct {
	reviewID int64
	diffusionTime Timestamp
}

// func (node *GraphNode) String() string {
// 	return fmt.Sprintf("%v", node.id)
// }

type Graph struct {
	// Nodes []*GraphNode	
	nodes map[int64]graph.Node
	// edges map[int64][]*GraphNode
	lines map[int64]map[int64]map[int64]graph.Line
	// map lineId->diffusionTime
	// edgeData map[int64]*EdgeData
	lineIDs []int64
	nodeIDs []int64
	lock sync.RWMutex
}

// ID satisfies the Node interface
func (node Node) ID() int64 {
	return int64(node)
}

func NewUndirecctedMultiGraph() *Graph {
	return &Graph{
		nodes: make(map[int64]graph.Node),
		lines: make(map[int64]map[int64]map[int64]graph.Line),


		nodeIDs: make([]int64, 0),
		lineIDs: make([]int64, 0),
	}
}

// NewGraphNode Returns a new graph node
// func NewGraphNode(id int64) *GraphNode {
// 	return &GraphNode{id: id}
// }

// Node satisfies the Multigraph interface
func (g *Graph) Node(id int64) graph.Node {
	return g.nodes[id]
}

// AddNode implements the NodeAdder interface
func (g *Graph) AddNode(n graph.Node) {
	if _, exists := g.nodes[n.ID()+1]; exists {
		panic(fmt.Sprintf("simple: node ID collision: %d", n.ID()))
	}

	g.nodes[n.ID()] = n
	g.lines[n.ID()] = make(map[int64]map[int64]graph.Line)
	g.nodeIDs = append(g.nodeIDs, n.ID())
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
func (line Line) From() graph.Node {
	return line.from
} 

// To satisfies the Line interface
func (line Line) To() graph.Node {
	return line.to
} 

// ID satisfies the Line interface
func (line Line) ID() int64 {
	return line.id
}

// // ReversedLine satisfies the Line interface
// func (line *Line) ReversedLine() *Line {
// 	return &Line{from: line.to, to: line.from, id: line.id}
// }

// HasEdgeBetween returns whether an edge exists between nodes x and y.
func (g *Graph) HasEdgeBetween(xid, yid int64) bool {
	_, ok := g.lines[xid][yid]
	return ok
}

// LinesBetween returns the lines between nodes x and y.
func (g *Graph) LinesBetween(xid, yid int64) graph.Lines {
	if !g.HasEdgeBetween(xid, yid) {
		return graph.Empty
	}
	var lines []graph.Line
	for _, l := range g.lines[xid][yid] {
		if l.From().ID() != xid {
			l = l.ReversedLine()
		}
		lines = append(lines, l)
	}
	return iterator.NewOrderedLines(lines)
}


// Lines returns the lines from u to v if such an edge exists and nil otherwise.
// The node v must be directly reachable from u as defined by the From method.
func (g *Graph) Lines(uid, vid int64) graph.Lines {
	return g.LinesBetween(uid, vid)
}


// ReversedLine returns a new Line with the F and T fields
// swapped. The UID of the new Line is the same as the
// UID of the receiver. The Lines within the Edge are
// not altered.
func (line Line) ReversedLine() graph.Line { line.from, line.to = line.to, line.from; return line}


// NewLine returns a new Line from the source to the destination node.
// The returned Line will have a graph-unique ID.
// The Line's ID does not become valid in g until the Line is added to g.
func (g *Graph) NewLine(from, to graph.Node, reviewID int64, diffusionTime Timestamp) graph.Line {
	return &Line{
		from: from, 
		to: to, 
		LineData: LineData{
			reviewID: reviewID,
			diffusionTime: diffusionTime,
		}, 
		id: int64(len(g.lines)),
	}
}

func (ld *LineData) DiffusionTime() Timestamp {
	return ld.diffusionTime
}

// func (g *Graph) GetLineData(uid, vid, lineId int64) Timestamp {
// 	return g.lines[uid][vid][lineId]
// }

func (g *Graph) NewNode(id int64) graph.Node {
	// if len(g.nodes) == 0 {
	// 	return Node(0)
	// }
	return Node(id)
}

// SetLine adds l, a line from one node to another. If the nodes do not exist, they are added and set to the nodes of the line otherwise.
func (g *Graph) SetLine(l graph.Line) {
	var (
		from = l.From()
		fid = from.ID()
		to = l.To()
		tid = to.ID()
		lid = l.ID()
	)

	if _, ok := g.nodes[fid]; !ok {
		g.AddNode(from)
	} else {
		g.nodes[fid] = from
	}

	if g.lines[fid][tid] == nil {
		g.lines[fid][tid] = make(map[int64]graph.Line)
	}
	if _, ok := g.nodes[tid]; !ok {
		g.AddNode(to)
	} else {
		g.nodes[tid] = to
	}

	if g.lines[tid][fid] == nil {
		g.lines[tid][fid] = make(map[int64]graph.Line)
	}

	g.lines[fid][tid][lid] = l
	g.lines[tid][fid][lid] = l
	g.lineIDs = append(g.lineIDs, lid)
}

