const assert = require('assert')
const longestOnes = require('./index')

const tests = [
    {
        nums: [1,1,1,0,0,0,1,1,1,1,0],
        k: 2,
        expectedOutput: 6,
    },
    {
        nums: [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1],
        k: 3,
        expectedOutput: 10,
    },
]

for (const t of tests) {
    assert.strict.deepEqual(longestOnes(t.nums, t.k), t.expectedOutput)
}
