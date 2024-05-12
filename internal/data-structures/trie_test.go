package datastructures

import (
	"testing"
)

func Test_TrieInsert(t *testing.T) {
	trie := NewTrie()

	trie.Insert("foo")
	trie.Insert("bar")

	if !trie.Search("foo") {
		t.Errorf("Expected true, got false")
	}
	if !trie.Search("bar") {
		t.Errorf("Expected true, got false")
	}
	if trie.Search("holanda") {
		t.Errorf("Expected false, got true")
	}
}

func Test_TrieInsertMany(t *testing.T) {
	trie := NewTrie()

	trie.InsertMany([]string{"foo", "bar"})

	if !trie.Search("foo") {
		t.Errorf("Expected true, got false")
	}
	if !trie.Search("bar") {
		t.Errorf("Expected true, got false")
	}
	if trie.Search("holanda") {
		t.Errorf("Expected false, got true")
	}
}

func Test_StartsWith(t *testing.T) {
	trie := NewTrie()

	trie.InsertMany([]string{"foo", "bar"})

	if !trie.StartsWith("fo") {
		t.Errorf("Expected true, got false")
	}
	if !trie.StartsWith("b") {
		t.Errorf("Expected true, got false")
	}
	if trie.StartsWith("ho") {
		t.Errorf("Expected false, got true")
	}
}
