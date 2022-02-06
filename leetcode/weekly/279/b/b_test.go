// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_b(t *testing.T) {
	examples := [][]string{
		{
			`310`, 
			`103`,
		},
		{
			`-7605`, 
			`-7650`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, smallestNumber, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-279/problems/smallest-value-of-the-rearranged-number/
