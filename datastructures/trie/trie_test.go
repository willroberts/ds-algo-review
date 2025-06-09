package trie

import (
	"os"
	"strings"
	"testing"
)

const wordsFile = "/usr/share/dict/words"

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.InsertWord("foo")
	if !trie.IsCompleteWord("foo") {
		t.Fail()
	}
}

func loadWords() ([]string, error) {
	b, err := os.ReadFile(wordsFile)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(string(b), "\n"), nil
}
