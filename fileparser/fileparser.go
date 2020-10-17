package fileparser

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/oskaryil/independent-cascade-model/multigraph"
	"github.com/oskaryil/independent-cascade-model/utils"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

func scanLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var edges []string

	for scanner.Scan() {
		edges = append(edges, scanner.Text())
	}

	return edges, nil
}

// GenerateGraphFromFile takes a file path and a Graph pointer and returns a populated Graph
func GenerateGraphFromFile(fpath string, g *multigraph.Graph) *multigraph.Graph {
	// Read the edges from file
	edges, err := scanLines(fpath)
	utils.CheckError(err)

	// Parse the edges
	for _, edge := range edges {
		spaceSplit := strings.Split(edge, " ")
		nodeUID, _ := strconv.Atoi(spaceSplit[0])
		nodeVID, _ := strconv.Atoi(spaceSplit[1])
		bestCaseTimestamp := spaceSplit[4][11:] + " " + spaceSplit[5][:len(spaceSplit[5])-3]
		reviewID, _ := strconv.Atoi(spaceSplit[13][:len(spaceSplit[13])-1])
		parsedTimestamp, _ := time.Parse(timeLayout, bestCaseTimestamp)

		newNodeU := g.NewNode(int64(nodeUID))
		newNodeV := g.NewNode(int64(nodeVID))
		nodeU := g.AddNode(newNodeU)
		nodeV := g.AddNode(newNodeV)

		newEdge := g.NewEdge(nodeU, nodeV, int64(reviewID), parsedTimestamp, 0)
		// if nodeUId == 1045553 {
		// 	fmt.Println(newEdge.From().ID())
		// }
		g.SetEdge(newEdge)
	}
	return g
}
