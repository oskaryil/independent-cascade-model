package main

import (
	"oskaryil/icm/fileparser"
	"oskaryil/icm/multigraph"
	"testing"
)

func BenchmarkAndroid(b *testing.B) {
	for n := 0; n < b.N; n++ {

		g := multigraph.NewUndirectedMultiGraph()

		const fname string = "android.csv"

		fileparser.GenerateGraphFromFile(fname, g)
		seed := make([]int64, 0)
		seed = append(seed, 1000205)
		g.DiffuseInformation(seed, "best_case")
	}
}
