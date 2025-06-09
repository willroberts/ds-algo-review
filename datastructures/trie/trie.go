package trie

// Trie (AKA PrefixTree) implements a character-based tree for fast word lookups.
type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

type TrieNode struct {
	children       map[rune]*TrieNode
	isCompleteWord bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode, 0),
	}
}

func (t *Trie) InsertWord(s string) {
	currentNode := t.root
	for _, r := range s {
		currentNode.children[r] = NewTrieNode()
		currentNode = currentNode.children[r]
	}
	currentNode.isCompleteWord = true
}

func (t *Trie) IsCompleteWord(s string) bool {
	currentNode := t.root
	for _, r := range s {
		if _, ok := currentNode.children[r]; !ok {
			return false
		}
		currentNode = currentNode.children[r]
	}
	return true
}

func debugKeys(m map[rune]*TrieNode) []rune {
	keys := make([]rune, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
