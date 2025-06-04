package hashtable

import (
	"container/list"
	"errors"

	"github.com/willroberts/ds-algo-review/algorithms/fnv1hash"
)

type HashTable struct {
	size  uint64
	table []*list.List
}

func NewHashTable(size uint64) *HashTable {
	return &HashTable{
		size:  size,
		table: make([]*list.List, size),
	}
}

func (ht *HashTable) Insert(key string, value string) {
	hash := fnv1hash.FNV1([]byte(key))
	mapped := hash % ht.size
	if ht.table[mapped] == nil {
		ht.table[mapped] = list.New()
	}
	list := ht.table[mapped]
	list.PushBack(key)
	list.PushBack(value)
}

func (ht *HashTable) Get(key string) (string, error) {
	hash := fnv1hash.FNV1([]byte(key))
	mapped := hash % ht.size
	list := ht.table[mapped]
	node := list.Front()
	if node == nil {
		return "", errors.New("list does not exist")
	}

	for {
		if node.Next() == nil {
			break
		}
		if node.Value == key {
			return node.Next().Value.(string), nil
		}
		node = node.Next()
	}

	return "", errors.New("value not found")
}
