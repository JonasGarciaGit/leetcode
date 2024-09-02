package main

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

// Example 1:

// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]
// Example 2:

// Input: n = 1
// Output: ["()"]

// Constraints:

// 1 <= n <= 8

func generateParenthesis(n int) []string {
	response := []string{}
	backtracking(&response, 0, 0, n, "")
	return response
}

func backtracking(response *[]string, left, right, n int, pairs string) {
	if len(pairs) == n*2 {
		*response = append(*response, pairs)
		return
	}

	if left < n {
		backtracking(response, left+1, right, n, pairs+"(")
	}

	if right < left {
		backtracking(response, left, right+1, n, pairs+")")
	}
}

func main() {
	generateParenthesis(2)
}
