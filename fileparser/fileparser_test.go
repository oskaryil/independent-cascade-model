package fileparser

import (
	"github.com/oskaryil/independent-cascade-model/multigraph"
	"testing"
)

func TestGenerateGraphFromFile(t *testing.T) {
	g := multigraph.NewUndirectedMultiGraph()

	fpath := "../android.csv"

	GenerateGraphFromFile(fpath, g)

	const expectedNodesLen int64 = 1128
	const expectedEdgeCount int64 = 12147

	if g.NodeCount() != expectedNodesLen {
		t.Errorf("Number of nodes are incorrent, actual: %d, expected: %d", g.NodeCount(), expectedNodesLen)
	}

	if g.EdgeCount() != expectedEdgeCount {
		t.Errorf("Number of edges are incorrent, actual: %d, expected: %d", g.EdgeCount(), expectedEdgeCount)
	}

}
