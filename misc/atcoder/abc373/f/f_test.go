// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc373/tasks/abc373_f
// 提交：https://atcoder.jp/contests/abc373/submit?taskScreenName=abc373_f
// 对拍：https://atcoder.jp/contests/abc373/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc373_f&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc373/submissions?f.Status=AC&f.Task=abc373_f&orderBy=source_length
func Test_f(t *testing.T) {
	testCases := [][2]string{
		{
			`2 10
3 4
3 2`,
			`5`,
		},
		{
			`3 6
1 4
2 3
2 7`,
			`14`,
		},
		{
			`1 10
1 7`,
			`12`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
