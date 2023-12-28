const assert = require('assert')
const moveZeroes = require('./index')

const tt = [
    {
        nums: [0,1,0,3,12],
        expected: [1,3,12,0,0,]
    },
    {
        nums: [0],
        expected: [0]
    },
    {
        nums: [0,1],
        expected: [1,0]
    }
]

for (const t of tt) {
    moveZeroes(t.nums)
    assert.strict.deepEqual(t.nums, t.expected)
}
