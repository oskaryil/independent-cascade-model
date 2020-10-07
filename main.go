package main

import (
	"oskaryil/icm/multigraph"
	// "github.com/sdboyer/gogl"
	"fmt"
	"time"
	// "os"
	// "bufio"
	// "log"
	// "encoding/csv"
	"os"
	"bufio"
	"strings"
	"strconv"
)

const (
	fname string = "android.csv"
	dsname string = "android"
	timeLayout = "2006-01-02 15:04:05"
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

	var lines[] string

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
		g.AddNode(newNodeU)
		g.AddNode(newNodeV)

		newLine := g.NewLine(newNodeU, newNodeV, int64(reviewId), multigraph.Timestamp(parsedTimestamp))
		g.SetLine(newLine)

		fmt.Println(bestCaseTimestamp)
		fmt.Println(reviewId)
	}

	// n1 := g.NewNode(1)
	// n2 := g.NewNode(2)
	// g.AddNode(n1)
	// g.AddNode(n2)
	// l1 := g.NewLine(n1, n2)
	// l2 := g.NewLine(n2, n1, 1233, multigraph.Timestamp(time.Now()))
	// g.SetLine(l1)
	// g.SetLine(l2)
	// gx := g.Nodes[0]
	// for _, line := range g.Lines(n1.ID(), n2.ID()).lines {
	// 	fmt.Println(line)	
	// }
	// g.Lines(n1.ID(), n2.ID()).lines
	// it := g.Lines(n1.ID(), n2.ID())
	// it.Next()
	// t, ok := g.(multigraph.Graph)
	// fmt.Println(it.Line().LineData)

	fmt.Println(g.LinesBetween(n1.ID(), n2.ID())[0].LineData.DiffusionTime())

	fmt.Println(g.HasEdgeBetween(n1.ID(), n2.ID()))

	adjacentNodes := g.AdjacentNodes(n1.ID())
	for key, val := range adjacentNodes {
		fmt.Println("Key:", key)
		fmt.Println("Val",  val)
	}
	fmt.Println(g.AdjacentNodes(n1.ID()))

	// graph := gogl.Spec().MultiGraph().Parallel().Undirected().Create().(gogl.DataGraph)
}