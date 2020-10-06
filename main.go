package main

import (
	"oskaryil/icm/multigraph"
	"fmt"
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

	var g multigraph.Graph
	gn := multigraph.NewGraphNode(1)
	g.AddNode(gn)
	gx := g.Nodes[0]
	fmt.Println(gx.ID())
}