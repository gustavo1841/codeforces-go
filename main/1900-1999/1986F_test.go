// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1986/F
// https://codeforces.com/problemset/status/1986/problem/F?friends=on
func Test_cf1986F(t *testing.T) {
	testCases := [][2]string{
		{
			`6
2 1
1 2
3 3
1 2
2 3
1 3
5 5
1 2
1 3
3 4
4 5
5 3
6 7
1 2
1 3
2 3
3 4
4 5
4 6
5 6
5 5
1 2
1 3
2 3
2 4
3 5
10 12
1 2
1 3
2 3
2 4
4 5
5 6
6 7
7 4
3 8
8 9
9 10
10 8`,
			`0
3
4
6
6
21`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1986F)
}
