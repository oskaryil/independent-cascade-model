package fileparser

import (
	"oskaryil/icm/multigraph"
	"testing"
)

func TestGenerateGraphFromFile(t *testing.T) {
	g := multigraph.NewUndirectedMultiGraph()

	fpath := "../android.csv"

	GenerateGraphFromFile(fpath, g)

	const expectedNodesLen int64 = 1128
	const expectedLineCount int64 = 12147

	if g.NodeCount() != expectedNodesLen {
		t.Errorf("Number of nodes are incorrent, actual: %d, expected: %d", g.NodeCount(), expectedNodesLen)
	}

	if g.LineCount() != expectedLineCount {
		t.Errorf("Number of lines are incorrent, actual: %d, expected: %d", g.LineCount(), expectedLineCount)
	}

}
