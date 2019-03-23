package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/meads/datastructures/pkg/trie"
)

func main() {
	// open file
	file, err := os.Open("pkg/trie/words.json")
	if err != nil {
		panic(fmt.Sprintf("error opening words.json\n%s", err))
	}
	fmt.Println("JSON file was read into memory successfully")
	defer file.Close()

	// read bytes
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("error reading from file handler\n%v", err))
	}

	// unmarshal into map
	words := make(map[string]int)
	err = json.Unmarshal(b, &words)
	if err != nil {
		panic(fmt.Sprintf("error unmarshalling words\n%v", err))
	}

	// create trie for holding the words and insert values
	trie := trie.NewTrie()
	for k := range words {
		trie.Insert(k)
	}
	fmt.Println("finished inserting all words")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		word := queryParams.Get("search")
		suggestions := trie.Search(word)

		fmt.Fprintf(w, "Here are the suggestions:\n%#v\n", suggestions)
	})
	http.ListenAndServe(":8080", nil)
}
