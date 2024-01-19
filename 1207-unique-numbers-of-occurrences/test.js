const assert = require('assert')
const uniqueOccurrences = require('./index')

const tests = [
    {
        arr: [1,2,2,1,1,3],
        output: true,
    },
    {
        arr: [1,2],
        output: false,
    },
    {
        arr: [-3,0,1,-3,1,1,1,-3,10,0],
        output: true,
    },
]

for (const t of tests) {
    assert.strict.deepEqual(uniqueOccurrences(t.arr), t.output)
}
