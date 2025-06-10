package problems

import "fmt"

// Fibonacci implements an unoptimized Fibonacci sequence with O(2^n) runtime.
// Space complexity is O(n).
func Fibonacci(i int) int {
	if i <= 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return Fibonacci(i-1) + Fibonacci(i-2)
}

// FibonacciWithMemoization implements an Fibonacci sequence with O(n) runtime via memoization.
// Space complexity is O(n).
// If memo is nil, a new map will be allocated to avoid operations on nil pointers.
func FibonacciWithMemoization(i int, memo map[int]int) int {
	if memo == nil {
		memo = make(map[int]int, 0)
	}
	if i <= 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	if v, ok := memo[i]; ok {
		return v
	}
	result := FibonacciWithMemoization(i-1, memo) + FibonacciWithMemoization(i-2, memo)
	memo[i] = result
	return result
}

// CountPaths returns the number of paths from coordinates (x,y) to (size-1,size-1) in the given
// square grid.
// Memoization reduces the recursive runtime from O(2^n-squared) to O(n-squared).
// A dynamic programming approach (summing paths from bottom-right to top-left) slightly reduces
// space complexity by eliminating the recursive call stack, though runtime is still O(n-squared).
func CountPaths(grid [][]int, size int, row int, col int, memo map[string]int) int {
	if memo == nil {
		memo = make(map[string]int, 0)
	}
	if !validCell(grid, row, col) {
		return 0
	}
	if col == size-1 && row == size-1 {
		return 1
	}
	if memo[coordsToKey(row, col)] == 0 {
		memo[coordsToKey(row, col)] = CountPaths(grid, size, row+1, col, memo) + CountPaths(grid, size, row, col+1, memo)
	}
	return memo[coordsToKey(row, col)]
}

func validCell(grid [][]int, row int, col int) bool {
	if row >= len(grid) {
		return false
	}
	if col >= len(grid[row]) {
		return false
	}
	if grid[row][col] == 1 {
		return false
	}
	return true
}

func coordsToKey(row int, col int) string {
	return fmt.Sprintf("%d,%d", row, col)
}
