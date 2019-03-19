package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

var (
	// ErrSuffixesFound is an error value for when a word cannot be deleted because it has dependent children TrieNodes'
	ErrSuffixesFound = errors.New("canot delete word, suffixes exist")
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
	w := strings.TrimSpace(word)
	if len(w) == 0 {
		return
	}
	node := t.RootNode
	letters := strings.Split(w, "")
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
	w := strings.TrimSpace(word)
	if len(w) == 0 {
		return false
	}
	node := t.RootNode
	letters := strings.Split(w, "")
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

// String writes the string representation of the Trie to the supplied Buffer
func (t *Trie) String(w io.Writer) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}

// Remove returns a string value indicating the result of attempting to remove an entire word from the Trie structure
// or an error indicating the problem encountered
func (t *Trie) Remove(word string) (string, error) {
	node := t.RootNode
	suffixes := []*TrieNode{}

	letters := strings.Split(word, "")

	// walk the trie structure for each letter of 'word', determining at the end if the word has children before proceeding
	for i := 0; i < len(letters); i++ {
		currentLetter := letters[i]
		if v, ok := node.Children[currentLetter]; ok {
			node = v
			// add the suffixes in reverse order a.k.a. unshift so 'word' will be [d,r,o,w] <- TrieNodes
			suffixes = append([]*TrieNode{node}, suffixes...)

			// can we even proceed with removal ?
			if i == len(letters) && len(node.Children) > 0 {
				return "", ErrSuffixesFound
			}
		}
	}

	// for each letter in 'word' work backwards from the edge removing a trie node from trie each go
	for j := 1; j < len(suffixes); j++ {
		childLetter := string(word[len(suffixes)-j]) // last character in the string "word", e.g. "d"

		// the node in 'suffixes' representing the parent node of the node with 'childLetter' as its 'Val' in the Trie
		parent := suffixes[j]

		if childNode, exists := parent.Children[childLetter]; exists {
			if len(childNode.Children) > 0 {
				return "", ErrSuffixesFound
			}
			delete(parent.Children, childLetter)
		}

		if parent.CompletesString || len(parent.Children) > 0 {
			return fmt.Sprintf("some suffixes of '%s', were removed from trie", word), nil
		}
	}

	// if we got this far, we are able to remove the root node of 'word' from the RootNode
	delete(t.RootNode.Children, string(word[0]))

	return fmt.Sprintf("removed '%s'; no other '%s' = words remain", word, string(word[0])), nil
}

func main() {}
