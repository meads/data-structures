package trie

import (
	"reflect"
	"sort"
	"testing"

	"github.com/pkg/errors"
)

func TestExists_Returns_True_When_Word_Is_Found(t *testing.T) {
	sut := NewTrie()
	sut.Insert("apple")
	if !sut.Exists("apple") {
		t.Error("expected 'word' to be found after insert")
	}
}

func TestExists_Returns_False_When_Word_Is_Not_Found(t *testing.T) {
	sut := NewTrie()
	word := "apple"
	sut.Insert(word)
	if sut.Exists("appl") {
		t.Errorf("expected '%s' to NOT be found when the entire term alone is not a valid word", word)
	}
}

func TestExists_Returns_False_When_Word_Is_Empty_String(t *testing.T) {
	sut := NewTrie()
	word := "apple"
	sut.Insert(word)
	if sut.Exists(" ") {
		t.Error("expected false as return value when supplying empty string")
	}
}

func TestInsert_Word_Inserted_Can_Be_Found(t *testing.T) {
	sut := NewTrie()
	word := "test"
	sut.Insert(word)
	if !sut.Exists(word) {
		t.Errorf("expected '%s' to be found after Insert but 'false' returned from Find operation", word)
	}
}

func TestInsert_Blank_Word_Or_Whitespace_Is_Ignored(t *testing.T) {
	sut := NewTrie()
	sut.Insert("")
	if sut.Exists("") {
		t.Errorf("expected blank string to be ignored from being added found TrieNode with empty string as Val")
	}
	sut = NewTrie()
	sut.Insert(" ")
	if sut.Exists(" ") {
		t.Errorf("expected blank string to be ignored from being added found TrieNode with empty string as Val")
	}
}

func TestInsert_Word_Suffix_Inserted_Can_Be_Found_Along_With_Parent(t *testing.T) {
	sut := NewTrie()

	word := "test"
	sut.Insert(word)

	wordSuffix := "testing"
	sut.Insert(wordSuffix)

	if !sut.Exists(word) {
		t.Errorf("expected '%s' to be found after Insert but 'false' returned from Find operation", word)
	}
	if !sut.Exists(wordSuffix) {
		t.Errorf("expected '%s' to be found after Insert but 'false' returned from Find operation", wordSuffix)
	}
}

func TestRemove_Word_Removed_Can_Not_Be_Found(t *testing.T) {
	sut := NewTrie()
	word := "test"
	sut.Insert(word)
	if !sut.Exists(word) {
		t.Errorf("expected '%s' to be found after Insert but 'false' returned from Find operation", word)
	}
	_, err := sut.Remove(word)
	if err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
	if sut.Exists(word) {
		t.Errorf("expected '%s' to NOT be found after Remove but 'true' returned from Find operation", word)
	}
}

func TestRemove_Non_Existing_Word_Supplied_Is_Ignored_No_Error(t *testing.T) {
	sut := NewTrie()
	if msg, err := sut.Remove("invalid"); err != nil {
		t.Errorf("expected '%v', got '%v'", nil, err)
		t.Errorf(msg)
	}
}

func TestRemove_Word_Doesnt_Allow_Removal_When_Children_Suffixes_Exist(t *testing.T) {
	sut := NewTrie()
	word := "test"
	sut.Insert(word)
	wordSuffix := "testing"
	sut.Insert(wordSuffix)

	_, err := sut.Remove(word)
	if errors.Cause(err) != ErrSuffixesFound {
		t.Errorf("expected '%v' got '%v'", ErrSuffixesFound, err)
	}
}

func TestRemove_Word_Suffix_Removed_Doesnt_Affect_Parent(t *testing.T) {
	sut := NewTrie()

	word := "test"
	sut.Insert(word)

	wordSuffix := "testing"
	sut.Insert(wordSuffix)

	_, err := sut.Remove(wordSuffix)
	if err != nil {
		t.Errorf("expected '%v' got '%v'", nil, err)
	}

	if !sut.Exists(word) {
		t.Errorf("expected '%s' to be found after leaf node removal", word)
	}
}

func TestSearch_Suggestions_Are_Empty_Given_Invalid_Term(t *testing.T) {
	sut := NewTrie()
	sut.Insert("apple")
	sut.Insert("orange")

	expected := []string{}

	actual := sut.Search("invalid")
	if !reflect.DeepEqual(expected, actual) {
		sut.String()
		t.Errorf("\nexpected\n%#v\ngot\n%#v\n", expected, actual)
	}
}

func TestSearch_Suggestions_Are_Returned(t *testing.T) {

	words := []string{
		"aardvark", // --dvark
		"aardvarks",
		"aardwolf",   // --dwolf
		"aardwolves", // --dwolves
		"aargh",      // --gh
		"aaron",      // --on
		"aaronic",
		"aaronical",
		"aaronite",
		"aaronitic",
		"aarrgh", // --rgh
		"aarrghh",
		"aaru",
	}

	sut := NewTrie()
	for _, w := range words {
		sut.Insert(w)
	}
	expected := []string{
		"dvark",
		"dwolf",
		"dwolves",
		"gh",
		"on",
		"rgh",
	}

	// verify that every word in the trie has been properly inserted and completesString flag is true too
	for _, el := range words {
		n := sut.FindCompletesString(el)
		if n == nil || !n.CompletesString {
			t.Errorf("expected string to exist in trie %s and complete the string", el)
		}
	}

	actual := sut.Search("aar")
	sort.Slice(expected, func(i, j int) bool { return expected[i] > expected[j] })
	sort.Slice(actual, func(i, j int) bool { return actual[i] > actual[j] })
	if !reflect.DeepEqual(expected, actual) {
		sut.String()
		t.Errorf("\nexpected\n%#v\ngot\n%#v\n", expected, actual)
	}
}
