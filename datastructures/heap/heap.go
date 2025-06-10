package heap

import "errors"

var (
	ErrHeapIsEmpty  = errors.New("heap is empty")
	ErrNoRootParent = errors.New("cannot get parent of root node")
)
