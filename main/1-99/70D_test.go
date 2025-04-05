// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/70/problem/D
// https://codeforces.com/problemset/status/70/problem/D?friends=on
func Test_cf70D(t *testing.T) {
	testCases := [][2]string{
		{
			`8
1 0 0
1 2 0
1 2 2
2 1 0
1 0 2
2 1 1
2 2 1
2 20 -1`,
			`YES
YES
YES
NO`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf70D)
}
