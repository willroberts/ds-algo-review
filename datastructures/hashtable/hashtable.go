package hashtable

import (
	"container/list"
	"errors"

	"github.com/willroberts/ds-algo-review/algorithms/hash/fnv1"
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
	hash := fnv1.FNV1([]byte(key))
	mapped := hash % ht.size
	if ht.table[mapped] == nil {
		ht.table[mapped] = list.New()
	}
	list := ht.table[mapped]
	list.PushBack(key)
	list.PushBack(value)
}

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
