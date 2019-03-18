package main

import (
	"bytes"
	"testing"
)

func Test_Word_Inserted_Can_Be_Found(t *testing.T) {
	sut := NewTrie()
	word := "test"
	sut.Insert(word)
	if !sut.Find(word) {
		t.Errorf("expected '%s' to be found after Insert but 'false' returned from Find operation", word)
	}
}

func Test_Word_Removed_Can_Not_Be_Found(t *testing.T) {
	sut := NewTrie()
	word := "test"
	sut.Insert(word)
	if !sut.Find(word) {
		t.Errorf("expected '%s' to be found after Insert but 'false' returned from Find operation", word)
	}
	_, err := sut.Remove(word)
	if err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
	if sut.Find(word) {
		t.Errorf("expected '%s' to NOT be found after Remove but 'true' returned from Find operation", word)
	}
}

func Test_Word_Suffix_Inserted_Can_Be_Found_Along_With_Parent(t *testing.T) {
	sut := NewTrie()

	word := "test"
	sut.Insert(word)

	wordSuffix := "testing"
	sut.Insert(wordSuffix)

	if !sut.Find(word) {
		t.Errorf("expected '%s' to be found after Insert but 'false' returned from Find operation", word)
	}
	if !sut.Find(wordSuffix) {
		t.Errorf("expected '%s' to be found after Insert but 'false' returned from Find operation", wordSuffix)
	}
}

func Test_Word_Doesnt_Allow_Removal_When_Children_Suffixes_Exist(t *testing.T) {
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

func Test_Word_Suffix_Removed_Doesnt_Affect_Parent(t *testing.T) {
	sut := NewTrie()

	word := "test"
	sut.Insert(word)

	wordSuffix := "testing"
	sut.Insert(wordSuffix)

	_, err := sut.Remove(wordSuffix)
	if err != nil {
		t.Errorf("expected '%v' got '%v'", nil, err)
	}

	if !sut.Find(word) {
		t.Errorf("expected '%s' to be found after leaf TrieNode removal but 'false' returned from Find operation", word)
	}
}

func Test_String_Properly_Displays_Trie_Structure(t *testing.T) {
	sut := NewTrie()
	sut.Insert("test")
	sut.Insert("testing")

	expected := `{"RootNode":{"Children":{"t":{"Children":{"e":{"Children":{"s":{"Children":{"t":{"Children":{"i":{"Children":{"n":{"Children":{"g":{"Children":{},"Val":"g","CompletesString":true}},"Val":"n","CompletesString":false}},"Val":"i","CompletesString":false}},"Val":"t","CompletesString":true}},"Val":"s","CompletesString":false}},"Val":"e","CompletesString":false}},"Val":"t","CompletesString":false}},"Val":"","CompletesString":false}}`

	var buf bytes.Buffer
	sut.String(&buf)
	trieString := buf.String()
	if trieString != expected {
		t.Errorf("expected\n\n'%s'\ngot\n\n'%s'", expected, trieString)
	}
}
