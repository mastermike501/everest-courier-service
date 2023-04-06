package main

import "github.com/mastermike501/everest-courier-service/ev_package"

/*
  This implementation of the 0-1 knapsack problem uses dynamic programming to
	compute the maximum value that can be achieved for each subset of items and
	capacity, and then backtracks through the table to determine which items were
	selected. The item struct represents an item with a given weight and value.
	The knapsack function takes an array of items and a capacity, and returns an
	array of the indices of the selected items.
*/

type KItem struct {
	weight float64
	value  float64
	pkg    *ev_package.Package
}

func KnapsackSolver(items []KItem, capacity float64) []int {
	// Create a 2D slice to store the maximum values for each capacity and subset of items
	dp := make([][]float64, len(items)+1)
	for i := range dp {
		dp[i] = make([]float64, int(capacity)+1)
	}

	// Iterate through each item and capacity, computing the maximum value for each subset of items and capacity
	for i := 1; i <= len(items); i++ {
		for j := 1; j <= int(capacity); j++ {
			if items[i-1].weight > float64(j) {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-int(items[i-1].weight)]+items[i-1].value)
			}
		}
	}

	// Determine which items were selected by backtracking through the dp table
	selected := make([]int, 0)
	i, j := len(items), int(capacity)
	for i > 0 && j > 0 {
		if dp[i][j] != dp[i-1][j] {
			selected = append(selected, i-1)
			j -= int(items[i-1].weight)
		}
		i--
	}

	// Reverse the order of the selected items and return them
	for i, j := 0, len(selected)-1; i < j; i, j = i+1, j-1 {
		selected[i], selected[j] = selected[j], selected[i]
	}
	return selected
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
