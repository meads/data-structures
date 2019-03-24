package trie

import (
	"reflect"
	"sort"
	"testing"
)

func TestSearch_Suggestions_Are_Returned(t *testing.T) {

	words := []string{
		"aardvark",
		"aardvarks",
		"aardwolf",
		"aardwolves",
		"aargh",
		"aaron",
		"aaronic",
		"aaronical",
		"aaronite",
		"aaronitic",
		"aarrgh",
		"aarrghh",
		"aaru",
	}

	/*
						 ""
						 |
						"a"
						 |
						"a"
						 |
						"r"
						 |
			  "d"    "g"     "o"  "r"  "u"
			   |      |       |    |
			--------  |       |    |
		  "v"  "w"  |	     "n"  "g"
			 |	  |	  |       |    |
			"a"	 "o" "h"      |   "h"
			|	  |		     "i"   |
			"r"	" l"  	      |   "h"
			 |	  |	          |
			"k"  -------
			|    "f" "v"      -----------
			"s"	 			|	  	 "t"     "c"
				   		 "e"	   -------    |
				 				|      "e" "i"   "a"
				   		 "s"		    |     |
						     				"c"   "l"


	*/

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
		"u", // fail ??
	}

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
	if err == nil {
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
		t.Errorf("expected '%s' to be found after leaf TrieNode removal but 'false' returned from Find operation", word)
	}
}
