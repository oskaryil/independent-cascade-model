package fileparser

import (
	"bufio"
	"os"
	"github.com/oskaryil/independent-cascade-model/multigraph"
	"github.com/oskaryil/independent-cascade-model/utils"
	"strconv"
	"strings"
	"time"
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

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

// GenerateGraphFromFile takes a file path and a Graph pointer and returns a populated Graph
func GenerateGraphFromFile(fpath string, g *multigraph.Graph) *multigraph.Graph {
	// Read the lines from file
	lines, err := scanLines(fpath)
	utils.CheckError(err)

	// Parse the lines
	for _, line := range lines {
		spaceSplit := strings.Split(line, " ")
		nodeUID, _ := strconv.Atoi(spaceSplit[0])
		nodeVID, _ := strconv.Atoi(spaceSplit[1])
		bestCaseTimestamp := spaceSplit[4][11:] + " " + spaceSplit[5][:len(spaceSplit[5])-3]
		reviewID, _ := strconv.Atoi(spaceSplit[13][:len(spaceSplit[13])-1])
		parsedTimestamp, _ := time.Parse(timeLayout, bestCaseTimestamp)

		newNodeU := g.NewNode(int64(nodeUID))
		newNodeV := g.NewNode(int64(nodeVID))
		nodeU := g.AddNode(newNodeU)
		nodeV := g.AddNode(newNodeV)

		newLine := g.NewLine(nodeU, nodeV, int64(reviewID), parsedTimestamp, 0)
		// if nodeUId == 1045553 {
		// 	fmt.Println(newLine.From().ID())
		// }
		g.SetLine(newLine)
	}
	return g
}
