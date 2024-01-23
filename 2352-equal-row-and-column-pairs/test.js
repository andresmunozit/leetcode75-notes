const assert = require('assert')
const equalPairs = require('./index')

const tests = [
    {
        grid: [[3,2,1],[1,7,6],[2,7,7]],
        expectedOutput: 1
    },
    {
        grid: [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]],
        expectedOutput: 3,
    },
    {
        grid: [[11,1],[1,11]],
        expectedOutput: 2,
    },
]

for (const t of tests) {
    assert.strict.strictEqual(equalPairs(t.grid), t.expectedOutput)
}
