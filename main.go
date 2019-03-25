package main

import (
	"flag"
	"fmt"

	"github.com/meads/datastructures/pkg/trie"
)

func main() {
	example := flag.String("example", "", "run the trie search example with interactive browser")
	flag.Parse()

	switch *example {
	case "trie":
		fmt.Println("loading trie search example.")
		trie.LoadSearch()
	default:
		fmt.Printf("please specify example to run e.g. \n\ngo run main.go -example trie\n")
	}

	fmt.Println("exiting...")
}
