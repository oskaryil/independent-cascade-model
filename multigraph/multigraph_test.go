package multigraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUndirectedMultiGraph(t *testing.T) {
	testGraph := NewUndirectedMultiGraph()

	assert.Equal(t, testGraph.lineCount, int64(0), "Line count should be 0")
	assert.Equal(t, testGraph.lineIDs, make([]int64, 0), "LineIDs slice should be empty")
}

func TestNewNode(t *testing.T) {
	testGraph := NewUndirectedMultiGraph()
	testNode := testGraph.NewNode(1)
	assert.Equal(t, testNode.id, int64(1), "Node id should match")
}

func TestNodeID(t *testing.T) {
	testGraph := NewUndirectedMultiGraph()
	testNode := testGraph.NewNode(1)
	assert.Equal(t, testNode.ID(), int64(1), "ID() should return the correct node id")
}
