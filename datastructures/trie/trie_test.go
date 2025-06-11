package trie

import "testing"

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.InsertWord("foo")
	if !trie.IsCompleteWord("foo") {
		t.Fail()
	}
}
