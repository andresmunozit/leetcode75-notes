const assert = require('assert')
const maxOperations = require('./index')

const tests = [
    {
        nums: [1,2,3,4],
        k: 5,
        expected: 2,
    },
    {
        nums: [3,1,3,4,3],
        k: 6,
        expected: 1,
    },
]

for (const t of tests) {
    assert.strict.deepEqual(maxOperations(t.nums, t.k), t.expected)
}
