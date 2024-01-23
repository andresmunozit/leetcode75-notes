const assert = require('assert')
const closeStrings = require('./index')

const tests = [
    {
        word1: 'abc',
        word2: 'bca',
        expectedOutput: true
    },
    {
        word1: 'a',
        word2: 'aa',
        expectedOutput: false
    },
    {
        word1: 'cabbba',
        word2: 'abbccc',
        expectedOutput: true
    },
    {
        word1: 'abbzzca',
        word2: 'babzzcz',
        expectedOutput: false
    },
]

for (const t of tests) {
    assert.strictEqual(closeStrings(t.word1, t.word2), t.expectedOutput)
}
