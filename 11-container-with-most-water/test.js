const assert = require('assert')
const maxArea = require('./index')

const tt = [
    {
        height: [1,8,6,2,5,4,8,3,7],
        expected: 49
    },
    {
        height: [1,1],
        expected: 1
    },
]

for (const t of tt) {
    assert.strict.deepEqual(maxArea(t.height), t.expected)
}
