package hashtable

import (
	"container/list"
	"errors"

	"github.com/willroberts/ds-algo-review/algorithms/hash/fnv1"
)

// HashTable implements a hash table with a fixed number of slots, which are modulo-mapped using
// the FNV1a hash algorithm. Each slot corresponds to a linked list of keys and values in order
// to handle collisions.
type HashTable struct {
	size  uint64
	table []*list.List
}

// NewHashTable instantiates a HashTable and preallocates its slots.
func NewHashTable(size uint64) *HashTable {
	return &HashTable{
		size:  size,
		table: make([]*list.List, size),
	}
}

// Insert adds the given key and value to the HashTable.
func (ht *HashTable) Insert(key string, value string) {
	hash := fnv1.FNV1([]byte(key))
	mapped := hash % ht.size
	if ht.table[mapped] == nil {
		ht.table[mapped] = list.New()
	}
	list := ht.table[mapped]
	list.PushBack(key)
	list.PushBack(value)
}

// Get retrieves the value for the given key from the HashTable.
func (ht *HashTable) Get(key string) (string, error) {
	hash := fnv1.FNV1([]byte(key))
	mapped := hash % ht.size
	list := ht.table[mapped]
	node := list.Front()
	if node == nil {
		return "", errors.New("list does not exist")
	}
	for e := node; e != nil; e = e.Next() {
		if e.Value == key {
			return e.Next().Value.(string), nil
		}
	}
	return "", errors.New("value not found")
}
