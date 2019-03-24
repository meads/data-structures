package trie

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

var (
	// ErrSuffixesFound is an error value for when a word cannot be deleted because it has dependent children TrieNodes'
	ErrSuffixesFound = errors.New("dependent suffixes exist preventing word deletion")
)

var alph = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

// Node represents a node a Trie data structure
type Node struct {
	Children        map[string]*Node
	Val             string
	CompletesString bool
}

// NewNode creates an instance of TrieNode with the supplied letter for its' value
func NewNode(letter string, parent *Node) *Node {
	return &Node{
		Children: make(map[string]*Node),
		Val:      letter,
	}
}

// Trie is a Tree like data structure which associates a prefix string with "branches" of suffixes
type Trie struct {
	RootNode *Node
}

// NewTrie creates an instance of Trie with a root node
func NewTrie() *Trie {
	return &Trie{
		RootNode: NewNode("", nil),
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
	val := ""
	for i := 0; i < len(letters); i++ {
		currentLetter := letters[i]
		val += currentLetter
		if v, ok := node.Children[currentLetter]; ok {
			node = v
		} else {
			newNode := NewNode(val, node)
			node.Children[currentLetter] = newNode
			node = newNode
		}
	}
	node.CompletesString = true
}

// Exists returns a boolean indicating that the word exists in the Trie
func (t *Trie) Exists(word string) bool {
	w := strings.TrimSpace(word)
	if len(w) == 0 {
		return false
	}
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

// FindCompletesString finds the leaf node for a word in the trie
func (t *Trie) FindCompletesString(word string) *Node {
	w := strings.TrimSpace(word)
	if len(w) == 0 {
		return nil
	}
	node := t.RootNode
	letters := strings.Split(word, "")
	for i := 0; i < len(letters); i++ {
		currentLetter := letters[i]
		if v, ok := node.Children[currentLetter]; ok {
			node = v
		} else {
			return nil
		}
	}

	return node
}

// Search given a prefix string will suggest next words.
func (t *Trie) Search(prefix string) []string {
	node := t.RootNode
	letters := strings.Split(prefix, "")
	possibleSuffixes := &[]string{}

	for i := 0; i < len(letters); i++ {
		currentLetter := letters[i]
		if childNode, ok := node.Children[currentLetter]; ok {
			node = childNode
		}
	}

	// recursively walk the sub trees of child nodes of the last node processed in the 'prefix' to find the closest
	// completesString suffixes that are contained in nearby subtrees
	for i := 0; i < 25; i++ {
		if v, exists := node.Children[alph[i]]; exists {
			w := ""
			searchRecur(prefix, w, v, possibleSuffixes)
		}
	}

	sort.Slice(*possibleSuffixes, func(i, j int) bool {
		return (*possibleSuffixes)[i] < (*possibleSuffixes)[j]
	})

	return *possibleSuffixes

}

func searchRecur(prefix, accumulator string, node *Node, suffixes *[]string) string {
	accumulator += node.Val

	if node.CompletesString {
		return node.Val
	}

	for _, v := range node.Children {
		if retval := searchRecur(prefix, accumulator, v, suffixes); retval != "" {
			val := v.Val
			*suffixes = append(*suffixes, strings.Replace(val, prefix, "", 1))
			accumulator = ""
		}
	}

	return ""
}

// String writes the string representation of the Trie to the supplied Buffer
func (t *Trie) String() {
	b, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

// Remove returns a string value indicating the result of attempting to remove an entire word from the Trie structure
// or an error indicating the problem encountered
func (t *Trie) Remove(word string) (string, error) {
	node := t.RootNode
	suffixes := []*Node{}

	letters := strings.Split(word, "")

	// walk the trie structure for each letter of 'word', determining at the end if the word has children before proceeding
	for i := 0; i < len(letters); i++ {
		currentLetter := letters[i]
		if v, ok := node.Children[currentLetter]; ok {
			node = v
			// add the suffixes in reverse order a.k.a. unshift so 'word' will be [d,r,o,w] <- TrieNodes
			suffixes = append([]*Node{node}, suffixes...)

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
