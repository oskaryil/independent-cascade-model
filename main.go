package main

import (
	"flag"
	"fmt"
	"oskaryil/icm/fileparser"
	"oskaryil/icm/multigraph"
)

func main() {

	// Setup the graph
	g := multigraph.NewUndirectedMultiGraph()

	var fname string

	flag.StringVar(&fname, "f", "", "Relative path to the input data file")

	flag.Parse()

	if len(fname) == 0 {
		panic("No input file specified")
	}

	fileparser.GenerateGraphFromFile(fname, g)

	seed := make([]int64, 0)
	seed = append(seed, 1000205)

	informedNodes := g.DiffuseInformation(seed, "best_case")
	cnt := 0
	const timeFormat = "2006-01-02 15:04:05"
	fmt.Printf("{")
	for i, val := range informedNodes {
		fmt.Printf("%d: '%v'\n", i, val.Format(timeFormat))
		cnt++
	}
	fmt.Printf("Number of nodes reached: %d \n", cnt)

}
