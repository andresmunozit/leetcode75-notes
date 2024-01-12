const assert = require('assert')
const findMaxAverage = require('./index')

const tests = [
    {
        nums: [1,12,-5,-6,50,3],
        k: 4,
        expectedMaxAv: 12.75,
    },
    {
        nums: [5],
        k: 1,
        expectedMaxAv: 5,
    },
]

for (const t of tests) {
    assert.strict.deepEqual(findMaxAverage(t.nums, t.k), t.expectedMaxAv)
}
