package main

import "fmt"

// Given an integer numRows, return the first numRows of Pascal's triangle.

// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

// Example 1:

// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

func generate(numRows int) [][]int {
	if numRows == 0 {
		return make([][]int, 0)
	}

	if numRows == 1 {
		var slice [][]int
		return append(slice, []int{1})
	}

	prevRows := generate(numRows - 1)
	newRow := []int{}

	for i := 0; i < numRows; i++ {
		newRow = append(newRow, 1)
	}

	for i := 1; i < numRows-1; i++ {
		newRow[i] = prevRows[numRows-2][i-1] + prevRows[numRows-2][i]
	}

	prevRows = append(prevRows, newRow)
	return prevRows
}

func main() {
	fmt.Println(generate(4))
}
