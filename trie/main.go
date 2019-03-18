package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// TrieNode represents a node a Trie data structure
type TrieNode struct {
	Children        map[string]*TrieNode
	Val             string
	CompletesString bool
}

// NewTrieNode creates an instance of TrieNode with the supplied letter for its' value
func NewTrieNode(letter string) *TrieNode {
	return &TrieNode{
		Children: make(map[string]*TrieNode),
		Val:      letter,
	}
}

// Trie is a Tree like data structure which associates a prefix string with "branches" of suffixes
type Trie struct {
	RootNode *TrieNode
}

// NewTrie creates an instance of Trie with a root node
func NewTrie() *Trie {
	return &Trie{
		RootNode: NewTrieNode(""),
	}
}

// Insert adds a word in the Trie structure
func (t *Trie) Insert(word string) {
	node := t.RootNode
	letters := strings.Split(word, "")
	for i := 0; i < len(letters); i++ {
		currentLetter := letters[i]
		if v, ok := node.Children[currentLetter]; ok {
			node = v
		} else {
			newNode := NewTrieNode(currentLetter)
			node.Children[currentLetter] = newNode
			node = newNode
		}
	}
	node.CompletesString = true
}

// Find returns a boolean indicating that the word exists in the Trie
func (t *Trie) Find(word string) bool {
	node := t.RootNode
	letters := strings.Split(word, "")
	for i := 0; i < len(letters); i++ {
		currentLetter := letters[i]
		if v, ok := node.Children[currentLetter]; ok {
			node = v
		} else {
			return false
		}
	}

	return true
}

// String prints the string value of the Trie
func (t *Trie) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

// Remove returns a string value indicating the result of attempting to remove an entire word from the Trie structure
func (t *Trie) Remove(word string) string {
	node := t.RootNode
	suffixes := []*TrieNode{}

	letters := strings.Split(word, "")
	// case where no part of 'word' can be removed from trie
	for i := 0; i < len(letters); i++ {
		currentLetter := letters[i]
		if v, ok := node.Children[currentLetter]; ok {
			node = v
			suffixes = append([]*TrieNode{node}, suffixes...)
			if i == len(letters) && len(node.Children) > 0 {
				panic(fmt.Sprintf("suffixes in trie depend on %s", word))
			}
		}
	}

	// for case where some parts of 'word' can be removed from trie
	for j := 1; j < len(suffixes); j++ {
		parent := suffixes[j] // first element which was the last visited in above loop

		childLetter := string(word[len(suffixes)-j]) // last character in the string "word"

		if childNode, exists := parent.Children[childLetter]; exists {
			if len(childNode.Children) == 0 {
				delete(parent.Children, childLetter)
			} else {
				return fmt.Sprintf("word '%s' cannot be removed because there are dependent suffixes", word)
			}
		}

		if parent.CompletesString || len(parent.Children) > 0 {
			return fmt.Sprintf("some suffixes of '%s', were removed from trie", word)
		}
	}

	// for case where all parts of 'word' can be removed from trie
	delete(t.RootNode.Children, string(word[0]))

	return fmt.Sprintf("removed '%s'; no other '%s'=words remain", word, string(word[0]))
}

func main() {
	trie := NewTrie()
	trie.Insert("listen")
	fmt.Printf("'listen' found = %v\n", trie.Find("listen"))

	trie.Insert("listened")
	fmt.Printf("'listened' found = %v\n", trie.Find("listened"))

	fmt.Println(trie.Remove("listen"))

	fmt.Printf("'listen' found after remove? = %v\n", trie.Find("listen"))
	fmt.Printf("'listened' found after remove? = %v\n", trie.Find("listened"))
}
