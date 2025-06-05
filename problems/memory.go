package problems

// Memory_AllocateToStack stores data in thread-specific stack memory.
//go:noinline
func Memory_AllocateToStack() int {
	x := 1
	return x // Stays on the stack, does not escape.
}

// Memory_AllocateToStackWithCopy demonstrates copying values without heap allocations.
//go:noinline
func Memory_AllocateToStackWithCopy() int {
	x := 1
	y := x
	return y
}

// Memory_AllocateToStackWithPointer demonstrates using reference types without heap allocations.
//go:noinline
func Memory_AllocateToStackWithPointer() int {
	x := 1
	return memory_DereferenceInt(&x)
}

//go:noinline
func memory_DereferenceInt(i *int) int {
	return *i // Does not escape.
}

// Memory_AllocateToHeap stores data in the garbage-collected heap space.
//go:noinline
func Memory_AllocateToHeap() *int {
	x := 1
	return &x // Moves to heap.
}
