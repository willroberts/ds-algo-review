package problems

// AllocateToStack stores data in thread-specific stack memory.
//go:noinline
func Memory_AllocateToStack() int {
	x := 1
	return x // Stays on the stack, does not escape.
}

func Memory_AllocateToStackWithCopy() int {
	x := 1
	y := x
	return y
}

func Memory_AllocateToStackWithPointer() int {
	x := 1
	return memory_DereferenceInt(&x)
}

func memory_DereferenceInt(i *int) int {
	return *i // Does not escape.
}

// AllocateToHeap stores data in the garbage-collected heap space.
//go:noinline
func Memory_AllocateToHeap() *int {
	x := 1
	return &x // Moves to heap.
}
