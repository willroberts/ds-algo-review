package trie

// Trie (AKA PrefixTree) implements a character-based tree for fast word lookups.
type Trie struct {
	root *TrieNode
}

// NewTrie instantiates a Trie and preallocates its root node.
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

// TrieNode represents a node in the Trie. It contains a lookup table of child nodes.
type TrieNode struct {
	children       map[rune]*TrieNode
	isCompleteWord bool
}

// NewTrieNode instantiates a TrieNode and preallocates its hash table of children.
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode, 0),
	}
}

// InsertWord adds the given word to the Trie.
func (t *Trie) InsertWord(s string) {
	currentNode := t.root
	for _, r := range s {
		currentNode.children[r] = NewTrieNode()
		currentNode = currentNode.children[r]
	}
	currentNode.isCompleteWord = true
}

// IsCompleteWord returns 'true' when the given string is a complete word in the Trie.
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
