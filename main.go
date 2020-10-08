package main

import (
	"bufio"
	"fmt"
	"os"
	"oskaryil/icm/multigraph"
	"strconv"
	"strings"
	"time"
)

const (
	fname      string = "android.csv"
	dsname     string = "android"
	timeLayout        = "2006-01-02 15:04:05"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

func main() {

	// Read the lines from file
	lines, err := scanLines(fname)
	check(err)

	// Setup the graph
	g := multigraph.NewUndirecctedMultiGraph()

	// Parse the lines
	for _, line := range lines {
		spaceSplit := strings.Split(line, " ")
		nodeUId, _ := strconv.Atoi(spaceSplit[0])
		nodeVId, _ := strconv.Atoi(spaceSplit[1])
		bestCaseTimestamp := spaceSplit[4][11:] + " " + spaceSplit[5][:len(spaceSplit[5])-3]
		reviewId, _ := strconv.Atoi(spaceSplit[13][:len(spaceSplit[13])-1])
		parsedTimestamp, _ := time.Parse(timeLayout, bestCaseTimestamp)

		newNodeU := g.NewNode(int64(nodeUId))
		newNodeV := g.NewNode(int64(nodeVId))
		nodeU := g.AddNode(newNodeU)
		nodeV := g.AddNode(newNodeV)

		// fmt.Println(nodeU, nodeV)

		newLine := g.NewLine(nodeU, nodeV, int64(reviewId), parsedTimestamp)
		g.SetLine(newLine)

	}

	seed := make([]int64, 0)
	// seed = append(seed, 1101609)
	// seed = append(seed, 1478611)
	seed = append(seed, 1000205)

	// adjacentNodes, edges := g.AdjacentEdges(seed)
	// fmt.Println(adjacentNodes)
	// fmt.Println(len(adjacentNodes))
	// for _, val := range edges {
	// 	// fmt.Println("Key:", key)
	// 	// fmt.Println("Val",  val)
	// 	fmt.Println(val)
	// 	for k, _ := range val {
	// 		fmt.Println("adjacent node id: ", k)
	// 		for lineId, line := range val[k] {
	// 			fmt.Println(lineId, line.DiffusionTime().Before(time.Now().Add(24*time.Hour)))
	// 		}
	// 	}
	// }

	fmt.Println(g.NodesLen())
	informedNodes := g.DiffuseInformation(seed, 1.0, "best_case")
	cnt := 0
	for i, val := range informedNodes {
		fmt.Println(i, val)
		cnt++
	}
	fmt.Println(cnt)
	// fmt.Println(informedNodes)
	// fmt.Println(g.AdjacentNodes(n1.ID()))

	// graph := gogl.Spec().MultiGraph().Parallel().Undirected().Create().(gogl.DataGraph)
}
