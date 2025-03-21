// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2065/D
// https://codeforces.com/problemset/status/2065/problem/D?friends=on
func Test_cf2065D(t *testing.T) {
	testCases := [][2]string{
		{
			`3
2 2
4 4
6 1
3 4
2 2 2 2
3 2 1 2
4 1 2 1
2 3
3 4 5
1 1 9`,
			`41
162
72`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2065D)
}
