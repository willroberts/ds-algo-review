package problems

import "testing"

func TestFibonacci(t *testing.T) {
	if Fibonacci(6) != 8 {
		t.Fail()
	}
}

func TestFibonacciWithMemoization(t *testing.T) {
	if FibonacciWithMemoization(6, nil) != 8 {
		t.Fail()
	}
}

func TestCountPaths(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}
	result := CountPaths(grid, 4, 0, 0, nil)
	if result != 4 {
		t.Log("expected 4, got", result)
		t.Fail()
	}
}
