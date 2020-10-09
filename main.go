package main

import (
	"fmt"
	"oskaryil/icm/fileparser"
	"oskaryil/icm/multigraph"
)

const (
	fname string = "android.csv"
)

func main() {

	// Setup the graph
	g := multigraph.NewUndirecctedMultiGraph()

	fileparser.GenerateGraphFromFile(fname, g)

	seed := make([]int64, 0)
	// seed = append(seed, 1101609)
	// seed = append(seed, 1478611)
	seed = append(seed, 1000205)
	// seed := make(map[int64]time.Time)

	// seed[1000205] = time.Time{}

	informedNodes := g.DiffuseInformation(seed, 1.0, "best_case")
	cnt := 0
	for i, val := range informedNodes {
		fmt.Println(i, val)
		cnt++
	}
	fmt.Println(cnt)

}
