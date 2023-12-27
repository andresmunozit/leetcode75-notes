const assert = require('assert')
const compress = require('./index')

const tests = [
    {
        chars: ['a','a','b','b','c','c','c'],
        expectedLength: 6,
        expectedCompressed: ['a','2','b','2','c','3'],
    },
    {
        chars: ['a'],
        expectedLength: 1,
        expectedCompressed: ['a'],
    },
    {
        chars: ['a','b','b','b','b','b','b','b','b','b','b','b','b'],
        expectedLength: 4,
        expectedCompressed: ['a','b','1','2'],
    },
]

for (const t of tests) {
    assert.strict.deepEqual(compress(t.chars), t.expectedLength)
    assert.strict.deepEqual(t.chars.slice(0, t.expectedLength), t.expectedCompressed)
}
