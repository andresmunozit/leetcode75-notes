---
name: tests
description: create a test file in js (node) or go inside of the leetcode exercise name, main_test.go or index.test.js
---

Every folder of this repo has the following structure:

01-exercise-name:

110-balanced-binary-tree/
├── README.md
├── go
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── main_test.go
└── js
    ├── index.js
    ├── index.test.js
    ├── jest.config.js
    └── package.json
└── img
    ├── 110-1.jpg
    └── 110-2.jpg

Write the corresponding tests for go and js (node) files, use table tests. I'm going to give you an
example of both a go and a js test:
```go
package main

import (
	"testing"

	h "github.com/andresmunozit/leetcode75-notes/helpers-go"
	"github.com/stretchr/testify/require"
)

func Test_isBalanced(t *testing.T) {
	tcases := []struct {
		name     string
		root     *h.TreeNode
		expected bool
	}{
		{
			name: "balanced tree 1",
			root: &h.TreeNode{
				Val: 3,
				Left: &h.TreeNode{
					Val: 9,
				},
				Right: &h.TreeNode{
					Val: 20,
					Left: &h.TreeNode{
						Val: 15,
					},
					Right: &h.TreeNode{
						Val: 7,
					},
				},
			},
			expected: true,
		},
		{
			name: "unbalanced tree 1",
			root: &h.TreeNode{
				Val: 1,
				Left: &h.TreeNode{
					Val: 2,
					Left: &h.TreeNode{
						Val: 3,
						Left: &h.TreeNode{
							Val: 4,
						},
						Right: &h.TreeNode{
							Val: 4,
						},
					},
					Right: &h.TreeNode{
						Val: 3,
					},
				},
				Right: &h.TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
		{
			name:     "root is nil",
			expected: true,
		},
		{
			name: "unbalanced tree 2",
			root: &h.TreeNode{
				Val: 1,
				Left: &h.TreeNode{
					Val: 2,
					Left: &h.TreeNode{
						Val: 3,
						Left: &h.TreeNode{
							Val: 4,
						},
					},
				},
				Right: &h.TreeNode{
					Val: 2,
					Right: &h.TreeNode{
						Val: 3,
						Right: &h.TreeNode{
							Val: 4,
						},
					},
				},
			},
			expected: false,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, isBalanced(tc.root))
		})
	}
}

```

```js
const { longestCommonPrefix } = require('.');

describe('longestCommonPrefix', () => {
    const tests = [
        {
            name: 'there is common prefix',
            strs: ['flower', 'flow', 'flight'],
            expected: 'fl',
        },
        {
            name: 'there is no common prefix',
            strs: ['dog', 'racecar', 'car'],
            expected: '',
        },
        {
            name: 'all the strings contain the prefix',
            strs: ['cat', 'cat', 'cat'],
            expected: 'cat',
        },
    ];

    tests.forEach((t) => {
        test(t.name, () => {
            expect(longestCommonPrefix(t.strs)).toEqual(t.expected);
        });
    });
});

```

For both `js` and `go`, use the examples present in the `README.md` for creating the test cases, for
the current corresponding exercise, take the name of the tested function from the `index.test.js` or
`main_test.go` files.

For js use only `jest` and for go use only  `fastify/require`.

For go only write the `go.mod` file and then run `go mod tidy` to generate `go.sum`. For running
tests in go use `go test -v -cover *.go`.
