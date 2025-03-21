// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/30/E
// https://codeforces.com/problemset/status/30/problem/E?friends=on
func Test_cf30E(t *testing.T) {
	testCases := [][2]string{
		{
			`abacaba`,
			`1
1 7`,
		},
		{
			`axbya`,
			`3
1 1
2 1
5 1`,
		},
		{
			`xabyczba`,
			`3
2 2
4 1
7 2`,
		},
		{
			`ntomzzuttrtaapousysvfgelrpqrnljqvfmcyytiheqnjuhpln`,
			`3
1 1
9 3
50 1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf30E)
}
