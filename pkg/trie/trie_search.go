package trie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pkg/browser"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// LoadSearch starts a webserver that exposes an endpoint search over backing trie datastructure
// with some baseline dictionary loaded
func LoadSearch() {
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
	trie := NewTrie()
	for k := range words {
		trie.Insert(k)
	}
	fmt.Println("finished inserting all words")

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		word := queryParams.Get("search")
		suggestions := trie.Search(word)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		b, err := json.Marshal(suggestions)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(b)
	})
	handler := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(router)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", handler))
		return
	}()

	fs := http.FileServer(http.Dir("pkg/trie/www"))
	http.Handle("/", fs)
	log.Println("Listening...")

	go func() {
		log.Fatal(http.ListenAndServe(":3000", nil))
		return
	}()

	err = browser.OpenURL("http://localhost:3000")
	if err != nil {
		log.Fatalf("failed to launch browser url:\n'%s'", err)
		return
	}

	select {} // block main thread
}
