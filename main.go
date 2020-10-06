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
	// "io"
)

const (
	fname string = "android.csv"
	dsname string = "android"
)

func main() {

	// f, err := hdf5.OpenFile(fname, hdf5.F_ACC_RDONLY)
	// if err != nil {
	// 	panic(err)
	// }

	// group, err := f.OpenGroup(dsname)
	// if err != nil {
	// 	panic(err)
	// }

	// set := make([]s1Type, 100)
	// table, err := group.OpenAttribute("")
	// if err != nil {
	// 	panic(err)
	// }

	// // display the fields
	// fmt.Printf(":: data: %v\n", set)

	// // release resources
	// table.Close()
	// f.Close()


	// csvFile, _ := os.Open(fname)
	// reader := csv.NewReader(bufio.NewReader(csvFile))

	// for {
	// 	line, err := reader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(line)
	// }

	// var g multigraph.Graph
	g := multigraph.NewUndirecctedMultiGraph()
	n1 := g.NewNode(1)
	n2 := g.NewNode(2)
	g.AddNode(n1)
	g.AddNode(n2)
	// l1 := g.NewLine(n1, n2)
	l2 := g.NewLine(n2, n1, 1233, multigraph.Timestamp(time.Now()))
	// g.SetLine(l1)
	g.SetLine(l2)
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

	// graph := gogl.Spec().MultiGraph().Parallel().Undirected().Create().(gogl.DataGraph)
}