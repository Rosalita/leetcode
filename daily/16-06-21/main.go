package main

import (
	"fmt"
)

// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge-2021/605/week-3-june-15th-june-21st/3781/


func main() {

	// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

	// as there are only 8 possible inputs, it could probably be faster to just store the outputs in a map.
	// however as the outputs are thousands of lines long, so possibly better to generate them.

	// this problem is a recursion problem.

	// because brackets come in pairs, the length of each generated string will be n * 2
	// the string will be made of left brackets characters ( and right bracket characters )
	// when forming valid brackets, each left bracket must have a matching right bracket

	// there are two important cases
	// while ranging over string, the ith char can only become a left ( if there are < n ('s
	// also the ith character can only be ) if the count of ( is greater than the number of )'s

	result := generateParentheses(4)
	fmt.Println(result)
}

func recursion(n int, str string, solutions *[]string) {
	// if the lenght of the string is n * 2 then the string is a solution.
	// because for n there are two kinds of bracket.
	if len(str) == n * 2 {
		*solutions = append(*solutions, str)
	} else {
		// count the number of left and right brackets in the string
		leftCount := 0
		rightCount := 0
		for _, char := range str {
			if char == '(' {
				leftCount++
			}
			if char == ')' {
				rightCount++
			}
		}
		// if there are the same amount of ( and )
		if leftCount == rightCount {
			// then every bracket has a pair, so add a new left bracket
			recursion(n, str+"(", solutions)
		} else if leftCount > rightCount{
			if leftCount < n {
				recursion(n, str+"(", solutions)
			}
			// at this point, the leftCount is n so the final closing ) is needed 
			recursion(n, str+")", solutions)
		}
	}
}

func generateParentheses(n int) []string{
	var solutions []string
    recursion(n, "", &solutions);
    return solutions;
}