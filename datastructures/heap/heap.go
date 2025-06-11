package heap

import "errors"

var (
	// ErrHeapIsEmpty safeguards against invalid indexing of slices when the heap is empty.
	ErrHeapIsEmpty = errors.New("heap is empty")

	// ErrNoRootParent safeguards against nil pointer access when invalid node access occurs.
	ErrNoRootParent = errors.New("cannot get parent of root node")
)
